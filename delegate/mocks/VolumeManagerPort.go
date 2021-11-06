// Code generated by mockery v2.3.0. DO NOT EDIT.

package mocks

import (
	backupmodel "github.com/thomasduchatelle/dphoto/delegate/backup/backupmodel"

	mock "github.com/stretchr/testify/mock"
)

// VolumeManagerPort is an autogenerated mock type for the VolumeManagerPort type
type VolumeManagerPort struct {
	mock.Mock
}

// OnMountedVolume provides a mock function with given fields: volume
func (_m *VolumeManagerPort) OnMountedVolume(volume backupmodel.VolumeToBackup) {
	_m.Called(volume)
}

// OnUnMountedVolume provides a mock function with given fields: uuid
func (_m *VolumeManagerPort) OnUnMountedVolume(uuid string) {
	_m.Called(uuid)
}
