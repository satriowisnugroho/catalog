package v1

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/satriowisnugroho/catalog/internal/entity"
	"github.com/satriowisnugroho/catalog/internal/response"
	"github.com/satriowisnugroho/catalog/internal/usecase"
	"github.com/satriowisnugroho/catalog/pkg/logger"
)

type ProductHandler struct {
	Logger         logger.LoggerInterface
	ProductUsecase usecase.ProductUsecaseInterface
}

func newProductHandler(handler *gin.RouterGroup, l logger.LoggerInterface, pu usecase.ProductUsecaseInterface) {
	r := &ProductHandler{l, pu}

	h := handler.Group("/products")
	{
		h.POST("/", r.CreateProduct)
		h.GET("/:id", r.GetProductByID)
	}
}

// @Summary     Create Product
// @Description An API to create product
// @ID          create
// @Tags  	    product
// @Accept      json
// @Produce     json
// @Success     200 {object} response.SuccessBody{data=entity.Product,meta=response.MetaInfo}
// @Failure     404 {object} response.ErrorBody
// @Failure     500 {object} response.ErrorBody
// @Router      /products [post]
func (h *ProductHandler) CreateProduct(c *gin.Context) {
	product := &entity.Product{}
	err := h.ProductUsecase.CreateProduct(c.Request.Context(), product)
	if err != nil {
		h.Logger.Error(err, "http - v1 - CreateProduct")
		response.Error(c, err)

		return
	}

	response.OK(c, product, "")
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
	productID, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	product, err := h.ProductUsecase.GetProductByID(c.Request.Context(), productID)
	if err != nil {
		h.Logger.Error(err, "http - v1 - GetProductByID")
		response.Error(c, err)

		return
	}

	response.OK(c, product, "")
}
