// Code generated by mockery v2.3.0. DO NOT EDIT.

package mocks

import (
	backupmodel "github.com/thomasduchatelle/dphoto/delegate/backup/backupmodel"
	mock "github.com/stretchr/testify/mock"
)

// Filter is an autogenerated mock type for the Filter type
type Filter struct {
	mock.Mock
}

// Execute provides a mock function with given fields: found
func (_m *Filter) Execute(found backupmodel.FoundMedia) bool {
	ret := _m.Called(found)

	var r0 bool
	if rf, ok := ret.Get(0).(func(backupmodel.FoundMedia) bool); ok {
		r0 = rf(found)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}
