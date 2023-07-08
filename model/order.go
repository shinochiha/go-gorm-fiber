package model

import (
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	OrderID            string `gorm:"unique"`
	TransactionDetails string
}

// OrderRequest represents the request body for creating an order
type OrderRequest struct {
	OrderID     string  `json:"order_id"`
	Amount      float64 `json:"amount"`
	CustomerID  string  `json:"customer_id"`
	Description string  `json:"description"`
}

// OrderResponse represents the response body for creating an order
type OrderResponse struct {
	RedirectURL string `json:"redirect_url"`
}
