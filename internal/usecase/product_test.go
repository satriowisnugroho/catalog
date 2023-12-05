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
