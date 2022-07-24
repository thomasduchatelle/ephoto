// Code generated by mockery v2.12.1. DO NOT EDIT.

package mocks

import (
	testing "testing"

	mock "github.com/stretchr/testify/mock"
)

// JobQueueAdapter is an autogenerated mock type for the JobQueueAdapter type
type JobQueueAdapter struct {
	mock.Mock
}

// ReportMisfiredCache provides a mock function with given fields: owner, storeKey, width
func (_m *JobQueueAdapter) ReportMisfiredCache(owner string, storeKey string, width int) error {
	ret := _m.Called(owner, storeKey, width)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, int) error); ok {
		r0 = rf(owner, storeKey, width)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewJobQueueAdapter creates a new instance of JobQueueAdapter. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewJobQueueAdapter(t testing.TB) *JobQueueAdapter {
	mock := &JobQueueAdapter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}