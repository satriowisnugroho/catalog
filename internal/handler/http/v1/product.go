package v1

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/satriowisnugroho/catalog/internal/entity"
	"github.com/satriowisnugroho/catalog/internal/entity/types"
	"github.com/satriowisnugroho/catalog/internal/parser"
	"github.com/satriowisnugroho/catalog/internal/response"
	"github.com/satriowisnugroho/catalog/internal/usecase"
	"github.com/satriowisnugroho/catalog/pkg/logger"
)

type ProductHandler struct {
	Logger         logger.LoggerInterface
	ProductParser  parser.ProductParserInterface
	ProductUsecase usecase.ProductUsecaseInterface
}

func newProductHandler(
	handler *gin.RouterGroup,
	l logger.LoggerInterface,
	pp parser.ProductParserInterface,
	pu usecase.ProductUsecaseInterface,
) {
	r := &ProductHandler{l, pp, pu}

	h := handler.Group("/products")
	{
		h.POST("/", r.CreateProduct)
		h.POST("/bulk-reduce-qty", r.BulkReduceQtyProduct)
		h.GET("/:id", r.GetProductByID)
		h.GET("/", r.GetProducts)
		h.PUT("/:id", r.UpdateProduct)
	}
}

// @Summary     Create Product
// @Description An API to create product
// @ID          create
// @Tags  	    product
// @Accept      json
// @Produce     json
// @Param       request body entity.SwaggerProductPayload true "Product Payload"
// @Success     200 {object} response.SuccessBody{data=entity.Product,meta=response.MetaInfo}
// @Failure     404 {object} response.ErrorBody
// @Failure     422 {object} response.ErrorBody
// @Failure     500 {object} response.ErrorBody
// @Router      /products [post]
func (h *ProductHandler) CreateProduct(c *gin.Context) {
	functionName := "ProductHandler.CreateProduct"

	payload, err := h.ProductParser.ParseProductPayload(c.Request.Body)
	if err != nil {
		switch err.(type) {
		case response.CustomError:
			response.Error(c, err)
			return
		}

		err = errors.Wrap(fmt.Errorf("h.productParser.ParseProductPayload: %w", err), functionName)
		h.Logger.Error(err)
		response.Error(c, err)
		return
	}

	product, err := h.ProductUsecase.CreateProduct(c.Request.Context(), payload)
	if err != nil {
		err = errors.Wrap(fmt.Errorf("h.ProductUsecase.CreateProduct: %w", err), functionName)
		h.Logger.Error(err)
		response.Error(c, err)

		return
	}

	response.OK(c, product, "")
}

// @Summary     Bulk Reduce Quantity Product
// @Description An API to bulk reduce quantity product
// @ID          bulk-reduce-qty
// @Tags  	    product
// @Accept      json
// @Produce     json
// @Param       request body entity.BulkReduceQtyProductPayload true "Bulk Reduce Qty Product Payload"
// @Success     200 {object} response.SuccessBody{data=entity.Product,meta=response.MetaInfo}
// @Failure     404 {object} response.ErrorBody
// @Failure     422 {object} response.ErrorBody
// @Failure     500 {object} response.ErrorBody
// @Router      /products/bulk-reduce-qty [post]
func (h *ProductHandler) BulkReduceQtyProduct(c *gin.Context) {
	functionName := "ProductHandler.BulkReduceQtyProduct"

	payload, err := h.ProductParser.ParseBulkReduceQtyProductPayload(c.Request.Body)
	if err != nil {
		err = errors.Wrap(fmt.Errorf("h.productParser.ParseBulkReduceQtyProductPayload: %w", err), functionName)
		h.Logger.Error(err)
		response.Error(c, err)
		return
	}

	if _, err = h.ProductUsecase.BulkReduceQtyProduct(c.Request.Context(), payload); err != nil {
		switch err.(type) {
		case response.CustomError:
			response.Error(c, err)
			return
		}

		err = errors.Wrap(fmt.Errorf("h.ProductUsecase.BulkReduceQtyProduct: %w", err), functionName)
		h.Logger.Error(err)
		response.Error(c, err)

		return
	}

	response.OK(c, nil, "Successfully bulk reduce quantity")
}

