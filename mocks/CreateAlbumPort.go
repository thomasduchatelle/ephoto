// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	"github.com/thomasduchatelle/ephoto/cmd/dphoto/cmd/ui"
)

// CreateAlbumPort is an autogenerated mock type for the CreateAlbumPort type
type CreateAlbumPort struct {
	mock.Mock
}

// Create provides a mock function with given fields: createRequest
func (_m *CreateAlbumPort) Create(createRequest ui.RecordCreation) error {
	ret := _m.Called(createRequest)

	var r0 error
	if rf, ok := ret.Get(0).(func(ui.RecordCreation) error); ok {
		r0 = rf(createRequest)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewCreateAlbumPort interface {
	mock.TestingT
	Cleanup(func())
}

// NewCreateAlbumPort creates a new instance of CreateAlbumPort. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewCreateAlbumPort(t mockConstructorTestingTNewCreateAlbumPort) *CreateAlbumPort {
	mock := &CreateAlbumPort{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
