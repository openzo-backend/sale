package models

import "time"

type OrderItem struct {
	ID             int    `json:"id"`
	ProductID      string `json:"product_id" gorm:"size:36"`
	SizeVariantID  int    `json:"size_variant_id"`
	ColorVariantID int    `json:"color_variant_id"`
	SaleId         string `json:"sale_id" gorm:"size:36"`
	Quantity       int    `json:"quantity"`
	Status         string `json:"status,omitempty"`
}

type Sale struct {
	ID                 string   `json:"id" gorm:"primaryKey;size:36;unique"`
	Customer           Customer `json:"customer"`
	OrderTime          string   `json:"order_time"`
	TotalAmount        float64  `json:"total_amount"`
	UndiscountedAmount float64  `json:"undiscounted_amount"`
	AmountPaid         float64  `json:"amount_paid"`
	PaymentMode        string   `json:"payment_mode"`
	StoreID            string   `json:"store_id"`

	OrderItems []OrderItem `json:"order_items"`

	Type   string `json:"type"`
	Status string `json:"status,omitempty" `

	// Restaurant specific fields
	TableId   string `json:"table_id,omitempty"`
	Occupants int    `json:"occupants,omitempty"`

	Note string `json:"note,omitempty"`

	// Booking specific fields
	BookingTime *time.Time `json:"booking_time,omitempty"`
}

type Customer struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	Phone           string `json:"phone"`
	Address         string `json:"address"`
	UserDataID      string `json:"user_data_id" gorm:"size:36"`
	StoreCustomerID string `json:"store_customer_id" gorm:"size:36"`
	SaleId          string `json:"sale_id" gorm:"size:36"`
}
