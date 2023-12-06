// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate_swagger = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/products": {
            "get": {
                "description": "An API to show product list",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "product"
                ],
                "summary": "Show Product List",
                "operationId": "list",
                "parameters": [
                    {
                        "type": "string",
                        "default": "lorem",
                        "example": "lorem, ipsum",
                        "description": "Tenant Header",
                        "name": "X-Tenant",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "title search by keyword",
                        "name": "keyword",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "sku product",
                        "name": "sku",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "category product",
                        "name": "category",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "condition product",
                        "name": "condition",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "order by",
                        "name": "orderby",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "offset",
                        "name": "offset",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "limit",
                        "name": "limit",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.SuccessBody"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/entity.Product"
                                            }
                                        },
                                        "meta": {
                                            "$ref": "#/definitions/response.MetaInfo"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorBody"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorBody"
                        }
                    }
                }
            },
            "post": {
                "description": "An API to create product",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "product"
                ],
                "summary": "Create Product",
                "operationId": "create",
                "parameters": [
                    {
                        "description": "Product Payload",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entity.SwaggerProductPayload"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.SuccessBody"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/entity.Product"
                                        },
                                        "meta": {
                                            "$ref": "#/definitions/response.MetaInfo"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorBody"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorBody"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorBody"
                        }
                    }
                }
            }
        },
        "/products/bulk-reduce-qty": {
            "post": {
                "description": "An API to bulk reduce quantity product",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "product"
                ],
                "summary": "Bulk Reduce Quantity Product",
                "operationId": "bulk-reduce-qty",
                "parameters": [
                    {
                        "type": "string",
                        "default": "lorem",
                        "example": "lorem, ipsum",
                        "description": "Tenant Header",
                        "name": "X-Tenant",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "Payload",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entity.BulkReduceQtyProductPayload"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.SuccessBody"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/entity.Product"
                                        },
                                        "meta": {
                                            "$ref": "#/definitions/response.MetaInfo"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorBody"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorBody"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorBody"
                        }
                    }
                }
            }
        },
        "/products/{id}": {
            "get": {
                "description": "An API to show product detail",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "product"
                ],
                "summary": "Show Product Detail",
                "operationId": "detail",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Product ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.SuccessBody"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/entity.Product"
                                        },
                                        "meta": {
                                            "$ref": "#/definitions/response.MetaInfo"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorBody"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorBody"
                        }
                    }
                }
            },
            "put": {
                "description": "An API to update product",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "product"
                ],
                "summary": "Update Product",
                "operationId": "update",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Product ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Product Payload",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entity.SwaggerProductPayload"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.SuccessBody"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/entity.Product"
                                        },
                                        "meta": {
                                            "$ref": "#/definitions/response.MetaInfo"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorBody"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorBody"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorBody"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "entity.BulkReduceQtyProductItemPayload": {
            "type": "object",
            "properties": {
                "req_qty": {
                    "type": "integer"
                },
                "sku": {
                    "type": "string"
                }
            }
        },
        "entity.BulkReduceQtyProductPayload": {
            "type": "object",
            "properties": {
                "items": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entity.BulkReduceQtyProductItemPayload"
                    }
                }
            }
        },
        "entity.Product": {
            "type": "object",
            "properties": {
                "category": {
                    "type": "string"
                },
                "condition": {
                    "type": "integer"
                },
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "price": {
                    "type": "integer"
                },
                "qty": {
                    "type": "integer"
                },
                "sku": {
                    "type": "string"
                },
                "tenant": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "entity.SwaggerProductPayload": {
            "type": "object",
            "properties": {
                "category": {
                    "type": "string"
                },
                "condition": {
                    "type": "integer"
                },
                "price": {
                    "type": "integer"
                },
                "qty": {
                    "type": "integer"
                },
                "tenant": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "response.ErrorBody": {
            "type": "object",
            "properties": {
                "errors": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/response.ErrorInfo"
                    }
                },
                "meta": {}
            }
        },
        "response.ErrorInfo": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "field": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "response.MetaInfo": {
            "type": "object",
            "properties": {
                "http_status": {
                    "type": "integer"
                },
                "limit": {
                    "type": "integer"
                },
                "offset": {
                    "type": "integer"
                },
                "total": {
                    "type": "integer"
                }
            }
        },
        "response.SuccessBody": {
            "type": "object",
            "properties": {
                "data": {},
                "message": {
                    "type": "string"
                },
                "meta": {}
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:9999",
	BasePath:         "/v1",
	Schemes:          []string{},
	Title:            "Catalog API",
	Description:      "An API Documentation",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate_swagger,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
