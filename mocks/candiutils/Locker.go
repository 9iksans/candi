// Code generated by mockery v2.13.1. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Locker is an autogenerated mock type for the Locker type
type Locker struct {
	mock.Mock
}

// IsLocked provides a mock function with given fields: key
func (_m *Locker) IsLocked(key string) bool {
	ret := _m.Called(key)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Reset provides a mock function with given fields: key
func (_m *Locker) Reset(key string) {
	_m.Called(key)
}

// Unlock provides a mock function with given fields: key
func (_m *Locker) Unlock(key string) {
	_m.Called(key)
}

type mockConstructorTestingTNewLocker interface {
	mock.TestingT
	Cleanup(func())
}

// NewLocker creates a new instance of Locker. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewLocker(t mockConstructorTestingTNewLocker) *Locker {
	mock := &Locker{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
