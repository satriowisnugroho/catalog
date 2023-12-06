package parser

import (
	"encoding/json"
	"io"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/satriowisnugroho/catalog/internal/entity"
	"github.com/satriowisnugroho/catalog/internal/entity/types"
	"github.com/satriowisnugroho/catalog/internal/helper"
	"github.com/satriowisnugroho/catalog/internal/response"
)

// ProductParserInterface holds interface that parse data for product
type ProductParserInterface interface {
	ParseProductPayload(body io.Reader) (*entity.ProductPayload, error)
	ParseGetProductPayload(c *gin.Context) *entity.GetProductPayload
	ParseBulkReduceQtyProductPayload(body io.Reader) (*entity.BulkReduceQtyProductPayload, error)
}

// ProductParser struct for product parser initialization
type ProductParser struct{}

// NewProductParser create product parser
func NewProductParser() *ProductParser {
	return &ProductParser{}
}

// ParseProductPayload parse request product
func (p *ProductParser) ParseProductPayload(body io.Reader) (*entity.ProductPayload, error) {
	functionName := "ProductParser.ParseProductPayload"

	var payload entity.ProductPayload
	if err := json.NewDecoder(body).Decode(&payload); err != nil {
		if customErr, ok := err.(response.CustomError); ok {
			return nil, customErr
		}

		return nil, errors.Wrap(err, functionName)
	}

	return &payload, nil
}

// ParseGetProductPayload parse request get products
func (p *ProductParser) ParseGetProductPayload(c *gin.Context) *entity.GetProductPayload {
	offset, _ := strconv.Atoi(c.Query("offset"))
	limit, _ := strconv.Atoi(c.Query("limit"))
	payload := &entity.GetProductPayload{
		SKU:          c.Query("sku"),
		TitleKeyword: c.Query("keyword"),
		Category:     types.CategoryTypeNameToValue[c.Query("category")],
		Condition:    types.ConditionTypeNameToValue[c.Query("condition")],
		Tenant:       helper.GetTenant(c),
		OrderBy:      c.Query("orderby"),
		Offset:       offset,
		Limit:        limit,
	}

	return payload
}

// ParseBulkReduceQtyProductPayload parse request bulk reduce qty product
func (p *ProductParser) ParseBulkReduceQtyProductPayload(body io.Reader) (*entity.BulkReduceQtyProductPayload, error) {
	functionName := "ProductParser.ParseBulkReduceQtyProductPayload"

	var payload entity.BulkReduceQtyProductPayload
	if err := json.NewDecoder(body).Decode(&payload); err != nil {
		if customErr, ok := err.(response.CustomError); ok {
			return nil, customErr
		}

		return nil, errors.Wrap(err, functionName)
	}

	return &payload, nil
}
