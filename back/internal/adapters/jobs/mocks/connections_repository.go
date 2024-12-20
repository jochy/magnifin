// Code generated by mockery v2.42.1. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	model "magnifin/internal/app/model"
)

// ConnectionsRepository is an autogenerated mock type for the ConnectionsRepository type
type ConnectionsRepository struct {
	mock.Mock
}

type ConnectionsRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *ConnectionsRepository) EXPECT() *ConnectionsRepository_Expecter {
	return &ConnectionsRepository_Expecter{mock: &_m.Mock}
}

// ListConnectionsToSync provides a mock function with given fields: ctx
func (_m *ConnectionsRepository) ListConnectionsToSync(ctx context.Context) ([]model.Connection, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for ListConnectionsToSync")
	}

	var r0 []model.Connection
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]model.Connection, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []model.Connection); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Connection)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ConnectionsRepository_ListConnectionsToSync_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListConnectionsToSync'
type ConnectionsRepository_ListConnectionsToSync_Call struct {
	*mock.Call
}

// ListConnectionsToSync is a helper method to define mock.On call
//   - ctx context.Context
func (_e *ConnectionsRepository_Expecter) ListConnectionsToSync(ctx interface{}) *ConnectionsRepository_ListConnectionsToSync_Call {
	return &ConnectionsRepository_ListConnectionsToSync_Call{Call: _e.mock.On("ListConnectionsToSync", ctx)}
}

func (_c *ConnectionsRepository_ListConnectionsToSync_Call) Run(run func(ctx context.Context)) *ConnectionsRepository_ListConnectionsToSync_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *ConnectionsRepository_ListConnectionsToSync_Call) Return(_a0 []model.Connection, _a1 error) *ConnectionsRepository_ListConnectionsToSync_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ConnectionsRepository_ListConnectionsToSync_Call) RunAndReturn(run func(context.Context) ([]model.Connection, error)) *ConnectionsRepository_ListConnectionsToSync_Call {
	_c.Call.Return(run)
	return _c
}

// NewConnectionsRepository creates a new instance of ConnectionsRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewConnectionsRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *ConnectionsRepository {
	mock := &ConnectionsRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
