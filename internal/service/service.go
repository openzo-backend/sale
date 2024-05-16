package service

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/tanush-128/openzo_backend/sale/internal/models"
	"github.com/tanush-128/openzo_backend/sale/internal/pb"
	"github.com/tanush-128/openzo_backend/sale/internal/repository"
)

type SaleService interface {

	//CRUD
	CreateSale(ctx *gin.Context, req models.Sale) (models.Sale, error)
	GetSaleByID(ctx *gin.Context, id string) (models.Sale, error)
	GetSalesByStoreID(ctx *gin.Context, storeID string) ([]models.Sale, error)
	UpdateSale(ctx *gin.Context, req models.Sale) (models.Sale, error)
}

type saleService struct {
	saleRepository repository.SaleRepository
	productClient  pb.ProductServiceClient
}

func NewSaleService(saleRepository repository.SaleRepository,
	productClient pb.ProductServiceClient,
) SaleService {
	return &saleService{saleRepository: saleRepository, productClient: productClient}
}

func (s *saleService) CreateSale(ctx *gin.Context, req models.Sale) (models.Sale, error) {
	for _, item := range req.OrderItems {
		_, err := s.productClient.ChangeProductQuantity(ctx, &pb.ChangeProductQuantityRequest{
			ProductId: item.ProductID,
			Quantity:  int32(item.Quantity),
		})
		if err != nil {
			log.Println("Error changing product quantity: ", err)
			return models.Sale{}, err
		}

	}

	createdSale, err := s.saleRepository.CreateSale(req)
	if err != nil {
		return models.Sale{}, err // Propagate error
	}

	return createdSale, nil
}

func (s *saleService) UpdateSale(ctx *gin.Context, req models.Sale) (models.Sale, error) {
	updatedSale, err := s.saleRepository.UpdateSale(req)
	if err != nil {
		return models.Sale{}, err
	}

	return updatedSale, nil
}
