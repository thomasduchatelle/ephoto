// Code generated by mockery v2.3.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// FoundMediaWithHash is an autogenerated mock type for the FoundMediaWithHash type
type FoundMediaWithHash struct {
	mock.Mock
}

// Sha256Hash provides a mock function with given fields:
func (_m *FoundMediaWithHash) Sha256Hash() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}