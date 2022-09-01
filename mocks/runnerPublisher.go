// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	backup "github.com/thomasduchatelle/dphoto/domain/backup"
)

// runnerPublisher is an autogenerated mock type for the runnerPublisher type
type runnerPublisher struct {
	mock.Mock
}

// Execute provides a mock function with given fields: _a0, _a1
func (_m *runnerPublisher) Execute(_a0 chan backup.FoundMedia, _a1 chan *backup.ProgressEvent) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(chan backup.FoundMedia, chan *backup.ProgressEvent) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTnewRunnerPublisher interface {
	mock.TestingT
	Cleanup(func())
}

// newRunnerPublisher creates a new instance of runnerPublisher. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func newRunnerPublisher(t mockConstructorTestingTnewRunnerPublisher) *runnerPublisher {
	mock := &runnerPublisher{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
