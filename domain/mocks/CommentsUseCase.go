// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	domain "socialmedia-app/domain"

	mock "github.com/stretchr/testify/mock"
)

// CommentsUseCase is an autogenerated mock type for the CommentsUseCase type
type CommentsUseCase struct {
	mock.Mock
}

// AddComments provides a mock function with given fields: IDUser, newComments
func (_m *CommentsUseCase) AddComments(IDUser int, newComments domain.Comments) (domain.Comments, error) {
	ret := _m.Called(IDUser, newComments)

	var r0 domain.Comments
	if rf, ok := ret.Get(0).(func(int, domain.Comments) domain.Comments); ok {
		r0 = rf(IDUser, newComments)
	} else {
		r0 = ret.Get(0).(domain.Comments)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, domain.Comments) error); ok {
		r1 = rf(IDUser, newComments)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DelComments provides a mock function with given fields: IDComments
func (_m *CommentsUseCase) DelComments(IDComments int) (bool, error) {
	ret := _m.Called(IDComments)

	var r0 bool
	if rf, ok := ret.Get(0).(func(int) bool); ok {
		r0 = rf(IDComments)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(IDComments)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllC provides a mock function with given fields:
func (_m *CommentsUseCase) GetAllC() ([]domain.Comments, error) {
	ret := _m.Called()

	var r0 []domain.Comments
	if rf, ok := ret.Get(0).(func() []domain.Comments); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Comments)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetMyC provides a mock function with given fields: IDUser
func (_m *CommentsUseCase) GetMyC(IDUser int) ([]domain.Comments, error) {
	ret := _m.Called(IDUser)

	var r0 []domain.Comments
	if rf, ok := ret.Get(0).(func(int) []domain.Comments); ok {
		r0 = rf(IDUser)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Comments)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(IDUser)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewCommentsUseCase interface {
	mock.TestingT
	Cleanup(func())
}

// NewCommentsUseCase creates a new instance of CommentsUseCase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewCommentsUseCase(t mockConstructorTestingTNewCommentsUseCase) *CommentsUseCase {
	mock := &CommentsUseCase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
