package v1_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/satriowisnugroho/catalog/internal/entity"
	httpv1 "github.com/satriowisnugroho/catalog/internal/handler/http/v1"
	"github.com/satriowisnugroho/catalog/internal/response"
	testmock "github.com/satriowisnugroho/catalog/test/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

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
			productUsecase.On("GetProductByID", mock.Anything, mock.Anything).Return(&entity.Product{}, tc.uProductErr)

			h := &httpv1.ProductHandler{l, pp, productUsecase}
			h.GetProductByID(ctx)

			assert.Equal(t, tc.httpStatusCodeRes, w.Code)
		})
	}
}
