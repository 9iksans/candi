// Code generated by mockery v2.8.0. DO NOT EDIT.

package mocks

import (
	tracer "github.com/golangid/candi/tracer"
	mock "github.com/stretchr/testify/mock"
)

// OptionFunc is an autogenerated mock type for the OptionFunc type
type OptionFunc struct {
	mock.Mock
}

// Execute provides a mock function with given fields: _a0
func (_m *OptionFunc) Execute(_a0 *tracer.Option) {
	_m.Called(_a0)
}