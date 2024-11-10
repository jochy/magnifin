// Code generated by mockery v2.42.1. DO NOT EDIT.

package mocks

import (
	context "context"
	model "magnifin/internal/app/model"

	mock "github.com/stretchr/testify/mock"
)

// AccountRepository is an autogenerated mock type for the AccountRepository type
type AccountRepository struct {
	mock.Mock
}

type AccountRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *AccountRepository) EXPECT() *AccountRepository_Expecter {
	return &AccountRepository_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: ctx, account
func (_m *AccountRepository) Create(ctx context.Context, account *model.Account) (*model.Account, error) {
	ret := _m.Called(ctx, account)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 *model.Account
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.Account) (*model.Account, error)); ok {
		return rf(ctx, account)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *model.Account) *model.Account); ok {
		r0 = rf(ctx, account)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Account)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *model.Account) error); ok {
		r1 = rf(ctx, account)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AccountRepository_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type AccountRepository_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - ctx context.Context
//   - account *model.Account
func (_e *AccountRepository_Expecter) Create(ctx interface{}, account interface{}) *AccountRepository_Create_Call {
	return &AccountRepository_Create_Call{Call: _e.mock.On("Create", ctx, account)}
}

func (_c *AccountRepository_Create_Call) Run(run func(ctx context.Context, account *model.Account)) *AccountRepository_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*model.Account))
	})
	return _c
}

func (_c *AccountRepository_Create_Call) Return(_a0 *model.Account, _a1 error) *AccountRepository_Create_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *AccountRepository_Create_Call) RunAndReturn(run func(context.Context, *model.Account) (*model.Account, error)) *AccountRepository_Create_Call {
	_c.Call.Return(run)
	return _c
}

// GetByConnectionIDAndProviderAccountID provides a mock function with given fields: ctx, connectionID, providerAccountID
func (_m *AccountRepository) GetByConnectionIDAndProviderAccountID(ctx context.Context, connectionID int32, providerAccountID string) (*model.Account, error) {
	ret := _m.Called(ctx, connectionID, providerAccountID)

	if len(ret) == 0 {
		panic("no return value specified for GetByConnectionIDAndProviderAccountID")
	}

	var r0 *model.Account
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int32, string) (*model.Account, error)); ok {
		return rf(ctx, connectionID, providerAccountID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int32, string) *model.Account); ok {
		r0 = rf(ctx, connectionID, providerAccountID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Account)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int32, string) error); ok {
		r1 = rf(ctx, connectionID, providerAccountID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AccountRepository_GetByConnectionIDAndProviderAccountID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByConnectionIDAndProviderAccountID'
type AccountRepository_GetByConnectionIDAndProviderAccountID_Call struct {
	*mock.Call
}

// GetByConnectionIDAndProviderAccountID is a helper method to define mock.On call
//   - ctx context.Context
//   - connectionID int32
//   - providerAccountID string
func (_e *AccountRepository_Expecter) GetByConnectionIDAndProviderAccountID(ctx interface{}, connectionID interface{}, providerAccountID interface{}) *AccountRepository_GetByConnectionIDAndProviderAccountID_Call {
	return &AccountRepository_GetByConnectionIDAndProviderAccountID_Call{Call: _e.mock.On("GetByConnectionIDAndProviderAccountID", ctx, connectionID, providerAccountID)}
}

func (_c *AccountRepository_GetByConnectionIDAndProviderAccountID_Call) Run(run func(ctx context.Context, connectionID int32, providerAccountID string)) *AccountRepository_GetByConnectionIDAndProviderAccountID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int32), args[2].(string))
	})
	return _c
}

func (_c *AccountRepository_GetByConnectionIDAndProviderAccountID_Call) Return(_a0 *model.Account, _a1 error) *AccountRepository_GetByConnectionIDAndProviderAccountID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *AccountRepository_GetByConnectionIDAndProviderAccountID_Call) RunAndReturn(run func(context.Context, int32, string) (*model.Account, error)) *AccountRepository_GetByConnectionIDAndProviderAccountID_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: ctx, account
func (_m *AccountRepository) Update(ctx context.Context, account *model.Account) (*model.Account, error) {
	ret := _m.Called(ctx, account)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 *model.Account
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.Account) (*model.Account, error)); ok {
		return rf(ctx, account)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *model.Account) *model.Account); ok {
		r0 = rf(ctx, account)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Account)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *model.Account) error); ok {
		r1 = rf(ctx, account)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AccountRepository_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type AccountRepository_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - ctx context.Context
//   - account *model.Account
func (_e *AccountRepository_Expecter) Update(ctx interface{}, account interface{}) *AccountRepository_Update_Call {
	return &AccountRepository_Update_Call{Call: _e.mock.On("Update", ctx, account)}
}

func (_c *AccountRepository_Update_Call) Run(run func(ctx context.Context, account *model.Account)) *AccountRepository_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*model.Account))
	})
	return _c
}

func (_c *AccountRepository_Update_Call) Return(_a0 *model.Account, _a1 error) *AccountRepository_Update_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *AccountRepository_Update_Call) RunAndReturn(run func(context.Context, *model.Account) (*model.Account, error)) *AccountRepository_Update_Call {
	_c.Call.Return(run)
	return _c
}

// NewAccountRepository creates a new instance of AccountRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAccountRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *AccountRepository {
	mock := &AccountRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}