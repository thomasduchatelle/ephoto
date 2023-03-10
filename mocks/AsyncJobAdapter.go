// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	archive "github.com/thomasduchatelle/ephoto/pkg/archive"
)

// AsyncJobAdapter is an autogenerated mock type for the AsyncJobAdapter type
type AsyncJobAdapter struct {
	mock.Mock
}

// LoadImagesInCache provides a mock function with given fields: images
func (_m *AsyncJobAdapter) LoadImagesInCache(images ...*archive.ImageToResize) error {
	_va := make([]interface{}, len(images))
	for _i := range images {
		_va[_i] = images[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(...*archive.ImageToResize) error); ok {
		r0 = rf(images...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WarmUpCacheByFolder provides a mock function with given fields: owner, missedStoreKey, width
func (_m *AsyncJobAdapter) WarmUpCacheByFolder(owner string, missedStoreKey string, width int) error {
	ret := _m.Called(owner, missedStoreKey, width)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, int) error); ok {
		r0 = rf(owner, missedStoreKey, width)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewAsyncJobAdapter interface {
	mock.TestingT
	Cleanup(func())
}

// NewAsyncJobAdapter creates a new instance of AsyncJobAdapter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewAsyncJobAdapter(t mockConstructorTestingTNewAsyncJobAdapter) *AsyncJobAdapter {
	mock := &AsyncJobAdapter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
