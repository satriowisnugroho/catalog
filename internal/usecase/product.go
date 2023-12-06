package usecase

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"github.com/satriowisnugroho/catalog/internal/entity"
	"github.com/satriowisnugroho/catalog/internal/helper"
	repo "github.com/satriowisnugroho/catalog/internal/repository/postgres"
	"github.com/satriowisnugroho/catalog/internal/response"
)

// ProductUsecaseInterface define contract for product related functions to usecase
type ProductUsecaseInterface interface {
	CreateProduct(ctx context.Context, payload *entity.ProductPayload) (*entity.Product, error)
	BulkReduceQtyProduct(ctx context.Context, payload *entity.BulkReduceQtyProductPayload) ([]*entity.Product, error)
	GetProductByID(ctx context.Context, productID int) (*entity.Product, error)
	GetProducts(ctx context.Context, payload *entity.GetProductPayload) ([]*entity.Product, int, error)
	UpdateProduct(ctx context.Context, productID int, payload *entity.ProductPayload) (*entity.Product, error)
}

type ProductUsecase struct {
	repo repo.ProductRepositoryInterface
}

func NewProductUsecase(r repo.ProductRepositoryInterface) *ProductUsecase {
	return &ProductUsecase{
		repo: r,
	}
}

func (uc *ProductUsecase) CreateProduct(ctx context.Context, payload *entity.ProductPayload) (*entity.Product, error) {
	functionName := "ProductUsecase.CreateProduct"

	if err := helper.CheckDeadline(ctx); err != nil {
		return nil, errors.Wrap(err, functionName)
	}

	product := payload.ToEntity()
	if err := uc.repo.CreateProduct(ctx, product); err != nil {
		return nil, errors.Wrap(fmt.Errorf("uc.repo.CreateProduct: %w", err), functionName)
	}

	return product, nil
}

func (uc *ProductUsecase) BulkReduceQtyProduct(ctx context.Context, payload *entity.BulkReduceQtyProductPayload) ([]*entity.Product, error) {
	functionName := "ProductUsecase.BulkReduceQtyProduct"

	if err := helper.CheckDeadline(ctx); err != nil {
		return nil, errors.Wrap(err, functionName)
	}

	// TODO: Use transaction
	products := make([]*entity.Product, 0)
	for _, item := range payload.Items {
		product, err := uc.repo.GetProductBySKU(ctx, item.SKU)
		if err != nil {
			if err == response.ErrNotFound {
				return nil, err
			}

			return nil, errors.Wrap(fmt.Errorf("uc.repo.GetProductBySKU: %w", err), functionName)
		}

		// TODO: Check tenant, if it doesn't belong to product, return 403 forbiden

		product.Qty = product.Qty - item.ReqQty
		if product.Qty < 0 {
			return nil, response.ErrInsufficientStock
		}

		if err := uc.repo.UpdateProduct(ctx, product); err != nil {
			return nil, errors.Wrap(fmt.Errorf("uc.repo.UpdateProduct: %w", err), functionName)
		}
	}

	return products, nil
}

func (uc *ProductUsecase) GetProductByID(ctx context.Context, productID int) (*entity.Product, error) {
	functionName := "ProductUsecase.GetProductByID"

	if err := helper.CheckDeadline(ctx); err != nil {
		return nil, errors.Wrap(err, functionName)
	}

	product, err := uc.repo.GetProductByID(ctx, productID)
	if err != nil {
		if err == response.ErrNotFound {
			return nil, err
		}

		return nil, errors.Wrap(fmt.Errorf("uc.repo.GetProductByID: %w", err), functionName)
	}

	return product, nil
}

func (uc *ProductUsecase) GetProducts(ctx context.Context, payload *entity.GetProductPayload) ([]*entity.Product, int, error) {
	functionName := "ProductUsecase.GetProducts"

	if err := helper.CheckDeadline(ctx); err != nil {
		return nil, 0, errors.Wrap(err, functionName)
	}

	products, err := uc.repo.GetProducts(ctx, payload)
	if err != nil {
		return nil, 0, errors.Wrap(fmt.Errorf("uc.repo.GetProducts: %w", err), functionName)
	}

	count, err := uc.repo.GetProductsCount(ctx, payload)
	if err != nil {
		return nil, 0, errors.Wrap(fmt.Errorf("uc.repo.GetProductsCount: %w", err), functionName)
	}

	return products, count, nil
}

func (uc *ProductUsecase) UpdateProduct(ctx context.Context, productID int, payload *entity.ProductPayload) (*entity.Product, error) {
	functionName := "ProductUsecase.UpdateProduct"

	if err := helper.CheckDeadline(ctx); err != nil {
		return nil, errors.Wrap(err, functionName)
	}

	product, err := uc.repo.GetProductByID(ctx, productID)
	if err != nil {
		if err == response.ErrNotFound {
			return nil, err
		}

		return nil, errors.Wrap(fmt.Errorf("uc.repo.GetProductByID: %w", err), functionName)
	}

	product.Title = payload.Title
	product.Category = payload.Category
	product.Condition = payload.Condition
	product.Tenant = payload.Tenant
	product.Qty = payload.Qty
	product.Price = payload.Price
	if err := uc.repo.UpdateProduct(ctx, product); err != nil {
		return nil, errors.Wrap(fmt.Errorf("uc.repo.UpdateProduct: %w", err), functionName)
	}

	return product, nil
}
