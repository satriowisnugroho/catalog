// Code generated by mockery v2.9.5. DO NOT EDIT.

package mocks

import (
	gin "github.com/gin-gonic/gin"
	entity "github.com/satriowisnugroho/catalog/internal/entity"

	io "io"

	mock "github.com/stretchr/testify/mock"
)

// ProductParserInterface is an autogenerated mock type for the ProductParserInterface type
type ProductParserInterface struct {
	mock.Mock
}

// ParseBulkReduceQtyProductPayload provides a mock function with given fields: body
func (_m *ProductParserInterface) ParseBulkReduceQtyProductPayload(body io.Reader) (*entity.BulkReduceQtyProductPayload, error) {
	ret := _m.Called(body)

	var r0 *entity.BulkReduceQtyProductPayload
	if rf, ok := ret.Get(0).(func(io.Reader) *entity.BulkReduceQtyProductPayload); ok {
		r0 = rf(body)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.BulkReduceQtyProductPayload)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(io.Reader) error); ok {
		r1 = rf(body)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ParseGetProductPayload provides a mock function with given fields: c
func (_m *ProductParserInterface) ParseGetProductPayload(c *gin.Context) *entity.GetProductPayload {
	ret := _m.Called(c)

	var r0 *entity.GetProductPayload
	if rf, ok := ret.Get(0).(func(*gin.Context) *entity.GetProductPayload); ok {
		r0 = rf(c)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.GetProductPayload)
		}
	}

	return r0
}

// ParseProductPayload provides a mock function with given fields: body
func (_m *ProductParserInterface) ParseProductPayload(body io.Reader) (*entity.ProductPayload, error) {
	ret := _m.Called(body)

	var r0 *entity.ProductPayload
	if rf, ok := ret.Get(0).(func(io.Reader) *entity.ProductPayload); ok {
		r0 = rf(body)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.ProductPayload)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(io.Reader) error); ok {
		r1 = rf(body)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
