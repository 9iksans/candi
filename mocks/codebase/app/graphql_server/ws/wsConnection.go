// Code generated by mockery v2.13.1. DO NOT EDIT.

package mocks

import (
	time "time"

	mock "github.com/stretchr/testify/mock"
)

// wsConnection is an autogenerated mock type for the wsConnection type
type wsConnection struct {
	mock.Mock
}

// Close provides a mock function with given fields:
func (_m *wsConnection) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ReadJSON provides a mock function with given fields: v
func (_m *wsConnection) ReadJSON(v interface{}) error {
	ret := _m.Called(v)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}) error); ok {
		r0 = rf(v)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetReadLimit provides a mock function with given fields: limit
func (_m *wsConnection) SetReadLimit(limit int64) {
	_m.Called(limit)
}

// SetWriteDeadline provides a mock function with given fields: t
func (_m *wsConnection) SetWriteDeadline(t time.Time) error {
	ret := _m.Called(t)

	var r0 error
	if rf, ok := ret.Get(0).(func(time.Time) error); ok {
		r0 = rf(t)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WriteJSON provides a mock function with given fields: v
func (_m *wsConnection) WriteJSON(v interface{}) error {
	ret := _m.Called(v)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}) error); ok {
		r0 = rf(v)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTnewWsConnection interface {
	mock.TestingT
	Cleanup(func())
}

// newWsConnection creates a new instance of wsConnection. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func newWsConnection(t mockConstructorTestingTnewWsConnection) *wsConnection {
	mock := &wsConnection{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
