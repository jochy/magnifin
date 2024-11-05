// Code generated by mockery v2.42.1. DO NOT EDIT.

package mocks

import (
	context "context"
	model "magnifin/internal/app/model"

	mock "github.com/stretchr/testify/mock"
)

// ConnectorRepository is an autogenerated mock type for the ConnectorRepository type
type ConnectorRepository struct {
	mock.Mock
}

type ConnectorRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *ConnectorRepository) EXPECT() *ConnectorRepository_Expecter {
	return &ConnectorRepository_Expecter{mock: &_m.Mock}
}

// Upsert provides a mock function with given fields: ctx, connectors
func (_m *ConnectorRepository) Upsert(ctx context.Context, connectors *model.Connector) (*model.Connector, error) {
	ret := _m.Called(ctx, connectors)

	if len(ret) == 0 {
		panic("no return value specified for Upsert")
	}

	var r0 *model.Connector
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.Connector) (*model.Connector, error)); ok {
		return rf(ctx, connectors)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *model.Connector) *model.Connector); ok {
		r0 = rf(ctx, connectors)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Connector)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *model.Connector) error); ok {
		r1 = rf(ctx, connectors)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ConnectorRepository_Upsert_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Upsert'
type ConnectorRepository_Upsert_Call struct {
	*mock.Call
}

// Upsert is a helper method to define mock.On call
//   - ctx context.Context
//   - connectors *model.Connector
func (_e *ConnectorRepository_Expecter) Upsert(ctx interface{}, connectors interface{}) *ConnectorRepository_Upsert_Call {
	return &ConnectorRepository_Upsert_Call{Call: _e.mock.On("Upsert", ctx, connectors)}
}

func (_c *ConnectorRepository_Upsert_Call) Run(run func(ctx context.Context, connectors *model.Connector)) *ConnectorRepository_Upsert_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*model.Connector))
	})
	return _c
}

func (_c *ConnectorRepository_Upsert_Call) Return(_a0 *model.Connector, _a1 error) *ConnectorRepository_Upsert_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ConnectorRepository_Upsert_Call) RunAndReturn(run func(context.Context, *model.Connector) (*model.Connector, error)) *ConnectorRepository_Upsert_Call {
	_c.Call.Return(run)
	return _c
}

// NewConnectorRepository creates a new instance of ConnectorRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewConnectorRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *ConnectorRepository {
	mock := &ConnectorRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}