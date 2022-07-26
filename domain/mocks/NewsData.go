// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	domain "socialmedia-app/domain"

	mock "github.com/stretchr/testify/mock"
)

// NewsData is an autogenerated mock type for the NewsData type
type NewsData struct {
	mock.Mock
}

// Delete provides a mock function with given fields: IDNews
func (_m *NewsData) Delete(IDNews int) bool {
	ret := _m.Called(IDNews)

	var r0 bool
	if rf, ok := ret.Get(0).(func(int) bool); ok {
		r0 = rf(IDNews)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// GetAll provides a mock function with given fields:
func (_m *NewsData) GetAll() []domain.News {
	ret := _m.Called()

	var r0 []domain.News
	if rf, ok := ret.Get(0).(func() []domain.News); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.News)
		}
	}

	return r0
}

// GetMy provides a mock function with given fields: IDUser
func (_m *NewsData) GetMy(IDUser int) []domain.News {
	ret := _m.Called(IDUser)

	var r0 []domain.News
	if rf, ok := ret.Get(0).(func(int) []domain.News); ok {
		r0 = rf(IDUser)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.News)
		}
	}

	return r0
}

// Insert provides a mock function with given fields: newNews
func (_m *NewsData) Insert(newNews domain.News) domain.News {
	ret := _m.Called(newNews)

	var r0 domain.News
	if rf, ok := ret.Get(0).(func(domain.News) domain.News); ok {
		r0 = rf(newNews)
	} else {
		r0 = ret.Get(0).(domain.News)
	}

	return r0
}

// Update provides a mock function with given fields: IDNews, updatedNews
func (_m *NewsData) Update(IDNews int, updatedNews domain.News) domain.News {
	ret := _m.Called(IDNews, updatedNews)

	var r0 domain.News
	if rf, ok := ret.Get(0).(func(int, domain.News) domain.News); ok {
		r0 = rf(IDNews, updatedNews)
	} else {
		r0 = ret.Get(0).(domain.News)
	}

	return r0
}

type mockConstructorTestingTNewNewsData interface {
	mock.TestingT
	Cleanup(func())
}

// NewNewsData creates a new instance of NewsData. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewNewsData(t mockConstructorTestingTNewNewsData) *NewsData {
	mock := &NewsData{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
