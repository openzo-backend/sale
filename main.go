package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/tanush-128/openzo_backend/sale/config"
	handlers "github.com/tanush-128/openzo_backend/sale/internal/api"
	"github.com/tanush-128/openzo_backend/sale/internal/middlewares"
	"github.com/tanush-128/openzo_backend/sale/internal/pb"
	"github.com/tanush-128/openzo_backend/sale/internal/repository"
	"github.com/tanush-128/openzo_backend/sale/internal/service"
	"google.golang.org/grpc"
)

var UserClient pb.UserServiceClient

type User2 struct {
}

func main() {

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(fmt.Errorf("failed to load config: %w", err))
	}

	db, err := connectToDB(cfg) // Implement database connection logic
	if err != nil {
		log.Fatal(fmt.Errorf("failed to connect to database: %w", err))
	}

	// // Initialize gRPC server
	// grpcServer := grpc.NewServer()
	// Salepb.RegisterSaleServiceServer(grpcServer, service.NewGrpcSaleService(SaleRepository, SaleService))
	// reflection.Register(grpcServer) // Optional for server reflection

	//Initialize gRPC client
	conn, err := grpc.Dial(cfg.UserGrpc, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewUserServiceClient(conn)
	UserClient = c

	productConn, err := grpc.Dial(cfg.ProductGrpc, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer productConn.Close()
	productClient := pb.NewProductServiceClient(productConn)

	saleRepository := repository.NewSaleRepository(db)
	SaleService := service.NewSaleService(saleRepository, productClient)

	// Initialize HTTP server with Gin
	router := gin.Default()
	handler := handlers.NewHandler(&SaleService)

	router.GET("ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// router.Use(middlewares.JwtMiddleware(c))

	router.GET("/store/:id", handler.GetSalesByStoreID)
	router.POST("/", handler.CreateSale)
	router.GET("/:id", handler.GetSaleByID)
	router.Use(middlewares.NewMiddleware(c).JwtMiddleware)
	router.PUT("/", handler.UpdateSale)

	// router.Use(middlewares.JwtMiddleware)

	router.Run(fmt.Sprintf(":%s", cfg.HTTPPort))

}
