package service

import (
	"encoding/json"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/gin-gonic/gin"
	"github.com/tanush-128/openzo_backend/sale/internal/models"
	"github.com/tanush-128/openzo_backend/sale/internal/repository"
)

// package service

// import (
// 	"encoding/json"

// 	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
// 	"github.com/gin-gonic/gin"
// 	"github.com/tanush-128/openzo_backend/online_order/internal/models"
// 	"github.com/tanush-128/openzo_backend/online_order/internal/repository"
// )

// type OnlineOrderService interface {

// 	//CRUD
// 	CreateOnlineOrder(ctx *gin.Context, req models.OnlineOrder) (models.OnlineOrder, error)
// 	GetOnlineOrderByID(ctx *gin.Context, id string) (models.OnlineOrder, error)
// 	GetOnlineOrdersByStoreID(ctx *gin.Context, store_id string) ([]models.OnlineOrder, error)
// 	GetOnlineOrdersByUserDataId(ctx *gin.Context, user_data_id string) ([]models.OnlineOrder, error)
// 	ChangeOrderStatus(ctx *gin.Context, id string, status string) (models.OnlineOrder, error)
// 	UpdateOnlineOrder(ctx *gin.Context, req models.OnlineOrder) (models.OnlineOrder, error)
// 	DeleteOnlineOrder(ctx *gin.Context, id string) error
// }

// type online_orderService struct {
// 	online_orderRepository repository.OnlineOrderRepository
// 	kafkaProducer          *kafka.Producer
// }

// func NewOnlineOrderService(online_orderRepository repository.OnlineOrderRepository,
// 	kafkaProducer *kafka.Producer,
// ) OnlineOrderService {
// 	return &online_orderService{online_orderRepository: online_orderRepository, kafkaProducer: kafkaProducer}
// }

// func (s *online_orderService) CreateOnlineOrder(ctx *gin.Context, req models.OnlineOrder) (models.OnlineOrder, error) {

// 	topic := "onlineorder"
// 	orderMsg, err := json.Marshal(req)
// 	if err != nil {
// 		return models.OnlineOrder{}, err
// 	}

// 	s.kafkaProducer.Produce(&kafka.Message{
// 		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
// 		Value:          orderMsg,
// 		Key:            []byte("new_cart_created"),
// 	}, nil)
// 	s.kafkaProducer.Flush(15 * 1000)

// 	createdOnlineOrder, err := s.online_orderRepository.CreateOnlineOrder(req)
// 	if err != nil {
// 		return models.OnlineOrder{}, err // Propagate error
// 	}

// 	return createdOnlineOrder, nil
// }

// func (s *online_orderService) ChangeOrderStatus(ctx *gin.Context, id string, status string) (models.OnlineOrder, error) {

// 	if status == "not_placed" {
// 		return models.OnlineOrder{}, nil
// 	}

// 	OnlineOrder, err := s.online_orderRepository.GetOnlineOrderByID(id)
// 	if err != nil {
// 		return models.OnlineOrder{}, err
// 	}

// 	if status == "placed" {
// 		OnlineOrder.OrderStatus = models.OrderPlaced
// 		topic := "onlineorder"
// 		orderStatusMsg, err := json.Marshal(OnlineOrder)
// 		if err != nil {
// 			return models.OnlineOrder{}, err
// 		}

// 		s.kafkaProducer.Produce(&kafka.Message{
// 			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
// 			Value:          orderStatusMsg,
// 			Key:            []byte("order_placed"),
// 		}, nil)
// 		s.kafkaProducer.Flush(15 * 1000)
// 		changedOnlineOrder, err := s.online_orderRepository.UpdateOnlineOrder(OnlineOrder)
// 		if err != nil {
// 			return models.OnlineOrder{}, err
// 		}

// 		return changedOnlineOrder, nil

// 	} else if status == "accepted" {
// 		OnlineOrder.OrderStatus = models.OrderAccepted
// 	} else if status == "rejected" {
// 		OnlineOrder.OrderStatus = models.OrderRejected
// 	} else if status == "out_for_delivery" {
// 		OnlineOrder.OrderStatus = models.OrderOutForDel
// 	} else if status == "cancelled" {
// 		OnlineOrder.OrderStatus = models.OrderCancelled
// 	} else if status == "delivered" {
// 		OnlineOrder.OrderStatus = models.OrderDelivered
// 	}

// 	topic := "order-status-updates"
// 	orderStatusMsg, err := json.Marshal(OnlineOrder)
// 	if err != nil {
// 		return models.OnlineOrder{}, err
// 	}

// 	s.kafkaProducer.Produce(&kafka.Message{
// 		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
// 		Value:          orderStatusMsg,
// 		Key:            []byte("order_status"),
// 	}, nil)
// 	s.kafkaProducer.Flush(15 * 1000)

// 	changedOnlineOrder, err := s.online_orderRepository.UpdateOnlineOrder(OnlineOrder)
// 	if err != nil {
// 		return models.OnlineOrder{}, err
// 	}

// 	return changedOnlineOrder, nil
// }

// func (s *online_orderService) UpdateOnlineOrder(ctx *gin.Context, req models.OnlineOrder) (models.OnlineOrder, error) {

// 	go s.ChangeOrderStatus(ctx, req.ID, string(req.OrderStatus))

// 	updatedOnlineOrder, err := s.online_orderRepository.UpdateOnlineOrder(req)
// 	if err != nil {
// 		return models.OnlineOrder{}, err
// 	}

// 	return updatedOnlineOrder, nil
// }

// func (s *online_orderService) DeleteOnlineOrder(ctx *gin.Context, id string) error {
// 	err := s.online_orderRepository.DeleteOnlineOrder(id)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// converting onlineOrder to a sale with type:"online_order	"
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

// GetSalesByUserDataID(ctx *gin.Context, user_data_id string) ([]models.Sale, error)

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

//delete sale

func (s *saleService) DeleteSale(ctx *gin.Context, id string) error {
	err := s.saleRepository.DeleteSale(id)
	if err != nil {
		return err
	}

	return nil
}
