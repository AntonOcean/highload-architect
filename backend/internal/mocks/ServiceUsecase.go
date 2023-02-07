// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	context "context"
	domain "kek/internal/domain"

	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// ServiceUsecase is an autogenerated mock type for the ServiceUsecase type
type ServiceUsecase struct {
	mock.Mock
}

// AuthUser provides a mock function with given fields: ctx, userID, password
func (_m *ServiceUsecase) AuthUser(ctx context.Context, userID uuid.UUID, password string) (string, error) {
	ret := _m.Called(ctx, userID, password)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, string) string); ok {
		r0 = rf(ctx, userID, password)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID, string) error); ok {
		r1 = rf(ctx, userID, password)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateFriend provides a mock function with given fields: ctx, userID, friendID
func (_m *ServiceUsecase) CreateFriend(ctx context.Context, userID uuid.UUID, friendID uuid.UUID) error {
	ret := _m.Called(ctx, userID, friendID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, uuid.UUID) error); ok {
		r0 = rf(ctx, userID, friendID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateToken provides a mock function with given fields: ctx, userID
func (_m *ServiceUsecase) CreateToken(ctx context.Context, userID uuid.UUID) (string, error) {
	ret := _m.Called(ctx, userID)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) string); ok {
		r0 = rf(ctx, userID)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(ctx, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateUser provides a mock function with given fields: ctx, user
func (_m *ServiceUsecase) CreateUser(ctx context.Context, user *domain.User) error {
	ret := _m.Called(ctx, user)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.User) error); ok {
		r0 = rf(ctx, user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteFriend provides a mock function with given fields: ctx, userID, friendID
func (_m *ServiceUsecase) DeleteFriend(ctx context.Context, userID uuid.UUID, friendID uuid.UUID) error {
	ret := _m.Called(ctx, userID, friendID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, uuid.UUID) error); ok {
		r0 = rf(ctx, userID, friendID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetTokenData provides a mock function with given fields: ctx, token
func (_m *ServiceUsecase) GetTokenData(ctx context.Context, token string) (*domain.Claims, error) {
	ret := _m.Called(ctx, token)

	var r0 *domain.Claims
	if rf, ok := ret.Get(0).(func(context.Context, string) *domain.Claims); ok {
		r0 = rf(ctx, token)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Claims)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, token)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserByID provides a mock function with given fields: ctx, userID
func (_m *ServiceUsecase) GetUserByID(ctx context.Context, userID uuid.UUID) (*domain.User, error) {
	ret := _m.Called(ctx, userID)

	var r0 *domain.User
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) *domain.User); ok {
		r0 = rf(ctx, userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(ctx, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUsersByPrefix provides a mock function with given fields: ctx, firstName, lastName
func (_m *ServiceUsecase) GetUsersByPrefix(ctx context.Context, firstName string, lastName string) ([]*domain.User, error) {
	ret := _m.Called(ctx, firstName, lastName)

	var r0 []*domain.User
	if rf, ok := ret.Get(0).(func(context.Context, string, string) []*domain.User); ok {
		r0 = rf(ctx, firstName, lastName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*domain.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, firstName, lastName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewServiceUsecase interface {
	mock.TestingT
	Cleanup(func())
}

// NewServiceUsecase creates a new instance of ServiceUsecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewServiceUsecase(t mockConstructorTestingTNewServiceUsecase) *ServiceUsecase {
	mock := &ServiceUsecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
