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

// GetProductBySKU provides a mock function with given fields: ctx, productSKU
func (_m *ProductRepositoryInterface) GetProductBySKU(ctx context.Context, productSKU string) (*entity.Product, error) {
	ret := _m.Called(ctx, productSKU)

	var r0 *entity.Product
	if rf, ok := ret.Get(0).(func(context.Context, string) *entity.Product); ok {
		r0 = rf(ctx, productSKU)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Product)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, productSKU)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetProducts provides a mock function with given fields: ctx, payload
func (_m *ProductRepositoryInterface) GetProducts(ctx context.Context, payload *entity.GetProductPayload) ([]*entity.Product, error) {
	ret := _m.Called(ctx, payload)

	var r0 []*entity.Product
	if rf, ok := ret.Get(0).(func(context.Context, *entity.GetProductPayload) []*entity.Product); ok {
		r0 = rf(ctx, payload)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entity.Product)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *entity.GetProductPayload) error); ok {
		r1 = rf(ctx, payload)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetProductsCount provides a mock function with given fields: ctx, payload
func (_m *ProductRepositoryInterface) GetProductsCount(ctx context.Context, payload *entity.GetProductPayload) (int, error) {
	ret := _m.Called(ctx, payload)

	var r0 int
	if rf, ok := ret.Get(0).(func(context.Context, *entity.GetProductPayload) int); ok {
		r0 = rf(ctx, payload)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *entity.GetProductPayload) error); ok {
		r1 = rf(ctx, payload)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateProduct provides a mock function with given fields: ctx, dbTrx, product
func (_m *ProductRepositoryInterface) UpdateProduct(ctx context.Context, dbTrx interface{}, product *entity.Product) error {
	ret := _m.Called(ctx, dbTrx, product)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, interface{}, *entity.Product) error); ok {
		r0 = rf(ctx, dbTrx, product)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
