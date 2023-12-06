package v1_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/satriowisnugroho/catalog/internal/entity"
	"github.com/satriowisnugroho/catalog/internal/entity/types"
	httpv1 "github.com/satriowisnugroho/catalog/internal/handler/http/v1"
	"github.com/satriowisnugroho/catalog/internal/response"
	testmock "github.com/satriowisnugroho/catalog/test/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateProduct(t *testing.T) {
	testcases := []struct {
		name              string
		pProductRes       *entity.ProductPayload
		pProductErr       error
		uProductRes       *entity.Product
		uProductErr       error
		httpStatusCodeRes int
	}{
		{
			name:              "parser return custom error",
			pProductErr:       response.ErrInvalidCategory,
			httpStatusCodeRes: http.StatusUnprocessableEntity,
		},
		{
			name:              "failed to parse product payload",
			pProductErr:       errors.New("error parse product payload"),
			httpStatusCodeRes: http.StatusInternalServerError,
		},
		{
			name:              "invalid type for tenant",
			pProductRes:       &entity.ProductPayload{},
			uProductErr:       response.ErrInvalidTenant,
			httpStatusCodeRes: http.StatusUnprocessableEntity,
		},
		{
			name:              "failed to create product",
			pProductRes:       &entity.ProductPayload{},
			uProductErr:       errors.New("error create product"),
			httpStatusCodeRes: http.StatusInternalServerError,
		},
		{
			name:              "success",
			pProductRes:       &entity.ProductPayload{Category: types.CategoryBookType, Condition: types.ConditionNewType, Tenant: types.TenantLoremType},
			uProductRes:       &entity.Product{Category: types.CategoryBookType, Condition: types.ConditionNewType, Tenant: types.TenantLoremType},
			httpStatusCodeRes: http.StatusOK,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)

			ctx.Request = &http.Request{
				Header: make(http.Header),
				Method: "POST",
			}

			l := &testmock.LoggerInterface{}
			l.On("Error", mock.Anything, mock.Anything)

			pp := &testmock.ProductParserInterface{}
			pp.On("ParseProductPayload", mock.Anything).Return(tc.pProductRes, tc.pProductErr)

			productUsecase := &testmock.ProductUsecaseInterface{}
			productUsecase.On("CreateProduct", mock.Anything, mock.Anything).Return(tc.uProductRes, tc.uProductErr)

			h := &httpv1.ProductHandler{l, pp, productUsecase}
			h.CreateProduct(ctx)

			assert.Equal(t, tc.httpStatusCodeRes, w.Code)
		})
	}
}

func TestBulkReduceQtyProduct(t *testing.T) {
	testcases := []struct {
		name              string
		pProductRes       *entity.BulkReduceQtyProductPayload
		pProductErr       error
		uProductRes       []*entity.Product
		uProductErr       error
		httpStatusCodeRes int
	}{
		{
			name:              "failed to parse payload",
			pProductErr:       errors.New("error parse payload"),
			httpStatusCodeRes: http.StatusInternalServerError,
		},
		{
			name:              "insufficient error",
			uProductErr:       response.ErrInsufficientStock,
			httpStatusCodeRes: http.StatusUnprocessableEntity,
		},
		{
			name:              "failed to bulk reduce qty product",
			uProductErr:       errors.New("error bulk reduce qty product"),
			httpStatusCodeRes: http.StatusInternalServerError,
		},
		{
			name:              "success",
			pProductRes:       &entity.BulkReduceQtyProductPayload{},
			uProductRes:       []*entity.Product{{Category: types.CategoryBookType, Condition: types.ConditionNewType, Tenant: types.TenantLoremType}},
			httpStatusCodeRes: http.StatusOK,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)

			ctx.Request = &http.Request{
				Header: make(http.Header),
				Method: "POST",
			}

			l := &testmock.LoggerInterface{}
			l.On("Error", mock.Anything, mock.Anything)

			pp := &testmock.ProductParserInterface{}
			pp.On("ParseBulkReduceQtyProductPayload", mock.Anything).Return(tc.pProductRes, tc.pProductErr)

			productUsecase := &testmock.ProductUsecaseInterface{}
			productUsecase.On("BulkReduceQtyProduct", mock.Anything, mock.Anything, mock.Anything).Return(tc.uProductRes, tc.uProductErr)

			h := &httpv1.ProductHandler{l, pp, productUsecase}
			h.BulkReduceQtyProduct(ctx)

			assert.Equal(t, tc.httpStatusCodeRes, w.Code)
		})
	}
}

