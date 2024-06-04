package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tanush-128/openzo_backend/sale/internal/models"
	"github.com/tanush-128/openzo_backend/sale/internal/service"
)

type Handler struct {
	saleService service.SaleService
}

func NewHandler(saleService *service.SaleService) *Handler {
	return &Handler{saleService: *saleService}
}

func (h *Handler) CreateSale(ctx *gin.Context) {
	var sale models.Sale

	ctx.ShouldBindJSON(&sale)

	log.Printf("%+v", sale)

	createdSale, err := h.saleService.CreateSale(ctx, sale)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, createdSale)

}

func (h *Handler) GetSaleByID(ctx *gin.Context) {
	id := ctx.Param("id")

	sale, err := h.saleService.GetSaleByID(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, sale)
}

func (h *Handler) GetSalesByStoreID(ctx *gin.Context) {
	storeID := ctx.Param("id")

	sales, err := h.saleService.GetSalesByStoreID(ctx, storeID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, sales)
}

func (h *Handler) GetSalesByUserDataID(ctx *gin.Context) {
	userDataID := ctx.Param("id")

	sales, err := h.saleService.GetSalesByUserDataID(ctx, userDataID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, sales)

}

func (h *Handler) ChangeSaleStatus(ctx *gin.Context) {
	id := ctx.Param("id")
	status := ctx.Param("status")

	changedSale, err := h.saleService.ChangeSaleStatus(ctx, id, status)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, changedSale)
}

func (h *Handler) UpdateSale(ctx *gin.Context) {
	var sale models.Sale
	if err := ctx.BindJSON(&sale); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedSale, err := h.saleService.UpdateSale(ctx, sale)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, updatedSale)
}
