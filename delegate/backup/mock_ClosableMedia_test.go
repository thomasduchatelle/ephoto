// Code generated by mockery v2.3.0. DO NOT EDIT.

package backup

import mock "github.com/stretchr/testify/mock"

// MockClosableMedia is an autogenerated mock type for the ClosableMedia type
type MockClosableMedia struct {
	mock.Mock
}

// Close provides a mock function with given fields:
func (_m *MockClosableMedia) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
