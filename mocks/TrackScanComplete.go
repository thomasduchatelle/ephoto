// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	backup "github.com/thomasduchatelle/ephoto/pkg/backup"
)

// TrackScanComplete is an autogenerated mock type for the TrackScanComplete type
type TrackScanComplete struct {
	mock.Mock
}

// OnScanComplete provides a mock function with given fields: total
func (_m *TrackScanComplete) OnScanComplete(total backup.MediaCounter) {
	_m.Called(total)
}

type mockConstructorTestingTNewTrackScanComplete interface {
	mock.TestingT
	Cleanup(func())
}

// NewTrackScanComplete creates a new instance of TrackScanComplete. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewTrackScanComplete(t mockConstructorTestingTNewTrackScanComplete) *TrackScanComplete {
	mock := &TrackScanComplete{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
