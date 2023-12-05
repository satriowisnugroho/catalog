package postgres

import (
	"context"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/satriowisnugroho/catalog/internal/entity"
	"github.com/satriowisnugroho/catalog/internal/helper"
	dbEntity "github.com/satriowisnugroho/catalog/internal/repository/postgres/entity"
	"github.com/satriowisnugroho/catalog/internal/response"
)

// ProductRepositoryInterface define contract for product related functions to repository
type ProductRepositoryInterface interface {
	GetProductByID(ctx context.Context, productID int64) (*entity.Product, error)
}

// ProductRepository holds database connection
type ProductRepository struct {
	db *sqlx.DB
}

var (
	// ProductTableName hold table name for products
	ProductTableName = "products"
	// ProductColumns list all columns on products table
	ProductColumns = []string{"id", "sku", "title", "category", "condition", "tenant", "qty", "price", "created_at", "updated_at"}
	// ProductAttributes hold string format of all products table columns
	ProductAttributes = strings.Join(ProductColumns, ", ")
)

// NewProductRepository create initiate product repository with given database
func NewProductRepository(db *sqlx.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) fetch(ctx context.Context, query string, args ...interface{}) ([]*entity.Product, error) {
	rows, err := r.db.QueryxContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	result := make([]*entity.Product, 0)

	for rows.Next() {
		tmpEntity := dbEntity.Product{}
		if err := rows.StructScan(&tmpEntity); err != nil {
			return nil, errors.Wrap(err, "fetch")
		}

		result = append(result, tmpEntity.ToEntity())
	}

	return result, nil
}

// GetProductByID return product by id
func (r *ProductRepository) GetProductByID(ctx context.Context, productID int64) (*entity.Product, error) {
	functionName := "ProductRepository.GetProductByID"

	if err := helper.CheckDeadline(ctx); err != nil {
		return nil, errors.Wrap(err, functionName)
	}

	query := fmt.Sprintf("SELECT %s FROM %s WHERE id = $1 LIMIT 1", ProductAttributes, ProductTableName)
	rows, err := r.fetch(ctx, query, productID)
	if err != nil {
		return nil, errors.Wrap(err, functionName)
	}

	if len(rows) == 0 {
		return nil, response.ErrNotFound
	}

	return rows[0], nil
}
