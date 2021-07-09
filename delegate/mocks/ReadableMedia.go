// Code generated by mockery v2.3.0. DO NOT EDIT.

package mocks

import (
	io "io"

	backupmodel "duchatelle.io/dphoto/dphoto/backup/backupmodel"

	mock "github.com/stretchr/testify/mock"
)

// ReadableMedia is an autogenerated mock type for the ReadableMedia type
type ReadableMedia struct {
	mock.Mock
}

// ReadMedia provides a mock function with given fields:
func (_m *ReadableMedia) ReadMedia() (io.Reader, error) {
	ret := _m.Called()

	var r0 io.Reader
	if rf, ok := ret.Get(0).(func() io.Reader); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(io.Reader)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SimpleSignature provides a mock function with given fields:
func (_m *ReadableMedia) SimpleSignature() *backupmodel.SimpleMediaSignature {
	ret := _m.Called()

	var r0 *backupmodel.SimpleMediaSignature
	if rf, ok := ret.Get(0).(func() *backupmodel.SimpleMediaSignature); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*backupmodel.SimpleMediaSignature)
		}
	}

	return r0
}
