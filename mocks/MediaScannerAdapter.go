// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	backupmodel "github.com/thomasduchatelle/dphoto/dphoto/backup/backupmodel"
)

// MediaScannerAdapter is an autogenerated mock type for the MediaScannerAdapter type
type MediaScannerAdapter struct {
	mock.Mock
}

// FindMediaRecursively provides a mock function with given fields: volume, callback
func (_m *MediaScannerAdapter) FindMediaRecursively(volume backupmodel.VolumeToBackup, callback func(backupmodel.FoundMedia)) (uint, uint, error) {
	ret := _m.Called(volume, callback)

	var r0 uint
	if rf, ok := ret.Get(0).(func(backupmodel.VolumeToBackup, func(backupmodel.FoundMedia)) uint); ok {
		r0 = rf(volume, callback)
	} else {
		r0 = ret.Get(0).(uint)
	}

	var r1 uint
	if rf, ok := ret.Get(1).(func(backupmodel.VolumeToBackup, func(backupmodel.FoundMedia)) uint); ok {
		r1 = rf(volume, callback)
	} else {
		r1 = ret.Get(1).(uint)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(backupmodel.VolumeToBackup, func(backupmodel.FoundMedia)) error); ok {
		r2 = rf(volume, callback)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}
