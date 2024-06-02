package service

import (
	"encoding/json"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/gin-gonic/gin"
	"github.com/tanush-128/openzo_backend/sale/internal/models"
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
	kafkaProducer  *kafka.Producer
	// productClient  pb.ProductServiceClient
}

func NewSaleService(saleRepository repository.SaleRepository,
	producer *kafka.Producer,

) SaleService {
	return &saleService{saleRepository: saleRepository, kafkaProducer: producer}
}

func (s *saleService) CreateSale(ctx *gin.Context, req models.Sale) (models.Sale, error) {

	createdSale, err := s.saleRepository.CreateSale(req)
	if err != nil {
		return models.Sale{}, err // Propagate error
	}

	// Produce message to Kafka
	topic := "sales"
	saleMsg, _ := json.Marshal(createdSale)
	msg := &kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          saleMsg,
	}
	s.kafkaProducer.Produce(msg, nil)

	return createdSale, nil
}

func (s *saleService) UpdateSale(ctx *gin.Context, req models.Sale) (models.Sale, error) {
	updatedSale, err := s.saleRepository.UpdateSale(req)
	if err != nil {
		return models.Sale{}, err
	}

	topic := "sales"
	saleMsg, _ := json.Marshal(updatedSale)
	msg := &kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          saleMsg,
	}
	s.kafkaProducer.Produce(msg, nil)

	return updatedSale, nil
}
