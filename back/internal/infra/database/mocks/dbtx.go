// Code generated by mockery v2.42.1. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	sql "database/sql"
)

// DBTX is an autogenerated mock type for the DBTX type
type DBTX struct {
	mock.Mock
}

type DBTX_Expecter struct {
	mock *mock.Mock
}

func (_m *DBTX) EXPECT() *DBTX_Expecter {
	return &DBTX_Expecter{mock: &_m.Mock}
}

// ExecContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *DBTX) ExecContext(_a0 context.Context, _a1 string, _a2 ...interface{}) (sql.Result, error) {
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _a2...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ExecContext")
	}

	var r0 sql.Result
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, ...interface{}) (sql.Result, error)); ok {
		return rf(_a0, _a1, _a2...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, ...interface{}) sql.Result); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(sql.Result)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, ...interface{}) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DBTX_ExecContext_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ExecContext'
type DBTX_ExecContext_Call struct {
	*mock.Call
}

// ExecContext is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 string
//   - _a2 ...interface{}
func (_e *DBTX_Expecter) ExecContext(_a0 interface{}, _a1 interface{}, _a2 ...interface{}) *DBTX_ExecContext_Call {
	return &DBTX_ExecContext_Call{Call: _e.mock.On("ExecContext",
		append([]interface{}{_a0, _a1}, _a2...)...)}
}

func (_c *DBTX_ExecContext_Call) Run(run func(_a0 context.Context, _a1 string, _a2 ...interface{})) *DBTX_ExecContext_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(args[0].(context.Context), args[1].(string), variadicArgs...)
	})
	return _c
}

func (_c *DBTX_ExecContext_Call) Return(_a0 sql.Result, _a1 error) *DBTX_ExecContext_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *DBTX_ExecContext_Call) RunAndReturn(run func(context.Context, string, ...interface{}) (sql.Result, error)) *DBTX_ExecContext_Call {
	_c.Call.Return(run)
	return _c
}

// PrepareContext provides a mock function with given fields: _a0, _a1
func (_m *DBTX) PrepareContext(_a0 context.Context, _a1 string) (*sql.Stmt, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for PrepareContext")
	}

	var r0 *sql.Stmt
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*sql.Stmt, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *sql.Stmt); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sql.Stmt)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DBTX_PrepareContext_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PrepareContext'
type DBTX_PrepareContext_Call struct {
	*mock.Call
}

// PrepareContext is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 string
func (_e *DBTX_Expecter) PrepareContext(_a0 interface{}, _a1 interface{}) *DBTX_PrepareContext_Call {
	return &DBTX_PrepareContext_Call{Call: _e.mock.On("PrepareContext", _a0, _a1)}
}

func (_c *DBTX_PrepareContext_Call) Run(run func(_a0 context.Context, _a1 string)) *DBTX_PrepareContext_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *DBTX_PrepareContext_Call) Return(_a0 *sql.Stmt, _a1 error) *DBTX_PrepareContext_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *DBTX_PrepareContext_Call) RunAndReturn(run func(context.Context, string) (*sql.Stmt, error)) *DBTX_PrepareContext_Call {
	_c.Call.Return(run)
	return _c
}

// QueryContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *DBTX) QueryContext(_a0 context.Context, _a1 string, _a2 ...interface{}) (*sql.Rows, error) {
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _a2...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for QueryContext")
	}

	var r0 *sql.Rows
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, ...interface{}) (*sql.Rows, error)); ok {
		return rf(_a0, _a1, _a2...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, ...interface{}) *sql.Rows); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sql.Rows)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, ...interface{}) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DBTX_QueryContext_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'QueryContext'
type DBTX_QueryContext_Call struct {
	*mock.Call
}

// QueryContext is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 string
//   - _a2 ...interface{}
func (_e *DBTX_Expecter) QueryContext(_a0 interface{}, _a1 interface{}, _a2 ...interface{}) *DBTX_QueryContext_Call {
	return &DBTX_QueryContext_Call{Call: _e.mock.On("QueryContext",
		append([]interface{}{_a0, _a1}, _a2...)...)}
}

func (_c *DBTX_QueryContext_Call) Run(run func(_a0 context.Context, _a1 string, _a2 ...interface{})) *DBTX_QueryContext_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(args[0].(context.Context), args[1].(string), variadicArgs...)
	})
	return _c
}

func (_c *DBTX_QueryContext_Call) Return(_a0 *sql.Rows, _a1 error) *DBTX_QueryContext_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *DBTX_QueryContext_Call) RunAndReturn(run func(context.Context, string, ...interface{}) (*sql.Rows, error)) *DBTX_QueryContext_Call {
	_c.Call.Return(run)
	return _c
}

// QueryRowContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *DBTX) QueryRowContext(_a0 context.Context, _a1 string, _a2 ...interface{}) *sql.Row {
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _a2...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for QueryRowContext")
	}

	var r0 *sql.Row
	if rf, ok := ret.Get(0).(func(context.Context, string, ...interface{}) *sql.Row); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sql.Row)
		}
	}

	return r0
}

// DBTX_QueryRowContext_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'QueryRowContext'
type DBTX_QueryRowContext_Call struct {
	*mock.Call
}

// QueryRowContext is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 string
//   - _a2 ...interface{}
func (_e *DBTX_Expecter) QueryRowContext(_a0 interface{}, _a1 interface{}, _a2 ...interface{}) *DBTX_QueryRowContext_Call {
	return &DBTX_QueryRowContext_Call{Call: _e.mock.On("QueryRowContext",
		append([]interface{}{_a0, _a1}, _a2...)...)}
}

func (_c *DBTX_QueryRowContext_Call) Run(run func(_a0 context.Context, _a1 string, _a2 ...interface{})) *DBTX_QueryRowContext_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(args[0].(context.Context), args[1].(string), variadicArgs...)
	})
	return _c
}

func (_c *DBTX_QueryRowContext_Call) Return(_a0 *sql.Row) *DBTX_QueryRowContext_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *DBTX_QueryRowContext_Call) RunAndReturn(run func(context.Context, string, ...interface{}) *sql.Row) *DBTX_QueryRowContext_Call {
	_c.Call.Return(run)
	return _c
}

// NewDBTX creates a new instance of DBTX. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewDBTX(t interface {
	mock.TestingT
	Cleanup(func())
}) *DBTX {
	mock := &DBTX{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}