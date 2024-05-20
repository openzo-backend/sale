package models

import "time"

type OrderItem struct {
	ID        int    `json:"id"`
	ProductID string `json:"product_id" gorm:"size:36"`
	SaleId    string `json:"sale_id" gorm:"size:36"`
	Quantity  int    `json:"quantity"`
}

type Sale struct {
	ID          string      `json:"id"`
	OrderItems  []OrderItem `json:"order_items"`
	Customer    Customer    `json:"customer"`
	OrderTime   time.Time   `json:"order_time" gorm:"autoCreateTime"`
	TotalAmount float64     `json:"total_amount"`
	AmountPaid  float64     `json:"amount_paid"`
	PaymentMode string      `json:"payment_mode"`
	StoreID     string      `json:"store_id"`
}

type Customer struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Phone      string `json:"phone"`
	Address    string `json:"address"`
	UserDataId string `json:"user_data_id" gorm:"size:36"`
	SaleId     string `json:"sale_id" gorm:"size:36"`
}
