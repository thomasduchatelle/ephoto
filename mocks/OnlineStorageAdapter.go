// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	backupmodel "github.com/thomasduchatelle/dphoto/dphoto/backup/backupmodel"
)

// OnlineStorageAdapter is an autogenerated mock type for the OnlineStorageAdapter type
type OnlineStorageAdapter struct {
	mock.Mock
}

// MoveFile provides a mock function with given fields: owner, folderName, filename, destFolderName
func (_m *OnlineStorageAdapter) MoveFile(owner string, folderName string, filename string, destFolderName string) (string, error) {
	ret := _m.Called(owner, folderName, filename, destFolderName)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, string, string, string) string); ok {
		r0 = rf(owner, folderName, filename, destFolderName)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, string, string) error); ok {
		r1 = rf(owner, folderName, filename, destFolderName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UploadFile provides a mock function with given fields: owner, media, folderName, filename
func (_m *OnlineStorageAdapter) UploadFile(owner string, media backupmodel.ReadableMedia, folderName string, filename string) (string, error) {
	ret := _m.Called(owner, media, folderName, filename)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, backupmodel.ReadableMedia, string, string) string); ok {
		r0 = rf(owner, media, folderName, filename)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, backupmodel.ReadableMedia, string, string) error); ok {
		r1 = rf(owner, media, folderName, filename)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
