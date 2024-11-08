// Code generated by mockery v2.42.1. DO NOT EDIT.

package mocks

import (
	context "context"
	model "magnifin/internal/app/model"

	mock "github.com/stretchr/testify/mock"
)

// ProviderUserRepository is an autogenerated mock type for the ProviderUserRepository type
type ProviderUserRepository struct {
	mock.Mock
}

type ProviderUserRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *ProviderUserRepository) EXPECT() *ProviderUserRepository_Expecter {
	return &ProviderUserRepository_Expecter{mock: &_m.Mock}
}

// GetByID provides a mock function with given fields: ctx, id
func (_m *ProviderUserRepository) GetByID(ctx context.Context, id int32) (*model.ProviderUser, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for GetByID")
	}

	var r0 *model.ProviderUser
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int32) (*model.ProviderUser, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int32) *model.ProviderUser); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.ProviderUser)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int32) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ProviderUserRepository_GetByID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByID'
type ProviderUserRepository_GetByID_Call struct {
	*mock.Call
}

// GetByID is a helper method to define mock.On call
//   - ctx context.Context
//   - id int32
func (_e *ProviderUserRepository_Expecter) GetByID(ctx interface{}, id interface{}) *ProviderUserRepository_GetByID_Call {
	return &ProviderUserRepository_GetByID_Call{Call: _e.mock.On("GetByID", ctx, id)}
}

func (_c *ProviderUserRepository_GetByID_Call) Run(run func(ctx context.Context, id int32)) *ProviderUserRepository_GetByID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int32))
	})
	return _c
}

func (_c *ProviderUserRepository_GetByID_Call) Return(_a0 *model.ProviderUser, _a1 error) *ProviderUserRepository_GetByID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ProviderUserRepository_GetByID_Call) RunAndReturn(run func(context.Context, int32) (*model.ProviderUser, error)) *ProviderUserRepository_GetByID_Call {
	_c.Call.Return(run)
	return _c
}

// GetByProviderIDAndUserID provides a mock function with given fields: ctx, providerID, userID
func (_m *ProviderUserRepository) GetByProviderIDAndUserID(ctx context.Context, providerID int32, userID int32) (*model.ProviderUser, error) {
	ret := _m.Called(ctx, providerID, userID)

	if len(ret) == 0 {
		panic("no return value specified for GetByProviderIDAndUserID")
	}

	var r0 *model.ProviderUser
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int32, int32) (*model.ProviderUser, error)); ok {
		return rf(ctx, providerID, userID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int32, int32) *model.ProviderUser); ok {
		r0 = rf(ctx, providerID, userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.ProviderUser)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int32, int32) error); ok {
		r1 = rf(ctx, providerID, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ProviderUserRepository_GetByProviderIDAndUserID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByProviderIDAndUserID'
type ProviderUserRepository_GetByProviderIDAndUserID_Call struct {
	*mock.Call
}

// GetByProviderIDAndUserID is a helper method to define mock.On call
//   - ctx context.Context
//   - providerID int32
//   - userID int32
func (_e *ProviderUserRepository_Expecter) GetByProviderIDAndUserID(ctx interface{}, providerID interface{}, userID interface{}) *ProviderUserRepository_GetByProviderIDAndUserID_Call {
	return &ProviderUserRepository_GetByProviderIDAndUserID_Call{Call: _e.mock.On("GetByProviderIDAndUserID", ctx, providerID, userID)}
}

func (_c *ProviderUserRepository_GetByProviderIDAndUserID_Call) Run(run func(ctx context.Context, providerID int32, userID int32)) *ProviderUserRepository_GetByProviderIDAndUserID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int32), args[2].(int32))
	})
	return _c
}

func (_c *ProviderUserRepository_GetByProviderIDAndUserID_Call) Return(_a0 *model.ProviderUser, _a1 error) *ProviderUserRepository_GetByProviderIDAndUserID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ProviderUserRepository_GetByProviderIDAndUserID_Call) RunAndReturn(run func(context.Context, int32, int32) (*model.ProviderUser, error)) *ProviderUserRepository_GetByProviderIDAndUserID_Call {
	_c.Call.Return(run)
	return _c
}

// Save provides a mock function with given fields: ctx, providerID, userID, providerUserID
func (_m *ProviderUserRepository) Save(ctx context.Context, providerID int32, userID int32, providerUserID string) (*model.ProviderUser, error) {
	ret := _m.Called(ctx, providerID, userID, providerUserID)

	if len(ret) == 0 {
		panic("no return value specified for Save")
	}

	var r0 *model.ProviderUser
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int32, int32, string) (*model.ProviderUser, error)); ok {
		return rf(ctx, providerID, userID, providerUserID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int32, int32, string) *model.ProviderUser); ok {
		r0 = rf(ctx, providerID, userID, providerUserID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.ProviderUser)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int32, int32, string) error); ok {
		r1 = rf(ctx, providerID, userID, providerUserID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ProviderUserRepository_Save_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Save'
type ProviderUserRepository_Save_Call struct {
	*mock.Call
}

// Save is a helper method to define mock.On call
//   - ctx context.Context
//   - providerID int32
//   - userID int32
//   - providerUserID string
func (_e *ProviderUserRepository_Expecter) Save(ctx interface{}, providerID interface{}, userID interface{}, providerUserID interface{}) *ProviderUserRepository_Save_Call {
	return &ProviderUserRepository_Save_Call{Call: _e.mock.On("Save", ctx, providerID, userID, providerUserID)}
}

func (_c *ProviderUserRepository_Save_Call) Run(run func(ctx context.Context, providerID int32, userID int32, providerUserID string)) *ProviderUserRepository_Save_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int32), args[2].(int32), args[3].(string))
	})
	return _c
}

func (_c *ProviderUserRepository_Save_Call) Return(_a0 *model.ProviderUser, _a1 error) *ProviderUserRepository_Save_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ProviderUserRepository_Save_Call) RunAndReturn(run func(context.Context, int32, int32, string) (*model.ProviderUser, error)) *ProviderUserRepository_Save_Call {
	_c.Call.Return(run)
	return _c
}

// NewProviderUserRepository creates a new instance of ProviderUserRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewProviderUserRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *ProviderUserRepository {
	mock := &ProviderUserRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
