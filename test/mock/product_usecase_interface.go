// Code generated by mockery v2.9.5. DO NOT EDIT.

package mocks

import (
	context "context"

	entity "github.com/satriowisnugroho/catalog/internal/entity"
	mock "github.com/stretchr/testify/mock"
)

// ProductUsecaseInterface is an autogenerated mock type for the ProductUsecaseInterface type
type ProductUsecaseInterface struct {
	mock.Mock
}

// CreateProduct provides a mock function with given fields: ctx, productPayload
func (_m *ProductUsecaseInterface) CreateProduct(ctx context.Context, productPayload *entity.ProductPayload) (*entity.Product, error) {
	ret := _m.Called(ctx, productPayload)

	var r0 *entity.Product
	if rf, ok := ret.Get(0).(func(context.Context, *entity.ProductPayload) *entity.Product); ok {
		r0 = rf(ctx, productPayload)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Product)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *entity.ProductPayload) error); ok {
		r1 = rf(ctx, productPayload)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetProductByID provides a mock function with given fields: ctx, productID
func (_m *ProductUsecaseInterface) GetProductByID(ctx context.Context, productID int64) (*entity.Product, error) {
	ret := _m.Called(ctx, productID)

	var r0 *entity.Product
	if rf, ok := ret.Get(0).(func(context.Context, int64) *entity.Product); ok {
		r0 = rf(ctx, productID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Product)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, productID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}