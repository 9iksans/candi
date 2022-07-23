// Code generated by mockery v2.13.1. DO NOT EDIT.

package mocks

import (
	broker "github.com/golangid/candi/broker"
	mock "github.com/stretchr/testify/mock"
)

// RabbitMQOptionFunc is an autogenerated mock type for the RabbitMQOptionFunc type
type RabbitMQOptionFunc struct {
	mock.Mock
}

// Execute provides a mock function with given fields: _a0
func (_m *RabbitMQOptionFunc) Execute(_a0 *broker.RabbitMQBroker) {
	_m.Called(_a0)
}

type mockConstructorTestingTNewRabbitMQOptionFunc interface {
	mock.TestingT
	Cleanup(func())
}

// NewRabbitMQOptionFunc creates a new instance of RabbitMQOptionFunc. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRabbitMQOptionFunc(t mockConstructorTestingTNewRabbitMQOptionFunc) *RabbitMQOptionFunc {
	mock := &RabbitMQOptionFunc{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
