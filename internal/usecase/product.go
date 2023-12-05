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
	CreateProduct(ctx context.Context, productPayload *entity.ProductPayload) (*entity.Product, error)
	GetProductByID(ctx context.Context, productID int64) (*entity.Product, error)
}

type ProductUsecase struct {
	repo repo.ProductRepositoryInterface
}

func NewProductUsecase(r repo.ProductRepositoryInterface) *ProductUsecase {
	return &ProductUsecase{
		repo: r,
	}
}

func (uc *ProductUsecase) CreateProduct(ctx context.Context, productPayload *entity.ProductPayload) (*entity.Product, error) {
	functionName := "ProductUsecase.CreateProduct"

	if err := helper.CheckDeadline(ctx); err != nil {
		return nil, errors.Wrap(err, functionName)
	}

	product := productPayload.ToEntity()
	if err := uc.repo.CreateProduct(ctx, product); err != nil {
		return nil, errors.Wrap(fmt.Errorf("uc.repo.CreateProduct: %w", err), functionName)
	}

	return product, nil
}

func (uc *ProductUsecase) GetProductByID(ctx context.Context, productID int64) (*entity.Product, error) {
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
