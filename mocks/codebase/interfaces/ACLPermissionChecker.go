// Code generated by mockery v2.13.1. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// ACLPermissionChecker is an autogenerated mock type for the ACLPermissionChecker type
type ACLPermissionChecker struct {
	mock.Mock
}

// CheckPermission provides a mock function with given fields: ctx, userID, permissionCode
func (_m *ACLPermissionChecker) CheckPermission(ctx context.Context, userID string, permissionCode string) (string, error) {
	ret := _m.Called(ctx, userID, permissionCode)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, string, string) string); ok {
		r0 = rf(ctx, userID, permissionCode)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, userID, permissionCode)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewACLPermissionChecker interface {
	mock.TestingT
	Cleanup(func())
}

// NewACLPermissionChecker creates a new instance of ACLPermissionChecker. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewACLPermissionChecker(t mockConstructorTestingTNewACLPermissionChecker) *ACLPermissionChecker {
	mock := &ACLPermissionChecker{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
