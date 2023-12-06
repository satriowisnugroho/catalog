package parser

import (
	"encoding/json"
	"io"

	"github.com/pkg/errors"
	"github.com/satriowisnugroho/catalog/internal/entity"
	"github.com/satriowisnugroho/catalog/internal/response"
)

// ProductParserInterface holds interface that parse data for product
type ProductParserInterface interface {
	ParseProductPayload(body io.Reader) (*entity.ProductPayload, error)
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
		switch err.(type) {
		case response.CustomError:
			return nil, err
		}

		return nil, errors.Wrap(err, functionName)
	}

	return &payload, nil
}

// ParseBulkReduceQtyProductPayload parse request bulk reduce qty product
func (p *ProductParser) ParseBulkReduceQtyProductPayload(body io.Reader) (*entity.BulkReduceQtyProductPayload, error) {
	functionName := "ProductParser.ParseBulkReduceQtyProductPayload"

	var payload entity.BulkReduceQtyProductPayload
	if err := json.NewDecoder(body).Decode(&payload); err != nil {
		switch err.(type) {
		case response.CustomError:
			return nil, err
		}

		return nil, errors.Wrap(err, functionName)
	}

	return &payload, nil
}
