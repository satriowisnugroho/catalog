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

// ProductPayload holds product payload representative
type ProductPayload struct {
	Title     string `json:"title"`
	Category  string `json:"category"`
	Condition int8   `json:"condition"`
	Tenant    int8   `json:"tenant"`
	Qty       int    `json:"qty"`
	Price     int    `json:"price"`
}

// ToEntity to convert product payload to entity contract
func (p *ProductPayload) ToEntity() *Product {
	return &Product{
		Title:     p.Title,
		Category:  p.Category,
		Condition: p.Condition,
		Tenant:    p.Tenant,
		Qty:       p.Qty,
		Price:     p.Price,
	}
}
