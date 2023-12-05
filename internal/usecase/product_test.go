package usecase_test

import (
	"context"
	"errors"
	"testing"

	"github.com/satriowisnugroho/catalog/internal/entity"
	"github.com/satriowisnugroho/catalog/internal/response"
	"github.com/satriowisnugroho/catalog/internal/usecase"
	"github.com/satriowisnugroho/catalog/test/fixture"
	testmock "github.com/satriowisnugroho/catalog/test/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateProduct(t *testing.T) {
	testcases := []struct {
		name        string
		ctx         context.Context
		data        *entity.ProductPayload
		rProductErr error
		wantErr     bool
	}{
		{
			name:    "deadline context",
			ctx:     fixture.CtxEnded(),
			wantErr: true,
		},
		{
			name:        "failed to create product",
			ctx:         context.Background(),
			data:        &entity.ProductPayload{},
			rProductErr: errors.New("error create product"),
			wantErr:     true,
		},
		{
			name:    "success",
			ctx:     context.Background(),
			data:    &entity.ProductPayload{Title: "New Product"},
			wantErr: false,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			productRepo := &testmock.ProductRepositoryInterface{}
			productRepo.On("CreateProduct", mock.Anything, mock.Anything).Return(tc.rProductErr)

			uc := usecase.NewProductUsecase(productRepo)
			_, err := uc.CreateProduct(tc.ctx, tc.data)
			assert.Equal(t, tc.wantErr, err != nil)
		})
	}
}

func TestGetProductByID(t *testing.T) {
	testcases := []struct {
		name        string
		ctx         context.Context
		rProductRes *entity.Product
		rProductErr error
		wantErr     bool
	}{
		{
			name:    "deadline context",
			ctx:     fixture.CtxEnded(),
			wantErr: true,
		},
		{
			name:        "product is not found",
			ctx:         context.Background(),
			rProductErr: response.ErrNotFound,
			wantErr:     true,
		},
		{
			name:        "failed to get product",
			ctx:         context.Background(),
			rProductErr: errors.New("error get product"),
			wantErr:     true,
		},
		{
			name:        "success",
			ctx:         context.Background(),
			rProductRes: &entity.Product{ID: 123, Title: "New Product"},
			wantErr:     false,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			productRepo := &testmock.ProductRepositoryInterface{}
			productRepo.On("GetProductByID", mock.Anything, mock.Anything).Return(tc.rProductRes, tc.rProductErr)

			uc := usecase.NewProductUsecase(productRepo)
			_, err := uc.GetProductByID(tc.ctx, 123)
			assert.Equal(t, tc.wantErr, err != nil)
		})
	}
}

func TestUpdateProduct(t *testing.T) {
	testcases := []struct {
		name           string
		ctx            context.Context
		productID      int
		data           *entity.ProductPayload
		rGetProductRes *entity.Product
		rGetProductErr error
		rProductErr    error
		wantErr        bool
	}{
		{
			name:    "deadline context",
			ctx:     fixture.CtxEnded(),
			wantErr: true,
		},
		{
			name:           "product is not found",
			ctx:            context.Background(),
			rGetProductErr: response.ErrNotFound,
			wantErr:        true,
		},
		{
			name:           "failed to get product",
			ctx:            context.Background(),
			rGetProductErr: errors.New("error get product"),
			wantErr:        true,
		},
		{
			name:           "failed to update product",
			ctx:            context.Background(),
			productID:      123,
			rGetProductRes: &entity.Product{},
			data:           &entity.ProductPayload{},
			rProductErr:    errors.New("error update product"),
			wantErr:        true,
		},
		{
			name:           "success",
			ctx:            context.Background(),
			productID:      123,
			rGetProductRes: &entity.Product{},
			data:           &entity.ProductPayload{Title: "New Product"},
			wantErr:        false,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			productRepo := &testmock.ProductRepositoryInterface{}
			productRepo.On("GetProductByID", mock.Anything, mock.Anything).Return(tc.rGetProductRes, tc.rGetProductErr)
			productRepo.On("UpdateProduct", mock.Anything, mock.Anything).Return(tc.rProductErr)

			uc := usecase.NewProductUsecase(productRepo)
			_, err := uc.UpdateProduct(tc.ctx, tc.productID, tc.data)
			assert.Equal(t, tc.wantErr, err != nil)
		})
	}
}
