// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// ARepositoryAdapter is an autogenerated mock type for the ARepositoryAdapter type
type ARepositoryAdapter struct {
	mock.Mock
}

// AddLocation provides a mock function with given fields: owner, id, key
func (_m *ARepositoryAdapter) AddLocation(owner string, id string, key string) error {
	ret := _m.Called(owner, id, key)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, string) error); ok {
		r0 = rf(owner, id, key)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindById provides a mock function with given fields: owner, id
func (_m *ARepositoryAdapter) FindById(owner string, id string) (string, error) {
	ret := _m.Called(owner, id)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, string) string); ok {
		r0 = rf(owner, id)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(owner, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByIds provides a mock function with given fields: owner, ids
func (_m *ARepositoryAdapter) FindByIds(owner string, ids []string) (map[string]string, error) {
	ret := _m.Called(owner, ids)

	var r0 map[string]string
	if rf, ok := ret.Get(0).(func(string, []string) map[string]string); ok {
		r0 = rf(owner, ids)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, []string) error); ok {
		r1 = rf(owner, ids)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindIdsFromKeyPrefix provides a mock function with given fields: keyPrefix
func (_m *ARepositoryAdapter) FindIdsFromKeyPrefix(keyPrefix string) (map[string]string, error) {
	ret := _m.Called(keyPrefix)

	var r0 map[string]string
	if rf, ok := ret.Get(0).(func(string) map[string]string); ok {
		r0 = rf(keyPrefix)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(keyPrefix)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateLocations provides a mock function with given fields: owner, locations
func (_m *ARepositoryAdapter) UpdateLocations(owner string, locations map[string]string) error {
	ret := _m.Called(owner, locations)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, map[string]string) error); ok {
		r0 = rf(owner, locations)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewARepositoryAdapter interface {
	mock.TestingT
	Cleanup(func())
}

// NewARepositoryAdapter creates a new instance of ARepositoryAdapter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewARepositoryAdapter(t mockConstructorTestingTNewARepositoryAdapter) *ARepositoryAdapter {
	mock := &ARepositoryAdapter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
