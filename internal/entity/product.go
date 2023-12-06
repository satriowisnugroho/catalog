package entity

import (
	"time"

	"github.com/satriowisnugroho/catalog/internal/entity/types"
)

// Product struct holds entity of product
type Product struct {
	ID        int              `json:"id"`
	SKU       string           `json:"sku"`
	Title     string           `json:"title"`
	Category  string           `json:"category"`
	Condition int8             `json:"condition"`
	Tenant    types.TenantType `json:"tenant"`
	Qty       int              `json:"qty"`
	Price     int              `json:"price"`
	CreatedAt time.Time        `json:"created_at"`
	UpdatedAt time.Time        `json:"updated_at"`
}

// GetProductPayload holds get product payload representative
type GetProductPayload struct {
	SKU          string
	TitleKeyword string
	Category     string
	Condition    int8
	Tenant       types.TenantType
	OrderBy      string
	Offset       int
	Limit        int
}

// BulkReduceQtyProductPayload holds bulk reduce qty product payload representative
type BulkReduceQtyProductPayload struct {
	Items []BulkReduceQtyProductItemPayload `json:"items"`
}

// BulkReduceQtyProductItemPayload holds bulk reduce qty product item payload representative
type BulkReduceQtyProductItemPayload struct {
	SKU    string `json:"sku"`
	ReqQty int    `json:"req_qty"`
}

// ProductPayload holds product payload representative
type ProductPayload struct {
	Title     string           `json:"title"`
	Category  string           `json:"category"`
	Condition int8             `json:"condition"`
	Tenant    types.TenantType `json:"tenant"`
	Qty       int              `json:"qty"`
	Price     int              `json:"price"`
}

// SwaggerProductPayload holds product payload for swagger docs
type SwaggerProductPayload struct {
	Title     string `json:"title"`
	Category  string `json:"category"`
	Condition int8   `json:"condition"`
	Tenant    string `json:"tenant"`
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
