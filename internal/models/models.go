package models

type OrderItem struct {
	ID        int    `json:"id"`
	ProductID string `json:"product_id" gorm:"size:36"`
	SaleId    string `json:"sale_id" gorm:"size:36"`
	Quantity  int    `json:"quantity"`
	Status    string `json:"status,omitempty"`
}

type Sale struct {
	ID          string   `json:"id"`
	Customer    Customer `json:"customer"`
	OrderTime   string   `json:"order_time"`
	TotalAmount float64  `json:"total_amount"`
	AmountPaid  float64  `json:"amount_paid"`
	PaymentMode string   `json:"payment_mode"`
	StoreID     string   `json:"store_id"`

	OrderItems []OrderItem `json:"order_items"`

	Type string `json:"type"`
	Note string `json:"note,omitempty"`

	// Restaurant specific fields
	TableId   string `json:"table_id,omitempty"`
	Occupants int    `json:"occupants,omitempty"`
	Status    string `json:"status,omitempty" `
}

type Customer struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Phone      string `json:"phone"`
	Address    string `json:"address"`
	UserDataId string `json:"user_data_id" gorm:"size:36"`
	SaleId     string `json:"sale_id" gorm:"size:36"`
}
