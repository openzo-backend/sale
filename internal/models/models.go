// package models

// type OrderItem struct {
// 	ID             int    `json:"id"`
// 	ProductID      string `json:"product_id" gorm:"size:36"`
// 	SizeVariantID  int    `json:"size_variant_id"`
// 	ColorVariantID int    `json:"color_variant_id"`
// 	SaleId         string `json:"sale_id" gorm:"size:36"`
// 	Quantity       int    `json:"quantity"`
// 	Status         string `json:"status,omitempty"`
// }

// type Sale struct {
// 	ID                 string   `json:"id"`
// 	Customer           Customer `json:"customer"`
// 	OrderTime          string   `json:"order_time"`
// 	TotalAmount        float64  `json:"total_amount"`
// 	UndiscountedAmount float64  `json:"undiscounted_amount"`
// 	AmountPaid         float64  `json:"amount_paid"`
// 	PaymentMode        string   `json:"payment_mode"`
// 	StoreID            string   `json:"store_id"`

// 	OrderItems []OrderItem `json:"order_items"`

// 	Type   string `json:"type"`
// 	Status string `json:"status,omitempty" `

// 	// Restaurant specific fields
// 	TableId   string `json:"table_id,omitempty"`
// 	Occupants int    `json:"occupants,omitempty"`
// 	Note      string `json:"note,omitempty"`
// }

// type Customer struct {
// 	ID         int    `json:"id"`
// 	Name       string `json:"name"`
// 	Phone      string `json:"phone"`
// 	Address    string `json:"address"`
// 	UserDataId string `json:"user_data_id" gorm:"size:36"`
// 	SaleId     string `json:"sale_id" gorm:"size:36"`
// }

package models

type OrderItem struct {
	ID             int    `json:"id" gorm:"primaryKey"`
	ProductID      string `json:"product_id" gorm:"size:36"`
	SizeVariantID  int    `json:"size_variant_id"`
	ColorVariantID int    `json:"color_variant_id"`
	SaleID         string `json:"sale_id" gorm:"size:36;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Quantity       int    `json:"quantity"`
	Status         string `json:"status,omitempty"`
}

type Sale struct {
	ID                 string      `json:"id" gorm:"primaryKey;size:36"`
	CustomerID         int         `json:"customer_id" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Customer           Customer    `json:"customer" gorm:"foreignKey:CustomerID"`
	OrderTime          string      `json:"order_time"`
	TotalAmount        float64     `json:"total_amount"`
	UndiscountedAmount float64     `json:"undiscounted_amount"`
	AmountPaid         float64     `json:"amount_paid"`
	PaymentMode        string      `json:"payment_mode"`
	StoreID            string      `json:"store_id"`
	OrderItems         []OrderItem `json:"order_items" gorm:"foreignKey:SaleID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Type               string      `json:"type"`
	Status             string      `json:"status,omitempty"`
	TableID            string      `json:"table_id,omitempty"`
	Occupants          int         `json:"occupants,omitempty"`
	Note               string      `json:"note,omitempty"`
}

type Customer struct {
	ID         int    `json:"id" gorm:"primaryKey"`
	Name       string `json:"name"`
	Phone      string `json:"phone"`
	Address    string `json:"address"`
	UserDataID string `json:"user_data_id" gorm:"size:36"`
	SaleID     string `json:"sale_id" gorm:"size:36"`
}
