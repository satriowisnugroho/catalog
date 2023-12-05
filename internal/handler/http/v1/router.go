package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginswagger "github.com/swaggo/gin-swagger"

	// Swagger docs.
	_ "github.com/satriowisnugroho/catalog/docs"
	"github.com/satriowisnugroho/catalog/internal/usecase"
	"github.com/satriowisnugroho/catalog/pkg/logger"
)

// NewRouter -.
// Swagger spec:
// @title       Catalog API
// @description An API Documentation
// @version     1.0
// @host        localhost:9999
// @BasePath    /v1
func NewRouter(handler *gin.Engine, l logger.LoggerInterface, p usecase.ProductUsecaseInterface) {
	// Options
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	// Swagger
	swaggerHandler := ginswagger.DisablingWrapHandler(swaggerfiles.Handler, "DISABLE_SWAGGER_HTTP_HANDLER")
	handler.GET("/swagger/*any", swaggerHandler)

	// K8s probe
	handler.GET("/healthz", func(c *gin.Context) { c.Status(http.StatusOK) })

	// Routers
	h := handler.Group("/v1")
	{
		newProductHandler(h, l, p)
	}
}
