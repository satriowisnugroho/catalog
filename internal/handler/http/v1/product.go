package v1

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/satriowisnugroho/catalog/internal/response"
	"github.com/satriowisnugroho/catalog/internal/usecase"
	"github.com/satriowisnugroho/catalog/pkg/logger"

	// Import entity for swagger success response
	_ "github.com/satriowisnugroho/catalog/internal/entity"
)

type ProductHandler struct {
	Logger         logger.LoggerInterface
	ProductUsecase usecase.ProductUsecaseInterface
}

func newProductHandler(handler *gin.RouterGroup, l logger.LoggerInterface, pu usecase.ProductUsecaseInterface) {
	r := &ProductHandler{l, pu}

	h := handler.Group("/products")
	{
		h.GET("/:id", r.GetProductByID)
	}
}

// @Summary     Show detail
// @Description Show product detail
// @ID          detail
// @Tags  	    product
// @Accept      json
// @Param      	id path int true "Product ID"
// @Produce     json
// @Success     200 {object} response.SuccessBody{data=entity.Product,meta=response.MetaInfo}
// @Failure     404 {object} response.ErrorBody
// @Failure     500 {object} response.ErrorBody
// @Router      /products/{id} [get]
func (r *ProductHandler) GetProductByID(c *gin.Context) {
	productID, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	product, err := r.ProductUsecase.GetProductByID(c.Request.Context(), productID)
	if err != nil {
		r.Logger.Error(err, "http - v1 - getProductByID")
		response.Error(c, err)

		return
	}

	response.OK(c, product, "")
}
