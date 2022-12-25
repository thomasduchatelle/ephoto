// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	"github.com/thomasduchatelle/ephoto/cmd/fphoto/screen"
)

// Segment is an autogenerated mock type for the Segment type
type Segment struct {
	mock.Mock
}

// Content provides a mock function with given fields: options
func (_m *Segment) Content(options screen.RenderingOptions) string {
	ret := _m.Called(options)

	var r0 string
	if rf, ok := ret.Get(0).(func(screen.RenderingOptions) string); ok {
		r0 = rf(options)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

type mockConstructorTestingTNewSegment interface {
	mock.TestingT
	Cleanup(func())
}

// NewSegment creates a new instance of Segment. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewSegment(t mockConstructorTestingTNewSegment) *Segment {
	mock := &Segment{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
