// Code generated by mockery v2.13.1. DO NOT EDIT.

package mocks

import (
	factory "github.com/golangid/candi/codebase/factory"
	dependency "github.com/golangid/candi/codebase/factory/dependency"
	config "github.com/golangid/candi/config"

	mock "github.com/stretchr/testify/mock"

	types "github.com/golangid/candi/codebase/factory/types"
)

// ServiceFactory is an autogenerated mock type for the ServiceFactory type
type ServiceFactory struct {
	mock.Mock
}

// GetApplications provides a mock function with given fields:
func (_m *ServiceFactory) GetApplications() []factory.AppServerFactory {
	ret := _m.Called()

	var r0 []factory.AppServerFactory
	if rf, ok := ret.Get(0).(func() []factory.AppServerFactory); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]factory.AppServerFactory)
		}
	}

	return r0
}

// GetConfig provides a mock function with given fields:
func (_m *ServiceFactory) GetConfig() *config.Config {
	ret := _m.Called()

	var r0 *config.Config
	if rf, ok := ret.Get(0).(func() *config.Config); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*config.Config)
		}
	}

	return r0
}

// GetDependency provides a mock function with given fields:
func (_m *ServiceFactory) GetDependency() dependency.Dependency {
	ret := _m.Called()

	var r0 dependency.Dependency
	if rf, ok := ret.Get(0).(func() dependency.Dependency); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(dependency.Dependency)
		}
	}

	return r0
}

// GetModules provides a mock function with given fields:
func (_m *ServiceFactory) GetModules() []factory.ModuleFactory {
	ret := _m.Called()

	var r0 []factory.ModuleFactory
	if rf, ok := ret.Get(0).(func() []factory.ModuleFactory); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]factory.ModuleFactory)
		}
	}

	return r0
}

// Name provides a mock function with given fields:
func (_m *ServiceFactory) Name() types.Service {
	ret := _m.Called()

	var r0 types.Service
	if rf, ok := ret.Get(0).(func() types.Service); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(types.Service)
	}

	return r0
}

type mockConstructorTestingTNewServiceFactory interface {
	mock.TestingT
	Cleanup(func())
}

// NewServiceFactory creates a new instance of ServiceFactory. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewServiceFactory(t mockConstructorTestingTNewServiceFactory) *ServiceFactory {
	mock := &ServiceFactory{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
