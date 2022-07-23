// Code generated by mockery v2.13.1. DO NOT EDIT.

package mocks

import (
	context "context"

	taskqueueworker "github.com/golangid/candi/codebase/app/task_queue_worker"
	mock "github.com/stretchr/testify/mock"
)

// Persistent is an autogenerated mock type for the Persistent type
type Persistent struct {
	mock.Mock
}

// AggregateAllTaskJob provides a mock function with given fields: ctx, filter
func (_m *Persistent) AggregateAllTaskJob(ctx context.Context, filter *taskqueueworker.Filter) []taskqueueworker.TaskSummary {
	ret := _m.Called(ctx, filter)

	var r0 []taskqueueworker.TaskSummary
	if rf, ok := ret.Get(0).(func(context.Context, *taskqueueworker.Filter) []taskqueueworker.TaskSummary); ok {
		r0 = rf(ctx, filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]taskqueueworker.TaskSummary)
		}
	}

	return r0
}

// CleanJob provides a mock function with given fields: ctx, filter
func (_m *Persistent) CleanJob(ctx context.Context, filter *taskqueueworker.Filter) int64 {
	ret := _m.Called(ctx, filter)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, *taskqueueworker.Filter) int64); ok {
		r0 = rf(ctx, filter)
	} else {
		r0 = ret.Get(0).(int64)
	}

	return r0
}

// CountAllJob provides a mock function with given fields: ctx, filter
func (_m *Persistent) CountAllJob(ctx context.Context, filter *taskqueueworker.Filter) int {
	ret := _m.Called(ctx, filter)

	var r0 int
	if rf, ok := ret.Get(0).(func(context.Context, *taskqueueworker.Filter) int); ok {
		r0 = rf(ctx, filter)
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// DeleteJob provides a mock function with given fields: ctx, id
func (_m *Persistent) DeleteJob(ctx context.Context, id string) (taskqueueworker.Job, error) {
	ret := _m.Called(ctx, id)

	var r0 taskqueueworker.Job
	if rf, ok := ret.Get(0).(func(context.Context, string) taskqueueworker.Job); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(taskqueueworker.Job)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindAllJob provides a mock function with given fields: ctx, filter
func (_m *Persistent) FindAllJob(ctx context.Context, filter *taskqueueworker.Filter) []taskqueueworker.Job {
	ret := _m.Called(ctx, filter)

	var r0 []taskqueueworker.Job
	if rf, ok := ret.Get(0).(func(context.Context, *taskqueueworker.Filter) []taskqueueworker.Job); ok {
		r0 = rf(ctx, filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]taskqueueworker.Job)
		}
	}

	return r0
}

// FindJobByID provides a mock function with given fields: ctx, id, excludeFields
func (_m *Persistent) FindJobByID(ctx context.Context, id string, excludeFields ...string) (taskqueueworker.Job, error) {
	_va := make([]interface{}, len(excludeFields))
	for _i := range excludeFields {
		_va[_i] = excludeFields[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, id)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 taskqueueworker.Job
	if rf, ok := ret.Get(0).(func(context.Context, string, ...string) taskqueueworker.Job); ok {
		r0 = rf(ctx, id, excludeFields...)
	} else {
		r0 = ret.Get(0).(taskqueueworker.Job)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, ...string) error); ok {
		r1 = rf(ctx, id, excludeFields...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Ping provides a mock function with given fields: ctx
func (_m *Persistent) Ping(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SaveJob provides a mock function with given fields: ctx, job, retryHistories
func (_m *Persistent) SaveJob(ctx context.Context, job *taskqueueworker.Job, retryHistories ...taskqueueworker.RetryHistory) {
	_va := make([]interface{}, len(retryHistories))
	for _i := range retryHistories {
		_va[_i] = retryHistories[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, job)
	_ca = append(_ca, _va...)
	_m.Called(_ca...)
}

// SetSummary provides a mock function with given fields: _a0
func (_m *Persistent) SetSummary(_a0 taskqueueworker.Summary) {
	_m.Called(_a0)
}

// Summary provides a mock function with given fields:
func (_m *Persistent) Summary() taskqueueworker.Summary {
	ret := _m.Called()

	var r0 taskqueueworker.Summary
	if rf, ok := ret.Get(0).(func() taskqueueworker.Summary); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(taskqueueworker.Summary)
		}
	}

	return r0
}

// UpdateJob provides a mock function with given fields: ctx, filter, updated, retryHistories
func (_m *Persistent) UpdateJob(ctx context.Context, filter *taskqueueworker.Filter, updated map[string]interface{}, retryHistories ...taskqueueworker.RetryHistory) (int64, int64, error) {
	_va := make([]interface{}, len(retryHistories))
	for _i := range retryHistories {
		_va[_i] = retryHistories[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, filter, updated)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, *taskqueueworker.Filter, map[string]interface{}, ...taskqueueworker.RetryHistory) int64); ok {
		r0 = rf(ctx, filter, updated, retryHistories...)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 int64
	if rf, ok := ret.Get(1).(func(context.Context, *taskqueueworker.Filter, map[string]interface{}, ...taskqueueworker.RetryHistory) int64); ok {
		r1 = rf(ctx, filter, updated, retryHistories...)
	} else {
		r1 = ret.Get(1).(int64)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, *taskqueueworker.Filter, map[string]interface{}, ...taskqueueworker.RetryHistory) error); ok {
		r2 = rf(ctx, filter, updated, retryHistories...)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

type mockConstructorTestingTNewPersistent interface {
	mock.TestingT
	Cleanup(func())
}

// NewPersistent creates a new instance of Persistent. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewPersistent(t mockConstructorTestingTNewPersistent) *Persistent {
	mock := &Persistent{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
