// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import (
	time "time"

	mock "github.com/stretchr/testify/mock"

	ui "github.com/thomasduchatelle/dphoto/dphoto/cmd/ui"
)

// InteractiveActionsPort is an autogenerated mock type for the InteractiveActionsPort type
type InteractiveActionsPort struct {
	mock.Mock
}

// BackupSuggestion provides a mock function with given fields: record, existing, listener
func (_m *InteractiveActionsPort) BackupSuggestion(record *ui.SuggestionRecord, existing *ui.ExistingRecord, listener ui.InteractiveRendererPort) error {
	ret := _m.Called(record, existing, listener)

	var r0 error
	if rf, ok := ret.Get(0).(func(*ui.SuggestionRecord, *ui.ExistingRecord, ui.InteractiveRendererPort) error); ok {
		r0 = rf(record, existing, listener)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Create provides a mock function with given fields: createRequest
func (_m *InteractiveActionsPort) Create(createRequest ui.RecordCreation) error {
	ret := _m.Called(createRequest)

	var r0 error
	if rf, ok := ret.Get(0).(func(ui.RecordCreation) error); ok {
		r0 = rf(createRequest)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteAlbum provides a mock function with given fields: folderName
func (_m *InteractiveActionsPort) DeleteAlbum(folderName string) error {
	ret := _m.Called(folderName)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(folderName)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RenameAlbum provides a mock function with given fields: folderName, newName, renameFolder
func (_m *InteractiveActionsPort) RenameAlbum(folderName string, newName string, renameFolder bool) error {
	ret := _m.Called(folderName, newName, renameFolder)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, bool) error); ok {
		r0 = rf(folderName, newName, renameFolder)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateAlbum provides a mock function with given fields: folderName, start, end
func (_m *InteractiveActionsPort) UpdateAlbum(folderName string, start time.Time, end time.Time) error {
	ret := _m.Called(folderName, start, end)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, time.Time, time.Time) error); ok {
		r0 = rf(folderName, start, end)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
