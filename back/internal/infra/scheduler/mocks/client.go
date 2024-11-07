// Code generated by mockery v2.42.1. DO NOT EDIT.

package mocks

import (
	context "context"

	river "github.com/riverqueue/river"
	mock "github.com/stretchr/testify/mock"
)

// Client is an autogenerated mock type for the Client type
type Client struct {
	mock.Mock
}

type Client_Expecter struct {
	mock *mock.Mock
}

func (_m *Client) EXPECT() *Client_Expecter {
	return &Client_Expecter{mock: &_m.Mock}
}

// Start provides a mock function with given fields: ctx
func (_m *Client) Start(ctx context.Context) error {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for Start")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Client_Start_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Start'
type Client_Start_Call struct {
	*mock.Call
}

// Start is a helper method to define mock.On call
//   - ctx context.Context
func (_e *Client_Expecter) Start(ctx interface{}) *Client_Start_Call {
	return &Client_Start_Call{Call: _e.mock.On("Start", ctx)}
}

func (_c *Client_Start_Call) Run(run func(ctx context.Context)) *Client_Start_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *Client_Start_Call) Return(_a0 error) *Client_Start_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Client_Start_Call) RunAndReturn(run func(context.Context) error) *Client_Start_Call {
	_c.Call.Return(run)
	return _c
}

// Stop provides a mock function with given fields: ctx
func (_m *Client) Stop(ctx context.Context) error {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for Stop")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Client_Stop_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Stop'
type Client_Stop_Call struct {
	*mock.Call
}

// Stop is a helper method to define mock.On call
//   - ctx context.Context
func (_e *Client_Expecter) Stop(ctx interface{}) *Client_Stop_Call {
	return &Client_Stop_Call{Call: _e.mock.On("Stop", ctx)}
}

func (_c *Client_Stop_Call) Run(run func(ctx context.Context)) *Client_Stop_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *Client_Stop_Call) Return(_a0 error) *Client_Stop_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Client_Stop_Call) RunAndReturn(run func(context.Context) error) *Client_Stop_Call {
	_c.Call.Return(run)
	return _c
}

// Trigger provides a mock function with given fields: ctx, job
func (_m *Client) Trigger(ctx context.Context, job river.JobArgs) error {
	ret := _m.Called(ctx, job)

	if len(ret) == 0 {
		panic("no return value specified for Trigger")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, river.JobArgs) error); ok {
		r0 = rf(ctx, job)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Client_Trigger_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Trigger'
type Client_Trigger_Call struct {
	*mock.Call
}

// Trigger is a helper method to define mock.On call
//   - ctx context.Context
//   - job river.JobArgs
func (_e *Client_Expecter) Trigger(ctx interface{}, job interface{}) *Client_Trigger_Call {
	return &Client_Trigger_Call{Call: _e.mock.On("Trigger", ctx, job)}
}

func (_c *Client_Trigger_Call) Run(run func(ctx context.Context, job river.JobArgs)) *Client_Trigger_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(river.JobArgs))
	})
	return _c
}

func (_c *Client_Trigger_Call) Return(_a0 error) *Client_Trigger_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Client_Trigger_Call) RunAndReturn(run func(context.Context, river.JobArgs) error) *Client_Trigger_Call {
	_c.Call.Return(run)
	return _c
}

// NewClient creates a new instance of Client. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *Client {
	mock := &Client{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
