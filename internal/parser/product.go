package parser

import (
	"encoding/json"
	"io"

	"github.com/pkg/errors"
	"github.com/satriowisnugroho/catalog/internal/entity"
)

// ProductParserInterface holds interface that parse data for product
type ProductParserInterface interface {
	ParseProductPayload(body io.Reader) (*entity.ProductPayload, error)
}

// ProductParser struct for product parser initialization
type ProductParser struct{}

// NewProductParser create product parser
func NewProductParser() *ProductParser {
	return &ProductParser{}
}

// ParserProductPayload parse request product
func (p *ProductParser) ParseProductPayload(body io.Reader) (*entity.ProductPayload, error) {
	functionName := "ProductParser.ParserProductPayload"

	var productPayload entity.ProductPayload
	if err := json.NewDecoder(body).Decode(&productPayload); err != nil {
		return nil, errors.Wrap(err, functionName)
	}

	return &productPayload, nil
}
