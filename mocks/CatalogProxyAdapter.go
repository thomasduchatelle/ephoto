// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	catalogmodel "github.com/thomasduchatelle/dphoto/domain/catalogmodel"
)

// CatalogProxyAdapter is an autogenerated mock type for the CatalogProxyAdapter type
type CatalogProxyAdapter struct {
	mock.Mock
}

// Create provides a mock function with given fields: createRequest
func (_m *CatalogProxyAdapter) Create(createRequest catalogmodel.CreateAlbum) error {
	ret := _m.Called(createRequest)

	var r0 error
	if rf, ok := ret.Get(0).(func(catalogmodel.CreateAlbum) error); ok {
		r0 = rf(createRequest)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindAllAlbums provides a mock function with given fields:
func (_m *CatalogProxyAdapter) FindAllAlbums() ([]*catalogmodel.Album, error) {
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

// FindSignatures provides a mock function with given fields: signatures
func (_m *CatalogProxyAdapter) FindSignatures(signatures []*catalogmodel.MediaSignature) ([]*catalogmodel.MediaSignature, error) {
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

// InsertMedias provides a mock function with given fields: medias
func (_m *CatalogProxyAdapter) InsertMedias(medias []catalogmodel.CreateMediaRequest) error {
	ret := _m.Called(medias)

	var r0 error
	if rf, ok := ret.Get(0).(func([]catalogmodel.CreateMediaRequest) error); ok {
		r0 = rf(medias)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
