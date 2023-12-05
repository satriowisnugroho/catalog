// Code generated by mockery v2.9.5. DO NOT EDIT.

package mocks

import (
	io "io"

	entity "github.com/satriowisnugroho/catalog/internal/entity"

	mock "github.com/stretchr/testify/mock"
)

// ProductParserInterface is an autogenerated mock type for the ProductParserInterface type
type ProductParserInterface struct {
	mock.Mock
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