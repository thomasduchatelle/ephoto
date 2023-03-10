// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	backup "github.com/thomasduchatelle/ephoto/pkg/backup"
)

// runnerAnalyser is an autogenerated mock type for the runnerAnalyser type
type runnerAnalyser struct {
	mock.Mock
}

// Execute provides a mock function with given fields: found, progressChannel
func (_m *runnerAnalyser) Execute(found backup.FoundMedia, progressChannel chan *backup.ProgressEvent) (*backup.AnalysedMedia, error) {
	ret := _m.Called(found, progressChannel)

	var r0 *backup.AnalysedMedia
	if rf, ok := ret.Get(0).(func(backup.FoundMedia, chan *backup.ProgressEvent) *backup.AnalysedMedia); ok {
		r0 = rf(found, progressChannel)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*backup.AnalysedMedia)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(backup.FoundMedia, chan *backup.ProgressEvent) error); ok {
		r1 = rf(found, progressChannel)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTnewRunnerAnalyser interface {
	mock.TestingT
	Cleanup(func())
}

// newRunnerAnalyser creates a new instance of runnerAnalyser. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func newRunnerAnalyser(t mockConstructorTestingTnewRunnerAnalyser) *runnerAnalyser {
	mock := &runnerAnalyser{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