// @Summary     Show Product Detail
// @Description An API to show product detail
// @ID          detail
// @Tags  	    product
// @Accept      json
// @Param      	id path int true "Product ID"
// @Produce     json
// @Success     200 {object} response.SuccessBody{data=entity.Product,meta=response.MetaInfo}
// @Failure     404 {object} response.ErrorBody
// @Failure     500 {object} response.ErrorBody
// @Router      /products/{id} [get]
func (h *ProductHandler) GetProductByID(c *gin.Context) {
	productID, _ := strconv.Atoi(c.Param("id"))
	product, err := h.ProductUsecase.GetProductByID(c.Request.Context(), productID)
	if err != nil {
		h.Logger.Error(err, "http - v1 - GetProductByID")
		response.Error(c, err)

		return
	}

	response.OK(c, product, "")
}

// @Summary     Show Product List
// @Description An API to show product list
// @ID          list
// @Tags  	    product
// @Accept      json
// @Produce     json
// @Param       keyword query string false "title search by keyword"
// @Param       sku query string false "sku product"
// @Param       category query string false "category product"
// @Param       condition query integer false "condition product"
// @Param       tenant query string false "tenant product"
// @Param       orderby query string false "order by"
// @Param       offset query integer false "offset"
// @Param       limit query integer false "limit"
// @Success     200 {object} response.SuccessBody{data=[]entity.Product,meta=response.MetaInfo}
// @Failure     404 {object} response.ErrorBody
// @Failure     500 {object} response.ErrorBody
// @Router      /products [get]
func (h *ProductHandler) GetProducts(c *gin.Context) {
	condition, _ := strconv.Atoi(c.Query("condition"))
	offset, _ := strconv.Atoi(c.Query("offset"))
	limit, _ := strconv.Atoi(c.Query("limit"))
	payload := &entity.GetProductPayload{
		SKU:          c.Query("sku"),
		TitleKeyword: c.Query("keyword"),
		Category:     c.Query("category"),
		Condition:    int8(condition),
		// TODO: Adjust tenant payload
		Tenant:  types.TenantTypeNameToValue[c.Query("tenant")],
		OrderBy: c.Query("orderby"),
		Offset:  offset,
		Limit:   limit,
	}
	products, total, err := h.ProductUsecase.GetProducts(c.Request.Context(), payload)
	if err != nil {
		h.Logger.Error(err, "http - v1 - GetProducts")
		response.Error(c, err)

		return
	}

	response.OKWithPagination(c, products, "", total, payload.Offset, payload.Limit)
}

// @Summary     Update Product
// @Description An API to update product
// @ID          update
// @Tags  	    product
// @Param      	id path int true "Product ID"
// @Accept      json
// @Produce     json
// @Param       request body entity.SwaggerProductPayload true "Product Payload"
// @Success     200 {object} response.SuccessBody{data=entity.Product,meta=response.MetaInfo}
// @Failure     404 {object} response.ErrorBody
// @Failure     422 {object} response.ErrorBody
// @Failure     500 {object} response.ErrorBody
// @Router      /products/{id} [put]
func (h *ProductHandler) UpdateProduct(c *gin.Context) {
	functionName := "ProductHandler.UpdateProduct"

	payload, err := h.ProductParser.ParseProductPayload(c.Request.Body)
	if err != nil {
		switch err.(type) {
		case response.CustomError:
			response.Error(c, err)
			return
		}

		err = errors.Wrap(fmt.Errorf("h.productParser.ParseProductPayload: %w", err), functionName)
		h.Logger.Error(err)
		response.Error(c, err)
		return
	}

	productID, _ := strconv.Atoi(c.Param("id"))
	product, err := h.ProductUsecase.UpdateProduct(c.Request.Context(), productID, payload)
	if err != nil {
		err = errors.Wrap(fmt.Errorf("h.ProductUsecase.UpdateProduct: %w", err), functionName)
		h.Logger.Error(err)
		response.Error(c, err)

		return
	}

	response.OK(c, product, "")
}
