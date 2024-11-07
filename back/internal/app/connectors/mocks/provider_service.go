// Code generated by mockery v2.42.1. DO NOT EDIT.

package mocks

import (
	context "context"
	model "magnifin/internal/app/model"

	mock "github.com/stretchr/testify/mock"
)

// ProviderService is an autogenerated mock type for the ProviderService type
type ProviderService struct {
	mock.Mock
}

type ProviderService_Expecter struct {
	mock *mock.Mock
}

func (_m *ProviderService) EXPECT() *ProviderService_Expecter {
	return &ProviderService_Expecter{mock: &_m.Mock}
}

// Connect provides a mock function with given fields: ctx, user, connector, params
func (_m *ProviderService) Connect(ctx context.Context, user *model.User, connector *model.Connector, params *model.ConnectParams) (*model.ConnectInstruction, error) {
	ret := _m.Called(ctx, user, connector, params)

	if len(ret) == 0 {
		panic("no return value specified for Connect")
	}

	var r0 *model.ConnectInstruction
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.User, *model.Connector, *model.ConnectParams) (*model.ConnectInstruction, error)); ok {
		return rf(ctx, user, connector, params)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *model.User, *model.Connector, *model.ConnectParams) *model.ConnectInstruction); ok {
		r0 = rf(ctx, user, connector, params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.ConnectInstruction)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *model.User, *model.Connector, *model.ConnectParams) error); ok {
		r1 = rf(ctx, user, connector, params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ProviderService_Connect_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Connect'
type ProviderService_Connect_Call struct {
	*mock.Call
}

// Connect is a helper method to define mock.On call
//   - ctx context.Context
//   - user *model.User
//   - connector *model.Connector
//   - params *model.ConnectParams
func (_e *ProviderService_Expecter) Connect(ctx interface{}, user interface{}, connector interface{}, params interface{}) *ProviderService_Connect_Call {
	return &ProviderService_Connect_Call{Call: _e.mock.On("Connect", ctx, user, connector, params)}
}

func (_c *ProviderService_Connect_Call) Run(run func(ctx context.Context, user *model.User, connector *model.Connector, params *model.ConnectParams)) *ProviderService_Connect_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*model.User), args[2].(*model.Connector), args[3].(*model.ConnectParams))
	})
	return _c
}

func (_c *ProviderService_Connect_Call) Return(_a0 *model.ConnectInstruction, _a1 error) *ProviderService_Connect_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ProviderService_Connect_Call) RunAndReturn(run func(context.Context, *model.User, *model.Connector, *model.ConnectParams) (*model.ConnectInstruction, error)) *ProviderService_Connect_Call {
	_c.Call.Return(run)
	return _c
}

// ConnectCallback provides a mock function with given fields: ctx, connector, connectionID, providerConnectionID
func (_m *ProviderService) ConnectCallback(ctx context.Context, connector *model.Connector, connectionID string, providerConnectionID *string) error {
	ret := _m.Called(ctx, connector, connectionID, providerConnectionID)

	if len(ret) == 0 {
		panic("no return value specified for ConnectCallback")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.Connector, string, *string) error); ok {
		r0 = rf(ctx, connector, connectionID, providerConnectionID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ProviderService_ConnectCallback_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ConnectCallback'
type ProviderService_ConnectCallback_Call struct {
	*mock.Call
}

// ConnectCallback is a helper method to define mock.On call
//   - ctx context.Context
//   - connector *model.Connector
//   - connectionID string
//   - providerConnectionID *string
func (_e *ProviderService_Expecter) ConnectCallback(ctx interface{}, connector interface{}, connectionID interface{}, providerConnectionID interface{}) *ProviderService_ConnectCallback_Call {
	return &ProviderService_ConnectCallback_Call{Call: _e.mock.On("ConnectCallback", ctx, connector, connectionID, providerConnectionID)}
}

func (_c *ProviderService_ConnectCallback_Call) Run(run func(ctx context.Context, connector *model.Connector, connectionID string, providerConnectionID *string)) *ProviderService_ConnectCallback_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*model.Connector), args[2].(string), args[3].(*string))
	})
	return _c
}

func (_c *ProviderService_ConnectCallback_Call) Return(_a0 error) *ProviderService_ConnectCallback_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ProviderService_ConnectCallback_Call) RunAndReturn(run func(context.Context, *model.Connector, string, *string) error) *ProviderService_ConnectCallback_Call {
	_c.Call.Return(run)
	return _c
}

// NewProviderService creates a new instance of ProviderService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewProviderService(t interface {
	mock.TestingT
	Cleanup(func())
}) *ProviderService {
	mock := &ProviderService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
