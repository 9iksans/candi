// Code generated by mockery v2.13.1. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// GraphQLErrorResolver is an autogenerated mock type for the GraphQLErrorResolver type
type GraphQLErrorResolver struct {
	mock.Mock
}

// Error provides a mock function with given fields:
func (_m *GraphQLErrorResolver) Error() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Extensions provides a mock function with given fields:
func (_m *GraphQLErrorResolver) Extensions() map[string]interface{} {
	ret := _m.Called()

	var r0 map[string]interface{}
	if rf, ok := ret.Get(0).(func() map[string]interface{}); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]interface{})
		}
	}

	return r0
}

type mockConstructorTestingTNewGraphQLErrorResolver interface {
	mock.TestingT
	Cleanup(func())
}

// NewGraphQLErrorResolver creates a new instance of GraphQLErrorResolver. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewGraphQLErrorResolver(t mockConstructorTestingTNewGraphQLErrorResolver) *GraphQLErrorResolver {
	mock := &GraphQLErrorResolver{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
