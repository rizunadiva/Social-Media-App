// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	domain "socialmedia-app/domain"

	mock "github.com/stretchr/testify/mock"
)

// CommentsData is an autogenerated mock type for the CommentsData type
type CommentsData struct {
	mock.Mock
}

// Delete provides a mock function with given fields: IDComments
func (_m *CommentsData) Delete(IDComments int) bool {
	ret := _m.Called(IDComments)

	var r0 bool
	if rf, ok := ret.Get(0).(func(int) bool); ok {
		r0 = rf(IDComments)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// GetAll provides a mock function with given fields:
func (_m *CommentsData) GetAll() []domain.Comments {
	ret := _m.Called()

	var r0 []domain.Comments
	if rf, ok := ret.Get(0).(func() []domain.Comments); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Comments)
		}
	}

	return r0
}

// GetMy provides a mock function with given fields: IDUser
func (_m *CommentsData) GetMy(IDUser int) []domain.Comments {
	ret := _m.Called(IDUser)

	var r0 []domain.Comments
	if rf, ok := ret.Get(0).(func(int) []domain.Comments); ok {
		r0 = rf(IDUser)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Comments)
		}
	}

	return r0
}

// Insert provides a mock function with given fields: newComments
func (_m *CommentsData) Insert(newComments domain.Comments) domain.Comments {
	ret := _m.Called(newComments)

	var r0 domain.Comments
	if rf, ok := ret.Get(0).(func(domain.Comments) domain.Comments); ok {
		r0 = rf(newComments)
	} else {
		r0 = ret.Get(0).(domain.Comments)
	}

	return r0
}

type mockConstructorTestingTNewCommentsData interface {
	mock.TestingT
	Cleanup(func())
}

// NewCommentsData creates a new instance of CommentsData. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewCommentsData(t mockConstructorTestingTNewCommentsData) *CommentsData {
	mock := &CommentsData{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}