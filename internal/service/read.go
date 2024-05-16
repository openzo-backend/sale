package service

import (
	"github.com/gin-gonic/gin"
	"github.com/tanush-128/openzo_backend/sale/internal/models"
)

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