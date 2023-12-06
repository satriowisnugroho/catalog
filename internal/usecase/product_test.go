package usecase_test

import (
	"context"
	"errors"
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/satriowisnugroho/catalog/internal/entity"
	"github.com/satriowisnugroho/catalog/internal/entity/types"
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

			uc := usecase.NewProductUsecase(productRepo, &testmock.PostgresTransactionRepositoryInterface{})
			_, err := uc.CreateProduct(tc.ctx, tc.data)
			assert.Equal(t, tc.wantErr, err != nil)
		})
	}
}

func TestBulkReduceQtyProduct(t *testing.T) {
	testcases := []struct {
		name              string
		ctx               context.Context
		tenant            types.TenantType
		payload           *entity.BulkReduceQtyProductPayload
		rStartTrxErr      error
		rCommitTrxErr     error
		rGetProductRes    *entity.Product
		rGetProductErr    error
		rUpdateProductErr error
		wantErr           bool
	}{
		{
			name:    "deadline context",
			ctx:     fixture.CtxEnded(),
			wantErr: true,
		},
		{
			name:         "failed to start transaction",
			ctx:          context.Background(),
			rStartTrxErr: response.ErrNoSQLTransactionFound,
			wantErr:      true,
		},
		{
			name:           "product is not found",
			ctx:            context.Background(),
			payload:        &entity.BulkReduceQtyProductPayload{Items: []entity.BulkReduceQtyProductItemPayload{{SKU: "SKU-123", ReqQty: 1}}},
			rGetProductErr: response.ErrNotFound,
			wantErr:        true,
		},
		{
			name:           "failed to get product",
			ctx:            context.Background(),
			payload:        &entity.BulkReduceQtyProductPayload{Items: []entity.BulkReduceQtyProductItemPayload{{SKU: "SKU-123", ReqQty: 1}}},
			rGetProductErr: errors.New("error get product"),
			wantErr:        true,
		},
		{
			name:           "forbidden",
			ctx:            context.Background(),
			tenant:         types.TenantIpsumType,
			payload:        &entity.BulkReduceQtyProductPayload{Items: []entity.BulkReduceQtyProductItemPayload{{SKU: "SKU-123", ReqQty: 11}}},
			rGetProductRes: &entity.Product{Qty: 10, Tenant: types.TenantLoremType},
			wantErr:        true,
		},
		{
			name:           "insufficient stock",
			ctx:            context.Background(),
			payload:        &entity.BulkReduceQtyProductPayload{Items: []entity.BulkReduceQtyProductItemPayload{{SKU: "SKU-123", ReqQty: 11}}},
			rGetProductRes: &entity.Product{Qty: 10},
			wantErr:        true,
		},
		{
			name:              "failed to update product",
			ctx:               context.Background(),
			payload:           &entity.BulkReduceQtyProductPayload{Items: []entity.BulkReduceQtyProductItemPayload{{SKU: "SKU-123", ReqQty: 1}}},
			rGetProductRes:    &entity.Product{Qty: 10},
			rUpdateProductErr: errors.New("error update product"),
			wantErr:           true,
		},
		{
			name:           "failed to commit transaction",
			ctx:            context.Background(),
			payload:        &entity.BulkReduceQtyProductPayload{Items: []entity.BulkReduceQtyProductItemPayload{{SKU: "SKU-123", ReqQty: 1}}},
			rGetProductRes: &entity.Product{Qty: 10},
			rCommitTrxErr:  response.ErrNoSQLTransactionFound,
			wantErr:        true,
		},
		{
			name:           "success",
			ctx:            context.Background(),
			rGetProductRes: &entity.Product{Qty: 10},
			payload:        &entity.BulkReduceQtyProductPayload{Items: []entity.BulkReduceQtyProductItemPayload{{SKU: "SKU-123", ReqQty: 1}}},
			wantErr:        false,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			productRepo := &testmock.ProductRepositoryInterface{}
			productRepo.On("GetProductBySKU", mock.Anything, mock.Anything).Return(tc.rGetProductRes, tc.rGetProductErr)
			productRepo.On("UpdateProduct", mock.Anything, mock.Anything, mock.Anything).Return(tc.rUpdateProductErr)

			dbTransactionRepo := &testmock.PostgresTransactionRepositoryInterface{}
			dbTransactionRepo.On("StartTransactionQuery", mock.Anything).Return(&sqlx.Tx{}, tc.rStartTrxErr)
			dbTransactionRepo.On("CommitTransactionQuery", mock.Anything, mock.Anything).Return(tc.rCommitTrxErr)
			dbTransactionRepo.On("RollbackTransactionQuery", mock.Anything, mock.Anything).Return(nil)

			uc := usecase.NewProductUsecase(productRepo, dbTransactionRepo)
			_, err := uc.BulkReduceQtyProduct(tc.ctx, tc.tenant, tc.payload)
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

			uc := usecase.NewProductUsecase(productRepo, &testmock.PostgresTransactionRepositoryInterface{})
			_, err := uc.GetProductByID(tc.ctx, 123)
			assert.Equal(t, tc.wantErr, err != nil)
		})
	}
}

func TestGetProducts(t *testing.T) {
	testcases := []struct {
		name                 string
		ctx                  context.Context
		data                 *entity.GetProductPayload
		rGetProductsRes      []*entity.Product
		rGetProductsErr      error
		rGetProductsCountRes int
		rGetProductsCountErr error
		wantErr              bool
	}{
		{
			name:    "deadline context",
			ctx:     fixture.CtxEnded(),
			wantErr: true,
		},
		{
			name:            "failed to get products",
			ctx:             context.Background(),
			rGetProductsErr: errors.New("error get products"),
			wantErr:         true,
		},
		{
			name:                 "failed to get products count",
			ctx:                  context.Background(),
			rGetProductsCountErr: errors.New("error get products count"),
			wantErr:              true,
		},
		{
			name:    "success",
			ctx:     context.Background(),
			data:    &entity.GetProductPayload{Tenant: types.TenantLoremType},
			wantErr: false,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			productRepo := &testmock.ProductRepositoryInterface{}
			productRepo.On("GetProducts", mock.Anything, mock.Anything).Return(tc.rGetProductsRes, tc.rGetProductsErr)
			productRepo.On("GetProductsCount", mock.Anything, mock.Anything).Return(tc.rGetProductsCountRes, tc.rGetProductsCountErr)

			uc := usecase.NewProductUsecase(productRepo, &testmock.PostgresTransactionRepositoryInterface{})
			_, _, err := uc.GetProducts(tc.ctx, tc.data)
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
			productRepo.On("UpdateProduct", mock.Anything, mock.Anything, mock.Anything).Return(tc.rProductErr)

			uc := usecase.NewProductUsecase(productRepo, &testmock.PostgresTransactionRepositoryInterface{})
			_, err := uc.UpdateProduct(tc.ctx, tc.productID, tc.data)
			assert.Equal(t, tc.wantErr, err != nil)
		})
	}
}
