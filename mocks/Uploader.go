// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	backupmodel "github.com/thomasduchatelle/dphoto/dphoto/backup/backupmodel"
)

// Uploader is an autogenerated mock type for the Uploader type
type Uploader struct {
	mock.Mock
}

// Execute provides a mock function with given fields: buffer, progressChannel
func (_m *Uploader) Execute(buffer []*backupmodel.AnalysedMedia, progressChannel chan *backupmodel.ProgressEvent) error {
	ret := _m.Called(buffer, progressChannel)

	var r0 error
	if rf, ok := ret.Get(0).(func([]*backupmodel.AnalysedMedia, chan *backupmodel.ProgressEvent) error); ok {
		r0 = rf(buffer, progressChannel)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
