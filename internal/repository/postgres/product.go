package postgres

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/satriowisnugroho/catalog/internal/entity"
	"github.com/satriowisnugroho/catalog/internal/helper"
	dbentity "github.com/satriowisnugroho/catalog/internal/repository/postgres/entity"
	"github.com/satriowisnugroho/catalog/internal/response"
)

// ProductRepositoryInterface define contract for product related functions to repository
type ProductRepositoryInterface interface {
	CreateProduct(ctx context.Context, product *entity.Product) error
	GetProductByID(ctx context.Context, productID int) (*entity.Product, error)
	UpdateProduct(ctx context.Context, product *entity.Product) error
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

	// ProductCreationColumns list all columns used for create product
	ProductCreationColumns = ProductColumns[1:]
	// ProductCreationAttributes hold string format of all creation product columns
	ProductCreationAttributes = strings.Join(ProductCreationColumns, ", ")
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
		tmpEntity := dbentity.Product{}
		if err := rows.StructScan(&tmpEntity); err != nil {
			return nil, errors.Wrap(err, "fetch")
		}

		result = append(result, tmpEntity.ToEntity())
	}

	return result, nil
}

// CreateProduct insert product data into database
func (r *ProductRepository) CreateProduct(ctx context.Context, product *entity.Product) error {
	functionName := "ProductRepository.CreateProduct"

	if err := helper.CheckDeadline(ctx); err != nil {
		return errors.Wrap(err, functionName)
	}

	now := time.Now()
	product.CreatedAt = now
	product.UpdatedAt = now
	product.SKU = helper.GenerateSKU()

	query := fmt.Sprintf(`INSERT INTO %s (%s) VALUES (%s) RETURNING id`, ProductTableName, ProductCreationAttributes, EnumeratedBindvars(ProductCreationColumns))
	log.Print(query)

	err := r.db.QueryRowContext(ctx, query,
		product.SKU,
		product.Title,
		product.Category,
		product.Condition,
		product.Tenant,
		product.Qty,
		product.Price,
		product.CreatedAt,
		product.UpdatedAt,
	).Scan(&product.ID)
	if err != nil {
		return errors.Wrap(err, functionName)
	}

	return nil
}

// GetProductByID return product by id
func (r *ProductRepository) GetProductByID(ctx context.Context, productID int) (*entity.Product, error) {
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

// UpdateProduct update a product
func (r *ProductRepository) UpdateProduct(ctx context.Context, product *entity.Product) error {
	functionName := "ProductRepository.UpdateProduct"

	if err := helper.CheckDeadline(ctx); err != nil {
		return errors.Wrap(err, functionName)
	}

	now := time.Now()
	product.UpdatedAt = now

	query := fmt.Sprintf("UPDATE %s SET %s WHERE id = $%d", ProductTableName, UpdateColumnsValues(ProductCreationColumns), len(ProductColumns))

	_, err := r.db.ExecContext(
		ctx,
		query,
		product.SKU,
		product.Title,
		product.Category,
		product.Condition,
		product.Tenant,
		product.Qty,
		product.Price,
		product.CreatedAt,
		product.UpdatedAt,
		product.ID,
	)
	if err != nil {
		return errors.Wrap(err, functionName)
	}

	return nil
}
