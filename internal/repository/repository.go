package repository

import (
	"github.com/google/uuid"
	"github.com/tanush-128/openzo_backend/sale/internal/models"

	"gorm.io/gorm"
)

type SaleRepository interface {
	CreateSale(Sale models.Sale) (models.Sale, error)
	GetSaleByID(id string) (models.Sale, error)
	GetSalesByStoreID(storeID string) ([]models.Sale, error)
	UpdateSale(Sale models.Sale) (models.Sale, error)
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

func (r *saleRepository) UpdateSale(Sale models.Sale) (models.Sale, error) {
	tx := r.db.Save(&Sale)
	if tx.Error != nil {
		return models.Sale{}, tx.Error
	}

	return Sale, nil
}

// Implement other repository methods (GetSaleByID, GetSaleByEmail, UpdateSale, etc.) with proper error handling
