// Code generated by mockery v2.49.2. DO NOT EDIT.

package mocks

import (
	context "context"
	dto "user_service/internal/application/dto"

	mock "github.com/stretchr/testify/mock"
)

// UserService is an autogenerated mock type for the UserService type
type UserService struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, user
func (_m *UserService) Create(ctx context.Context, user *dto.CreateUser) (*dto.User, error) {
	ret := _m.Called(ctx, user)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 *dto.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *dto.CreateUser) (*dto.User, error)); ok {
		return rf(ctx, user)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *dto.CreateUser) *dto.User); ok {
		r0 = rf(ctx, user)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dto.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *dto.CreateUser) error); ok {
		r1 = rf(ctx, user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Follow provides a mock function with given fields: ctx, id, followerID
func (_m *UserService) Follow(ctx context.Context, id string, followerID string) error {
	ret := _m.Called(ctx, id, followerID)

	if len(ret) == 0 {
		panic("no return value specified for Follow")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, id, followerID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Followers provides a mock function with given fields: ctx, id, page, limit
func (_m *UserService) Followers(ctx context.Context, id string, page int, limit int) ([]dto.Follower, error) {
	ret := _m.Called(ctx, id, page, limit)

	if len(ret) == 0 {
		panic("no return value specified for Followers")
	}

	var r0 []dto.Follower
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, int, int) ([]dto.Follower, error)); ok {
		return rf(ctx, id, page, limit)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, int, int) []dto.Follower); ok {
		r0 = rf(ctx, id, page, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]dto.Follower)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, int, int) error); ok {
		r1 = rf(ctx, id, page, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Following provides a mock function with given fields: ctx, id, page, limit
func (_m *UserService) Following(ctx context.Context, id string, page int, limit int) ([]dto.Follower, error) {
	ret := _m.Called(ctx, id, page, limit)

	if len(ret) == 0 {
		panic("no return value specified for Following")
	}

	var r0 []dto.Follower
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, int, int) ([]dto.Follower, error)); ok {
		return rf(ctx, id, page, limit)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, int, int) []dto.Follower); ok {
		r0 = rf(ctx, id, page, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]dto.Follower)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, int, int) error); ok {
		r1 = rf(ctx, id, page, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetById provides a mock function with given fields: ctx, id
func (_m *UserService) GetById(ctx context.Context, id string) (*dto.User, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for GetById")
	}

	var r0 *dto.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*dto.User, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *dto.User); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dto.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Paginate provides a mock function with given fields: ctx, page, limit
func (_m *UserService) Paginate(ctx context.Context, page int, limit int) ([]dto.User, error) {
	ret := _m.Called(ctx, page, limit)

	if len(ret) == 0 {
		panic("no return value specified for Paginate")
	}

	var r0 []dto.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int, int) ([]dto.User, error)); ok {
		return rf(ctx, page, limit)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int, int) []dto.User); ok {
		r0 = rf(ctx, page, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]dto.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int, int) error); ok {
		r1 = rf(ctx, page, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Unfollow provides a mock function with given fields: ctx, id, followerID
func (_m *UserService) Unfollow(ctx context.Context, id string, followerID string) error {
	ret := _m.Called(ctx, id, followerID)

	if len(ret) == 0 {
		panic("no return value specified for Unfollow")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, id, followerID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Update provides a mock function with given fields: ctx, id, user
func (_m *UserService) Update(ctx context.Context, id string, user *dto.UpdateUser) (*dto.User, error) {
	ret := _m.Called(ctx, id, user)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 *dto.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, *dto.UpdateUser) (*dto.User, error)); ok {
		return rf(ctx, id, user)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, *dto.UpdateUser) *dto.User); ok {
		r0 = rf(ctx, id, user)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dto.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, *dto.UpdateUser) error); ok {
		r1 = rf(ctx, id, user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewUserService creates a new instance of UserService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUserService(t interface {
	mock.TestingT
	Cleanup(func())
}) *UserService {
	mock := &UserService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}