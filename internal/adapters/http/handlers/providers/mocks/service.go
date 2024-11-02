// Code generated by mockery v2.42.1. DO NOT EDIT.

package mocks

import (
	context "context"
	model "magnifin/internal/app/model"

	mock "github.com/stretchr/testify/mock"
)

// Service is an autogenerated mock type for the Service type
type Service struct {
	mock.Mock
}

type Service_Expecter struct {
	mock *mock.Mock
}

func (_m *Service) EXPECT() *Service_Expecter {
	return &Service_Expecter{mock: &_m.Mock}
}

// List provides a mock function with given fields: ctx
func (_m *Service) List(ctx context.Context) ([]model.Provider, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for List")
	}

	var r0 []model.Provider
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]model.Provider, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []model.Provider); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Provider)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Service_List_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'List'
type Service_List_Call struct {
	*mock.Call
}

// List is a helper method to define mock.On call
//   - ctx context.Context
func (_e *Service_Expecter) List(ctx interface{}) *Service_List_Call {
	return &Service_List_Call{Call: _e.mock.On("List", ctx)}
}

func (_c *Service_List_Call) Run(run func(ctx context.Context)) *Service_List_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *Service_List_Call) Return(_a0 []model.Provider, _a1 error) *Service_List_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Service_List_Call) RunAndReturn(run func(context.Context) ([]model.Provider, error)) *Service_List_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: ctx, provider
func (_m *Service) Update(ctx context.Context, provider model.Provider) (*model.Provider, error) {
	ret := _m.Called(ctx, provider)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 *model.Provider
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, model.Provider) (*model.Provider, error)); ok {
		return rf(ctx, provider)
	}
	if rf, ok := ret.Get(0).(func(context.Context, model.Provider) *model.Provider); ok {
		r0 = rf(ctx, provider)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Provider)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, model.Provider) error); ok {
		r1 = rf(ctx, provider)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Service_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type Service_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - ctx context.Context
//   - provider model.Provider
func (_e *Service_Expecter) Update(ctx interface{}, provider interface{}) *Service_Update_Call {
	return &Service_Update_Call{Call: _e.mock.On("Update", ctx, provider)}
}

func (_c *Service_Update_Call) Run(run func(ctx context.Context, provider model.Provider)) *Service_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(model.Provider))
	})
	return _c
}

func (_c *Service_Update_Call) Return(_a0 *model.Provider, _a1 error) *Service_Update_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Service_Update_Call) RunAndReturn(run func(context.Context, model.Provider) (*model.Provider, error)) *Service_Update_Call {
	_c.Call.Return(run)
	return _c
}

// NewService creates a new instance of Service. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewService(t interface {
	mock.TestingT
	Cleanup(func())
}) *Service {
	mock := &Service{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
