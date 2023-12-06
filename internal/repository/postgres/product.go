package postgres

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"github.com/pkg/errors"
	"github.com/satriowisnugroho/catalog/internal/config"
	"github.com/satriowisnugroho/catalog/internal/entity"
	"github.com/satriowisnugroho/catalog/internal/entity/types"
	"github.com/satriowisnugroho/catalog/internal/helper"
	dbentity "github.com/satriowisnugroho/catalog/internal/repository/postgres/entity"
	"github.com/satriowisnugroho/catalog/internal/response"
)

// ProductRepositoryInterface define contract for product related functions to repository
type ProductRepositoryInterface interface {
	CreateProduct(ctx context.Context, product *entity.Product) error
	GetProductByID(ctx context.Context, productID int) (*entity.Product, error)
	GetProductBySKU(ctx context.Context, productSKU string) (*entity.Product, error)
	GetProducts(ctx context.Context, payload *entity.GetProductPayload) ([]*entity.Product, error)
	GetProductsCount(ctx context.Context, payload *entity.GetProductPayload) (int, error)
	UpdateProduct(ctx context.Context, dbTrx interface{}, product *entity.Product) error
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

	// eligibleOrderByFields list all eligible order by field
	eligibleOrderByFields = []string{"created_at"}
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
		if postgresError, ok := err.(*pq.Error); ok {
			if postgresError.Code == pq.ErrorCode(config.UniqueConstraintViolationCode) && postgresError.Constraint == config.SKUTenantUniqueConstraint {
				return response.ErrDuplicateSKUTenant
			}
		}
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

// GetProductBySKU return product by sku
func (r *ProductRepository) GetProductBySKU(ctx context.Context, productSKU string) (*entity.Product, error) {
	functionName := "ProductRepository.GetProductBySKU"

	if err := helper.CheckDeadline(ctx); err != nil {
		return nil, errors.Wrap(err, functionName)
	}

	query := fmt.Sprintf("SELECT %s FROM %s WHERE sku = $1 LIMIT 1", ProductAttributes, ProductTableName)
	rows, err := r.fetch(ctx, query, productSKU)
	if err != nil {
		return nil, errors.Wrap(err, functionName)
	}

	if len(rows) == 0 {
		return nil, response.ErrNotFound
	}

	return rows[0], nil
}

// GetProducts query to get product list
func (r *ProductRepository) GetProducts(ctx context.Context, payload *entity.GetProductPayload) ([]*entity.Product, error) {
	functionName := "ProductRepository.GetProducts"
	if err := helper.CheckDeadline(ctx); err != nil {
		return []*entity.Product{}, errors.Wrap(err, functionName)
	}

	if payload.Limit == 0 {
		payload.Limit = 10
	} else if payload.Limit > 100 {
		payload.Limit = 100
	}

	orderBy := "id DESC"
	if len(payload.OrderBy) > 0 {
		parts := strings.Split(payload.OrderBy, " ")
		orderByField := parts[0]
		orderByType := parts[1]

		if helper.StringInArray(orderByField, eligibleOrderByFields) {
			if orderByType != "ASC" {
				orderByType = "DESC"
			}
			orderBy = fmt.Sprintf("%s %s", orderByField, orderByType)
		}
	}

	filterQuery, params := r.constructSearchQuery(payload)
	query := fmt.Sprintf("SELECT %s FROM %s %s ORDER BY %s OFFSET %d LIMIT %d", ProductAttributes, ProductTableName, filterQuery, orderBy, payload.Offset, payload.Limit)

	rows, err := r.fetch(ctx, query, params...)
	if err != nil {
		return rows, errors.Wrap(err, functionName)
	}

	return rows, nil
}

// GetProductsCount query to get count of product list
func (r *ProductRepository) GetProductsCount(ctx context.Context, payload *entity.GetProductPayload) (int, error) {
	functionName := "ProductRepository.GetProductsCount"
	if err := helper.CheckDeadline(ctx); err != nil {
		return 0, errors.Wrap(err, functionName)
	}

	filterQuery, params := r.constructSearchQuery(payload)
	query := fmt.Sprintf("SELECT COUNT(*) FROM %s %s", ProductTableName, filterQuery)

	count := 0
	rows := r.db.QueryRowxContext(ctx, query, params...)
	if err := rows.Scan(&count); err != nil {
		return count, errors.Wrap(err, functionName)
	}

	return count, nil
}

// UpdateProduct update a product
func (r *ProductRepository) UpdateProduct(ctx context.Context, dbTrx interface{}, product *entity.Product) error {
	functionName := "ProductRepository.UpdateProduct"

	if err := helper.CheckDeadline(ctx); err != nil {
		return errors.Wrap(err, functionName)
	}

	now := time.Now()
	product.UpdatedAt = now

	query := fmt.Sprintf("UPDATE %s SET %s WHERE id = $%d", ProductTableName, UpdateColumnsValues(ProductCreationColumns), len(ProductColumns))

	tx := Tx(r.db, dbTrx)
	_, err := tx.ExecContext(
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

// constructSearchQuery construct search query
func (r *ProductRepository) constructSearchQuery(payload *entity.GetProductPayload) (string, []interface{}) {
	var params []interface{}
	filterQuery := ""
	wheres := []string{}
	paramIndex := 1

	if len(payload.SKU) > 0 {
		wheres = append(wheres, fmt.Sprintf("sku = $%v", paramIndex))
		params = append(params, payload.SKU)
		paramIndex++
	}

	if len(payload.TitleKeyword) >= 3 {
		wheres = append(wheres, fmt.Sprintf("title ILIKE $%v", paramIndex))
		params = append(params, fmt.Sprintf("%%%s%%", payload.TitleKeyword))
		paramIndex++
	}

	if payload.Category != types.CategoryEmptyType {
		wheres = append(wheres, fmt.Sprintf("category = $%v", paramIndex))
		params = append(params, strconv.FormatInt(int64(payload.Category), 10))
		paramIndex++
	}

	if payload.Condition != types.ConditionEmptyType {
		wheres = append(wheres, fmt.Sprintf("condition = $%v", paramIndex))
		params = append(params, strconv.FormatInt(int64(payload.Condition), 10))
		paramIndex++
	}

	wheres = append(wheres, fmt.Sprintf("tenant = $%v", paramIndex))
	params = append(params, strconv.FormatInt(int64(payload.Tenant), 10))
	paramIndex++

	if len(wheres) > 0 {
		filterQuery = fmt.Sprintf("WHERE %s", strings.Join(wheres, " AND "))
	}

	return filterQuery, params
}
