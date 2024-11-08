// Code generated by mockery v2.42.1. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	model "magnifin/internal/app/model"
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

// HandleSyncError provides a mock function with given fields: ctx, connectionID, syncErr
func (_m *Service) HandleSyncError(ctx context.Context, connectionID int32, syncErr error) error {
	ret := _m.Called(ctx, connectionID, syncErr)

	if len(ret) == 0 {
		panic("no return value specified for HandleSyncError")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int32, error) error); ok {
		r0 = rf(ctx, connectionID, syncErr)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Service_HandleSyncError_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'HandleSyncError'
type Service_HandleSyncError_Call struct {
	*mock.Call
}

// HandleSyncError is a helper method to define mock.On call
//   - ctx context.Context
//   - connectionID int32
//   - syncErr error
func (_e *Service_Expecter) HandleSyncError(ctx interface{}, connectionID interface{}, syncErr interface{}) *Service_HandleSyncError_Call {
	return &Service_HandleSyncError_Call{Call: _e.mock.On("HandleSyncError", ctx, connectionID, syncErr)}
}

func (_c *Service_HandleSyncError_Call) Run(run func(ctx context.Context, connectionID int32, syncErr error)) *Service_HandleSyncError_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int32), args[2].(error))
	})
	return _c
}

func (_c *Service_HandleSyncError_Call) Return(_a0 error) *Service_HandleSyncError_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Service_HandleSyncError_Call) RunAndReturn(run func(context.Context, int32, error) error) *Service_HandleSyncError_Call {
	_c.Call.Return(run)
	return _c
}

// SynchronizeConnection provides a mock function with given fields: ctx, connectionID
func (_m *Service) SynchronizeConnection(ctx context.Context, connectionID int32) error {
	ret := _m.Called(ctx, connectionID)

	if len(ret) == 0 {
		panic("no return value specified for SynchronizeConnection")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int32) error); ok {
		r0 = rf(ctx, connectionID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Service_SynchronizeConnection_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SynchronizeConnection'
type Service_SynchronizeConnection_Call struct {
	*mock.Call
}

// SynchronizeConnection is a helper method to define mock.On call
//   - ctx context.Context
//   - connectionID int32
func (_e *Service_Expecter) SynchronizeConnection(ctx interface{}, connectionID interface{}) *Service_SynchronizeConnection_Call {
	return &Service_SynchronizeConnection_Call{Call: _e.mock.On("SynchronizeConnection", ctx, connectionID)}
}

func (_c *Service_SynchronizeConnection_Call) Run(run func(ctx context.Context, connectionID int32)) *Service_SynchronizeConnection_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int32))
	})
	return _c
}

func (_c *Service_SynchronizeConnection_Call) Return(_a0 error) *Service_SynchronizeConnection_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Service_SynchronizeConnection_Call) RunAndReturn(run func(context.Context, int32) error) *Service_SynchronizeConnection_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateConnectorsList provides a mock function with given fields: ctx
func (_m *Service) UpdateConnectorsList(ctx context.Context) ([]model.Connector, []error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for UpdateConnectorsList")
	}

	var r0 []model.Connector
	var r1 []error
	if rf, ok := ret.Get(0).(func(context.Context) ([]model.Connector, []error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []model.Connector); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Connector)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) []error); ok {
		r1 = rf(ctx)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]error)
		}
	}

	return r0, r1
}

// Service_UpdateConnectorsList_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateConnectorsList'
type Service_UpdateConnectorsList_Call struct {
	*mock.Call
}

// UpdateConnectorsList is a helper method to define mock.On call
//   - ctx context.Context
func (_e *Service_Expecter) UpdateConnectorsList(ctx interface{}) *Service_UpdateConnectorsList_Call {
	return &Service_UpdateConnectorsList_Call{Call: _e.mock.On("UpdateConnectorsList", ctx)}
}

func (_c *Service_UpdateConnectorsList_Call) Run(run func(ctx context.Context)) *Service_UpdateConnectorsList_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *Service_UpdateConnectorsList_Call) Return(_a0 []model.Connector, _a1 []error) *Service_UpdateConnectorsList_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Service_UpdateConnectorsList_Call) RunAndReturn(run func(context.Context) ([]model.Connector, []error)) *Service_UpdateConnectorsList_Call {
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
