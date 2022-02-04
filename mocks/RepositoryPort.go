// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	catalogmodel "github.com/thomasduchatelle/dphoto/domain/catalogmodel"
)

// RepositoryPort is an autogenerated mock type for the RepositoryPort type
type RepositoryPort struct {
	mock.Mock
}

// CountMedias provides a mock function with given fields: folderName
func (_m *RepositoryPort) CountMedias(folderName string) (int, error) {
	ret := _m.Called(folderName)

	var r0 int
	if rf, ok := ret.Get(0).(func(string) int); ok {
		r0 = rf(folderName)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(folderName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteEmptyAlbum provides a mock function with given fields: folderName
func (_m *RepositoryPort) DeleteEmptyAlbum(folderName string) error {
	ret := _m.Called(folderName)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(folderName)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteEmptyMoveTransaction provides a mock function with given fields: transactionId
func (_m *RepositoryPort) DeleteEmptyMoveTransaction(transactionId string) error {
	ret := _m.Called(transactionId)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(transactionId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindAlbum provides a mock function with given fields: folderName
func (_m *RepositoryPort) FindAlbum(folderName string) (*catalogmodel.Album, error) {
	ret := _m.Called(folderName)

	var r0 *catalogmodel.Album
	if rf, ok := ret.Get(0).(func(string) *catalogmodel.Album); ok {
		r0 = rf(folderName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*catalogmodel.Album)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(folderName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindAllAlbums provides a mock function with given fields:
func (_m *RepositoryPort) FindAllAlbums() ([]*catalogmodel.Album, error) {
	ret := _m.Called()

	var r0 []*catalogmodel.Album
	if rf, ok := ret.Get(0).(func() []*catalogmodel.Album); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*catalogmodel.Album)
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

// FindExistingSignatures provides a mock function with given fields: signatures
func (_m *RepositoryPort) FindExistingSignatures(signatures []*catalogmodel.MediaSignature) ([]*catalogmodel.MediaSignature, error) {
	ret := _m.Called(signatures)

	var r0 []*catalogmodel.MediaSignature
	if rf, ok := ret.Get(0).(func([]*catalogmodel.MediaSignature) []*catalogmodel.MediaSignature); ok {
		r0 = rf(signatures)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*catalogmodel.MediaSignature)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]*catalogmodel.MediaSignature) error); ok {
		r1 = rf(signatures)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindFilesToMove provides a mock function with given fields: transactionId, pageToken
func (_m *RepositoryPort) FindFilesToMove(transactionId string, pageToken string) ([]*catalogmodel.MovedMedia, string, error) {
	ret := _m.Called(transactionId, pageToken)

	var r0 []*catalogmodel.MovedMedia
	if rf, ok := ret.Get(0).(func(string, string) []*catalogmodel.MovedMedia); ok {
		r0 = rf(transactionId, pageToken)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*catalogmodel.MovedMedia)
		}
	}

	var r1 string
	if rf, ok := ret.Get(1).(func(string, string) string); ok {
		r1 = rf(transactionId, pageToken)
	} else {
		r1 = ret.Get(1).(string)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(string, string) error); ok {
		r2 = rf(transactionId, pageToken)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// FindMedias provides a mock function with given fields: folderName, filter
func (_m *RepositoryPort) FindMedias(folderName string, filter catalogmodel.FindMediaFilter) (*catalogmodel.MediaPage, error) {
	ret := _m.Called(folderName, filter)

	var r0 *catalogmodel.MediaPage
	if rf, ok := ret.Get(0).(func(string, catalogmodel.FindMediaFilter) *catalogmodel.MediaPage); ok {
		r0 = rf(folderName, filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*catalogmodel.MediaPage)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, catalogmodel.FindMediaFilter) error); ok {
		r1 = rf(folderName, filter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindReadyMoveTransactions provides a mock function with given fields:
func (_m *RepositoryPort) FindReadyMoveTransactions() ([]*catalogmodel.MoveTransaction, error) {
	ret := _m.Called()

	var r0 []*catalogmodel.MoveTransaction
	if rf, ok := ret.Get(0).(func() []*catalogmodel.MoveTransaction); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*catalogmodel.MoveTransaction)
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

// InsertAlbum provides a mock function with given fields: album
func (_m *RepositoryPort) InsertAlbum(album catalogmodel.Album) error {
	ret := _m.Called(album)

	var r0 error
	if rf, ok := ret.Get(0).(func(catalogmodel.Album) error); ok {
		r0 = rf(album)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// InsertMedias provides a mock function with given fields: media
func (_m *RepositoryPort) InsertMedias(media []catalogmodel.CreateMediaRequest) error {
	ret := _m.Called(media)

	var r0 error
	if rf, ok := ret.Get(0).(func([]catalogmodel.CreateMediaRequest) error); ok {
		r0 = rf(media)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateAlbum provides a mock function with given fields: album
func (_m *RepositoryPort) UpdateAlbum(album catalogmodel.Album) error {
	ret := _m.Called(album)

	var r0 error
	if rf, ok := ret.Get(0).(func(catalogmodel.Album) error); ok {
		r0 = rf(album)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateMedias provides a mock function with given fields: filter, newFolderName
func (_m *RepositoryPort) UpdateMedias(filter *catalogmodel.UpdateMediaFilter, newFolderName string) (string, int, error) {
	ret := _m.Called(filter, newFolderName)

	var r0 string
	if rf, ok := ret.Get(0).(func(*catalogmodel.UpdateMediaFilter, string) string); ok {
		r0 = rf(filter, newFolderName)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 int
	if rf, ok := ret.Get(1).(func(*catalogmodel.UpdateMediaFilter, string) int); ok {
		r1 = rf(filter, newFolderName)
	} else {
		r1 = ret.Get(1).(int)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(*catalogmodel.UpdateMediaFilter, string) error); ok {
		r2 = rf(filter, newFolderName)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// UpdateMediasLocation provides a mock function with given fields: transactionId, moves
func (_m *RepositoryPort) UpdateMediasLocation(transactionId string, moves []*catalogmodel.MovedMedia) error {
	ret := _m.Called(transactionId, moves)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, []*catalogmodel.MovedMedia) error); ok {
		r0 = rf(transactionId, moves)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
