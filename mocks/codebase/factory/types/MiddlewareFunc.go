// Code generated by mockery v2.13.1. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// MiddlewareFunc is an autogenerated mock type for the MiddlewareFunc type
type MiddlewareFunc struct {
	mock.Mock
}

// Execute provides a mock function with given fields: _a0
func (_m *MiddlewareFunc) Execute(_a0 context.Context) context.Context {
	ret := _m.Called(_a0)

	var r0 context.Context
	if rf, ok := ret.Get(0).(func(context.Context) context.Context); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(context.Context)
		}
	}

	return r0
}

type mockConstructorTestingTNewMiddlewareFunc interface {
	mock.TestingT
	Cleanup(func())
}

// NewMiddlewareFunc creates a new instance of MiddlewareFunc. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMiddlewareFunc(t mockConstructorTestingTNewMiddlewareFunc) *MiddlewareFunc {
	mock := &MiddlewareFunc{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
