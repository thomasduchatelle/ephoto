// Code generated by mockery v2.3.0. DO NOT EDIT.

package mocks

import (
	ui "github.com/thomasduchatelle/dphoto/delegate/cmd/ui"
	mock "github.com/stretchr/testify/mock"
)

// ExistingRecordRepositoryPort is an autogenerated mock type for the ExistingRecordRepositoryPort type
type ExistingRecordRepositoryPort struct {
	mock.Mock
}

// FindExistingRecords provides a mock function with given fields:
func (_m *ExistingRecordRepositoryPort) FindExistingRecords() ([]*ui.ExistingRecord, error) {
	ret := _m.Called()

	var r0 []*ui.ExistingRecord
	if rf, ok := ret.Get(0).(func() []*ui.ExistingRecord); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*ui.ExistingRecord)
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
