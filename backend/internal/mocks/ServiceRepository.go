// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	context "context"
	domain "kek/internal/domain"

	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// ServiceRepository is an autogenerated mock type for the ServiceRepository type
type ServiceRepository struct {
	mock.Mock
}

// CreateUser provides a mock function with given fields: ctx, user
func (_m *ServiceRepository) CreateUser(ctx context.Context, user *domain.User) error {
	ret := _m.Called(ctx, user)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.User) error); ok {
		r0 = rf(ctx, user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetUserByID provides a mock function with given fields: ctx, userID
func (_m *ServiceRepository) GetUserByID(ctx context.Context, userID uuid.UUID) (*domain.User, error) {
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
func (_m *ServiceRepository) GetUsersByPrefix(ctx context.Context, firstName string, lastName string) ([]*domain.User, error) {
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

type mockConstructorTestingTNewServiceRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewServiceRepository creates a new instance of ServiceRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewServiceRepository(t mockConstructorTestingTNewServiceRepository) *ServiceRepository {
	mock := &ServiceRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
