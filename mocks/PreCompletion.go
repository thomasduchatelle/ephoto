// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// PreCompletion is an autogenerated mock type for the PreCompletion type
type PreCompletion struct {
	mock.Mock
}

// Execute provides a mock function with given fields:
func (_m *PreCompletion) Execute() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
