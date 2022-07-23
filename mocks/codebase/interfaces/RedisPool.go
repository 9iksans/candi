// Code generated by mockery v2.13.1. DO NOT EDIT.

package mocks

import (
	context "context"

	interfaces "github.com/golangid/candi/codebase/interfaces"
	mock "github.com/stretchr/testify/mock"

	redis "github.com/gomodule/redigo/redis"
)

// RedisPool is an autogenerated mock type for the RedisPool type
type RedisPool struct {
	mock.Mock
}

// Cache provides a mock function with given fields:
func (_m *RedisPool) Cache() interfaces.Cache {
	ret := _m.Called()

	var r0 interfaces.Cache
	if rf, ok := ret.Get(0).(func() interfaces.Cache); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interfaces.Cache)
		}
	}

	return r0
}

// Disconnect provides a mock function with given fields: ctx
func (_m *RedisPool) Disconnect(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Health provides a mock function with given fields:
func (_m *RedisPool) Health() map[string]error {
	ret := _m.Called()

	var r0 map[string]error
	if rf, ok := ret.Get(0).(func() map[string]error); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]error)
		}
	}

	return r0
}

// ReadPool provides a mock function with given fields:
func (_m *RedisPool) ReadPool() *redis.Pool {
	ret := _m.Called()

	var r0 *redis.Pool
	if rf, ok := ret.Get(0).(func() *redis.Pool); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*redis.Pool)
		}
	}

	return r0
}

// WritePool provides a mock function with given fields:
func (_m *RedisPool) WritePool() *redis.Pool {
	ret := _m.Called()

	var r0 *redis.Pool
	if rf, ok := ret.Get(0).(func() *redis.Pool); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*redis.Pool)
		}
	}

	return r0
}

type mockConstructorTestingTNewRedisPool interface {
	mock.TestingT
	Cleanup(func())
}

// NewRedisPool creates a new instance of RedisPool. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRedisPool(t mockConstructorTestingTNewRedisPool) *RedisPool {
	mock := &RedisPool{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
