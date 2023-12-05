package entity

import (
	"time"
)

// Product struct holds entity of product
type Product struct {
	ID        int       `json:"id"`
	SKU       string    `json:"sku"`
	Title     string    `json:"title"`
	Category  string    `json:"category"`
	Condition int8      `json:"condition"`
	Tenant    int8      `json:"tenant"`
	Qty       int       `json:"qty"`
	Price     int       `json:"price"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
