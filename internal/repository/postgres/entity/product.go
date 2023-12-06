package entity

import (
	"time"

	"github.com/satriowisnugroho/catalog/internal/entity"
	"github.com/satriowisnugroho/catalog/internal/entity/types"
)

// Product struct holds attachment database representative
type Product struct {
	ID        int                 `db:"id"`
	SKU       string              `db:"sku"`
	Title     string              `db:"title"`
	Category  types.CategoryType  `db:"category"`
	Condition types.ConditionType `db:"condition"`
	Tenant    types.TenantType    `db:"tenant"`
	Qty       int                 `db:"qty"`
	Price     int                 `db:"price"`
	CreatedAt time.Time           `db:"created_at"`
	UpdatedAt time.Time           `db:"updated_at"`
}

// ToEntity to convert product from database to entity contract
func (p *Product) ToEntity() *entity.Product {
	return &entity.Product{
		ID:        p.ID,
		SKU:       p.SKU,
		Title:     p.Title,
		Category:  p.Category,
		Condition: p.Condition,
		Tenant:    p.Tenant,
		Qty:       p.Qty,
		Price:     p.Price,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}
}
