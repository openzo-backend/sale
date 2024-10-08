package repository

import (
	"log"

	"github.com/google/uuid"
	"github.com/tanush-128/openzo_backend/sale/internal/models"

	"gorm.io/gorm"
)

type SaleRepository interface {
	CreateSale(Sale models.Sale) (models.Sale, error)
	GetSaleByID(id string) (models.Sale, error)
	GetSalesByStoreID(storeID string) ([]models.Sale, error)
	UpdateSale(Sale models.Sale) (models.Sale, error)
	// GetSalesByUserDataID(ctx *gin.Context, user_data_id string) ([]models.Sale, error)
	GetSalesByUserDataID(user_data_id string) ([]models.Sale, error)
	ChangeSaleStatus(id string, status string) (models.Sale, error)
	DeleteSale(id string) error
	// ChangeSaleStatus(ctx *gin.Context, id string, status string) (models.Sale, error)
	// Add more methods for other Sale operations (GetSaleByEmail, UpdateSale, etc.)

}

type saleRepository struct {
	db *gorm.DB
}

func NewSaleRepository(db *gorm.DB) SaleRepository {

	return &saleRepository{db: db}
}

func (r *saleRepository) CreateSale(Sale models.Sale) (models.Sale, error) {
	Sale.ID = uuid.New().String()
	tx := r.db.Create(&Sale)

	if tx.Error != nil {
		return models.Sale{}, tx.Error
	}

	return Sale, nil
}

func (r *saleRepository) GetSaleByID(id string) (models.Sale, error) {
	var Sale models.Sale
	tx := r.db.Preload("OrderItems").Preload("Customer").Where("id = ?", id).First(&Sale)
	if tx.Error != nil {
		return models.Sale{}, tx.Error
	}

	return Sale, nil
}

func (r *saleRepository) GetSalesByStoreID(storeID string) ([]models.Sale, error) {
	var Sales []models.Sale

	tx := r.db.Preload("OrderItems").Preload("Customer").Where("store_id = ?", storeID).Find(&Sales)
	if tx.Error != nil {
		return []models.Sale{}, tx.Error
	}

	return Sales, nil
}

func (r *saleRepository) GetSalesByUserDataID(user_data_id string) ([]models.Sale, error) {
	var Sales []models.Sale
	tx := r.db.Joins("Customer").Where("user_data_id = ?", user_data_id).Preload("OrderItems").Preload("Customer").Find(&Sales)

	if tx.Error != nil {
		return []models.Sale{}, tx.Error
	}
	// log.Printf("Sales: %+v", Sales)
	for _, sale := range Sales {
		log.Printf("Sale: %+v", sale.ID)
	}
	return Sales, nil
}

func (r *saleRepository) ChangeSaleStatus(id string, status string) (models.Sale, error) {
	var Sale models.Sale
	tx := r.db.Model(&Sale).Where("id = ?", id).Update("status", status)
	if tx.Error != nil {
		return models.Sale{}, tx.Error
	}

	return Sale, nil
}

func (r *saleRepository) UpdateSale(sale models.Sale) (models.Sale, error) {

	// delete existing order items
	tx := r.db.Where("sale_id = ?", sale.ID).Delete(&models.OrderItem{})
	if tx.Error != nil {
		return models.Sale{}, tx.Error
	}

	tx = r.db.Session(&gorm.Session{FullSaveAssociations: true}).Save(&sale)
	if tx.Error != nil {
		return models.Sale{}, tx.Error
	}

	return sale, nil
}

func (r *saleRepository) DeleteSale(id string) error {
	tx := r.db.Where("id = ?", id).Delete(&models.Sale{})
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

// Implement other repository methods (GetSaleByID, GetSaleByEmail, UpdateSale, etc.) with proper error handling
