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
	GetSalesByUserDataID(ctx *gin.Context, user_data_id string) ([]models.Sale, error)
	ChangeSaleStatus(ctx *gin.Context, id string, status string) (models.Sale, error)
	UpdateSale(ctx *gin.Context, req models.Sale) (models.Sale, error)
	DeleteSale(ctx *gin.Context, id string) error
}

type saleService struct {
	saleRepository repository.SaleRepository
	kafkaProducer  *kafka.Producer
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



func (s *saleService) GetSalesByUserDataID(ctx *gin.Context, user_data_id string) ([]models.Sale, error) {
	sales, err := s.saleRepository.GetSalesByUserDataID(user_data_id)
	if err != nil {
		return []models.Sale{}, err
	}

	return sales, nil
}

func (s *saleService) ChangeSaleStatus(ctx *gin.Context, id string, status string) (models.Sale, error) {
	updatedSale, err := s.saleRepository.ChangeSaleStatus(id, status)
	if err != nil {
		return models.Sale{}, err
	}
	sale, _ := s.saleRepository.GetSaleByID(id)
	topic := "sales"
	saleMsg, _ := json.Marshal(sale)
	msg := &kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          saleMsg,
	}
	s.kafkaProducer.Produce(msg, nil)

	return updatedSale, nil
}



func (s *saleService) DeleteSale(ctx *gin.Context, id string) error {
	err := s.saleRepository.DeleteSale(id)
	if err != nil {
		return err
	}

	return nil
}

func (s *saleService) GetSaleByID(ctx *gin.Context, id string) (models.Sale, error) {
	sale, err := s.saleRepository.GetSaleByID(id)
	if err != nil {
		return models.Sale{}, err
	}

	return sale, nil
}

func (s *saleService) GetSalesByStoreID(ctx *gin.Context, storeID string) ([]models.Sale, error) {
	sales, err := s.saleRepository.GetSalesByStoreID(storeID)
	if err != nil {
		return []models.Sale{}, err
	}

	return sales, nil
}
