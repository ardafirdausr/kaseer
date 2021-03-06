// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	context "context"

	entity "github.com/ardafirdausr/kaseer/internal/entity"

	mock "github.com/stretchr/testify/mock"
)

// OrderRepository is an autogenerated mock type for the OrderRepository type
type OrderRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, param
func (_m *OrderRepository) Create(ctx context.Context, param entity.CreateOrderParam) (*entity.Order, error) {
	ret := _m.Called(ctx, param)

	var r0 *entity.Order
	if rf, ok := ret.Get(0).(func(context.Context, entity.CreateOrderParam) *entity.Order); ok {
		r0 = rf(ctx, param)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Order)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, entity.CreateOrderParam) error); ok {
		r1 = rf(ctx, param)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateOrderItems provides a mock function with given fields: ctx, orderId, items
func (_m *OrderRepository) CreateOrderItems(ctx context.Context, orderId int64, items []*entity.CreateOrderItemParam) error {
	ret := _m.Called(ctx, orderId, items)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, []*entity.CreateOrderItemParam) error); ok {
		r0 = rf(ctx, orderId, items)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAllOrders provides a mock function with given fields: ctx
func (_m *OrderRepository) GetAllOrders(ctx context.Context) ([]*entity.Order, error) {
	ret := _m.Called(ctx)

	var r0 []*entity.Order
	if rf, ok := ret.Get(0).(func(context.Context) []*entity.Order); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entity.Order)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAnnualIncome provides a mock function with given fields: ctx
func (_m *OrderRepository) GetAnnualIncome(ctx context.Context) ([]*entity.AnnualIncome, error) {
	ret := _m.Called(ctx)

	var r0 []*entity.AnnualIncome
	if rf, ok := ret.Get(0).(func(context.Context) []*entity.AnnualIncome); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entity.AnnualIncome)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDailyOrderCount provides a mock function with given fields: ctx
func (_m *OrderRepository) GetDailyOrderCount(ctx context.Context) (int, error) {
	ret := _m.Called(ctx)

	var r0 int
	if rf, ok := ret.Get(0).(func(context.Context) int); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLastDayIncome provides a mock function with given fields: ctx
func (_m *OrderRepository) GetLastDayIncome(ctx context.Context) (int, error) {
	ret := _m.Called(ctx)

	var r0 int
	if rf, ok := ret.Get(0).(func(context.Context) int); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLastMonthIncome provides a mock function with given fields: ctx
func (_m *OrderRepository) GetLastMonthIncome(ctx context.Context) (int, error) {
	ret := _m.Called(ctx)

	var r0 int
	if rf, ok := ret.Get(0).(func(context.Context) int); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetOrderItemsByID provides a mock function with given fields: ctx, ID
func (_m *OrderRepository) GetOrderItemsByID(ctx context.Context, ID int64) ([]*entity.OrderItem, error) {
	ret := _m.Called(ctx, ID)

	var r0 []*entity.OrderItem
	if rf, ok := ret.Get(0).(func(context.Context, int64) []*entity.OrderItem); ok {
		r0 = rf(ctx, ID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entity.OrderItem)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, ID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTotalOrderCount provides a mock function with given fields: ctx
func (_m *OrderRepository) GetTotalOrderCount(ctx context.Context) (int, error) {
	ret := _m.Called(ctx)

	var r0 int
	if rf, ok := ret.Get(0).(func(context.Context) int); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
