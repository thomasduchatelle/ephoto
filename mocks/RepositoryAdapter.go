// Code generated by mockery v2.12.1. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	catalog "github.com/thomasduchatelle/dphoto/domain/catalog"

	testing "testing"
)

// RepositoryAdapter is an autogenerated mock type for the RepositoryAdapter type
type RepositoryAdapter struct {
	mock.Mock
}

// DeleteEmptyAlbum provides a mock function with given fields: owner, folderName
func (_m *RepositoryAdapter) DeleteEmptyAlbum(owner string, folderName string) error {
	ret := _m.Called(owner, folderName)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(owner, folderName)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindAlbum provides a mock function with given fields: owner, folderName
func (_m *RepositoryAdapter) FindAlbum(owner string, folderName string) (*catalog.Album, error) {
	ret := _m.Called(owner, folderName)

	var r0 *catalog.Album
	if rf, ok := ret.Get(0).(func(string, string) *catalog.Album); ok {
		r0 = rf(owner, folderName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*catalog.Album)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(owner, folderName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindAllAlbums provides a mock function with given fields: owner
func (_m *RepositoryAdapter) FindAllAlbums(owner string) ([]*catalog.Album, error) {
	ret := _m.Called(owner)

	var r0 []*catalog.Album
	if rf, ok := ret.Get(0).(func(string) []*catalog.Album); ok {
		r0 = rf(owner)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*catalog.Album)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(owner)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindExistingSignatures provides a mock function with given fields: owner, signatures
func (_m *RepositoryAdapter) FindExistingSignatures(owner string, signatures []*catalog.MediaSignature) ([]*catalog.MediaSignature, error) {
	ret := _m.Called(owner, signatures)

	var r0 []*catalog.MediaSignature
	if rf, ok := ret.Get(0).(func(string, []*catalog.MediaSignature) []*catalog.MediaSignature); ok {
		r0 = rf(owner, signatures)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*catalog.MediaSignature)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, []*catalog.MediaSignature) error); ok {
		r1 = rf(owner, signatures)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindMediaIds provides a mock function with given fields: request
func (_m *RepositoryAdapter) FindMediaIds(request *catalog.FindMediaRequest) ([]string, error) {
	ret := _m.Called(request)

	var r0 []string
	if rf, ok := ret.Get(0).(func(*catalog.FindMediaRequest) []string); ok {
		r0 = rf(request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*catalog.FindMediaRequest) error); ok {
		r1 = rf(request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindMedias provides a mock function with given fields: request
func (_m *RepositoryAdapter) FindMedias(request *catalog.FindMediaRequest) ([]*catalog.MediaMeta, error) {
	ret := _m.Called(request)

	var r0 []*catalog.MediaMeta
	if rf, ok := ret.Get(0).(func(*catalog.FindMediaRequest) []*catalog.MediaMeta); ok {
		r0 = rf(request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*catalog.MediaMeta)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*catalog.FindMediaRequest) error); ok {
		r1 = rf(request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InsertAlbum provides a mock function with given fields: album
func (_m *RepositoryAdapter) InsertAlbum(album catalog.Album) error {
	ret := _m.Called(album)

	var r0 error
	if rf, ok := ret.Get(0).(func(catalog.Album) error); ok {
		r0 = rf(album)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// InsertMedias provides a mock function with given fields: owner, media
func (_m *RepositoryAdapter) InsertMedias(owner string, media []catalog.CreateMediaRequest) error {
	ret := _m.Called(owner, media)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, []catalog.CreateMediaRequest) error); ok {
		r0 = rf(owner, media)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// TransferMedias provides a mock function with given fields: owner, mediaIds, newFolderName
func (_m *RepositoryAdapter) TransferMedias(owner string, mediaIds []string, newFolderName string) error {
	ret := _m.Called(owner, mediaIds, newFolderName)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, []string, string) error); ok {
		r0 = rf(owner, mediaIds, newFolderName)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateAlbum provides a mock function with given fields: album
func (_m *RepositoryAdapter) UpdateAlbum(album catalog.Album) error {
	ret := _m.Called(album)

	var r0 error
	if rf, ok := ret.Get(0).(func(catalog.Album) error); ok {
		r0 = rf(album)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewRepositoryAdapter creates a new instance of RepositoryAdapter. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewRepositoryAdapter(t testing.TB) *RepositoryAdapter {
	mock := &RepositoryAdapter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
