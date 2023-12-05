// Code generated by mockery v2.9.5. DO NOT EDIT.

package mocks

import (
	context "context"

	entity "github.com/satriowisnugroho/catalog/internal/entity"
	mock "github.com/stretchr/testify/mock"
)

// ProductRepositoryInterface is an autogenerated mock type for the ProductRepositoryInterface type
type ProductRepositoryInterface struct {
	mock.Mock
}

// CreateProduct provides a mock function with given fields: ctx, product
func (_m *ProductRepositoryInterface) CreateProduct(ctx context.Context, product *entity.Product) error {
	ret := _m.Called(ctx, product)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *entity.Product) error); ok {
		r0 = rf(ctx, product)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetProductByID provides a mock function with given fields: ctx, productID
func (_m *ProductRepositoryInterface) GetProductByID(ctx context.Context, productID int) (*entity.Product, error) {
	ret := _m.Called(ctx, productID)

	var r0 *entity.Product
	if rf, ok := ret.Get(0).(func(context.Context, int) *entity.Product); ok {
		r0 = rf(ctx, productID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Product)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, productID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
