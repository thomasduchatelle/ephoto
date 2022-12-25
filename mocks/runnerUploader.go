// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	backup "github.com/thomasduchatelle/ephoto/pkg/backup"
)

// runnerUploader is an autogenerated mock type for the runnerUploader type
type runnerUploader struct {
	mock.Mock
}

// Execute provides a mock function with given fields: buffer, progressChannel
func (_m *runnerUploader) Execute(buffer []*backup.BackingUpMediaRequest, progressChannel chan *backup.ProgressEvent) error {
	ret := _m.Called(buffer, progressChannel)

	var r0 error
	if rf, ok := ret.Get(0).(func([]*backup.BackingUpMediaRequest, chan *backup.ProgressEvent) error); ok {
		r0 = rf(buffer, progressChannel)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTnewRunnerUploader interface {
	mock.TestingT
	Cleanup(func())
}

// newRunnerUploader creates a new instance of runnerUploader. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func newRunnerUploader(t mockConstructorTestingTnewRunnerUploader) *runnerUploader {
	mock := &runnerUploader{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