func TestGetProductByID(t *testing.T) {
	testcases := []struct {
		name              string
		uProductErr       error
		httpStatusCodeRes int
	}{
		{
			name:              "product is not found",
			uProductErr:       response.ErrNotFound,
			httpStatusCodeRes: http.StatusNotFound,
		},
		{
			name:              "failed to get product",
			uProductErr:       errors.New("error get product"),
			httpStatusCodeRes: http.StatusInternalServerError,
		},
		{
			name:              "success",
			httpStatusCodeRes: http.StatusOK,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)

			ctx.Request = &http.Request{
				Header: make(http.Header),
				Method: "GET",
			}

			l := &testmock.LoggerInterface{}
			l.On("Error", mock.Anything, mock.Anything)

			pp := &testmock.ProductParserInterface{}

			productUsecase := &testmock.ProductUsecaseInterface{}
			productUsecase.On("GetProductByID", mock.Anything, mock.Anything, mock.Anything).Return(&entity.Product{Category: types.CategoryBookType, Condition: types.ConditionNewType, Tenant: types.TenantLoremType}, tc.uProductErr)

			h := &httpv1.ProductHandler{l, pp, productUsecase}
			h.GetProductByID(ctx)

			assert.Equal(t, tc.httpStatusCodeRes, w.Code)
		})
	}
}

func TestGetProducts(t *testing.T) {
	testcases := []struct {
		name              string
		uProductErr       error
		httpStatusCodeRes int
	}{
		{
			name:              "failed to get products",
			uProductErr:       errors.New("error get products"),
			httpStatusCodeRes: http.StatusInternalServerError,
		},
		{
			name:              "success",
			httpStatusCodeRes: http.StatusOK,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)

			ctx.Request, _ = http.NewRequest("GET", "/products?", nil)

			l := &testmock.LoggerInterface{}
			l.On("Error", mock.Anything, mock.Anything)

			pp := &testmock.ProductParserInterface{}

			productUsecase := &testmock.ProductUsecaseInterface{}
			productUsecase.On("GetProducts", mock.Anything, mock.Anything).Return([]*entity.Product{{Category: types.CategoryBookType, Condition: types.ConditionNewType, Tenant: types.TenantLoremType}}, 10, tc.uProductErr)

			h := &httpv1.ProductHandler{l, pp, productUsecase}
			h.GetProducts(ctx)

			assert.Equal(t, tc.httpStatusCodeRes, w.Code)
		})
	}
}

func TestUpdateProduct(t *testing.T) {
	testcases := []struct {
		name              string
		pProductRes       *entity.ProductPayload
		pProductErr       error
		uProductRes       *entity.Product
		uProductErr       error
		httpStatusCodeRes int
	}{
		{
			name:              "parser return custom error",
			pProductErr:       response.ErrInvalidCategory,
			httpStatusCodeRes: http.StatusUnprocessableEntity,
		},
		{
			name:              "failed to parse product payload",
			pProductErr:       errors.New("error parse product payload"),
			httpStatusCodeRes: http.StatusInternalServerError,
		},
		{
			name:              "forbidden",
			pProductRes:       &entity.ProductPayload{},
			uProductErr:       response.ErrForbidden,
			httpStatusCodeRes: http.StatusForbidden,
		},
		{
			name:              "failed to update product",
			pProductRes:       &entity.ProductPayload{},
			uProductErr:       errors.New("error update product"),
			httpStatusCodeRes: http.StatusInternalServerError,
		},
		{
			name:              "success",
			pProductRes:       &entity.ProductPayload{Category: types.CategoryBookType, Condition: types.ConditionNewType, Tenant: types.TenantLoremType},
			uProductRes:       &entity.Product{Category: types.CategoryBookType, Condition: types.ConditionNewType, Tenant: types.TenantLoremType},
			httpStatusCodeRes: http.StatusOK,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)

			ctx.Request = &http.Request{
				Header: make(http.Header),
				Method: "PUT",
			}

			l := &testmock.LoggerInterface{}
			l.On("Error", mock.Anything, mock.Anything)

			pp := &testmock.ProductParserInterface{}
			pp.On("ParseProductPayload", mock.Anything).Return(tc.pProductRes, tc.pProductErr)

			productUsecase := &testmock.ProductUsecaseInterface{}
			productUsecase.On("UpdateProduct", mock.Anything, mock.Anything, mock.Anything).Return(tc.uProductRes, tc.uProductErr)

			h := &httpv1.ProductHandler{l, pp, productUsecase}
			h.UpdateProduct(ctx)

			assert.Equal(t, tc.httpStatusCodeRes, w.Code)
		})
	}
}
