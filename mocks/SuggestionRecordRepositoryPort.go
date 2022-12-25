// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	"github.com/thomasduchatelle/ephoto/cmd/dphoto/cmd/ui"
)

// SuggestionRecordRepositoryPort is an autogenerated mock type for the SuggestionRecordRepositoryPort type
type SuggestionRecordRepositoryPort struct {
	mock.Mock
}

// Count provides a mock function with given fields:
func (_m *SuggestionRecordRepositoryPort) Count() int {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// FindSuggestionRecords provides a mock function with given fields:
func (_m *SuggestionRecordRepositoryPort) FindSuggestionRecords() []*ui.SuggestionRecord {
	ret := _m.Called()

	var r0 []*ui.SuggestionRecord
	if rf, ok := ret.Get(0).(func() []*ui.SuggestionRecord); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*ui.SuggestionRecord)
		}
	}

	return r0
}

// Rejects provides a mock function with given fields:
func (_m *SuggestionRecordRepositoryPort) Rejects() int {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

type mockConstructorTestingTNewSuggestionRecordRepositoryPort interface {
	mock.TestingT
	Cleanup(func())
}

// NewSuggestionRecordRepositoryPort creates a new instance of SuggestionRecordRepositoryPort. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewSuggestionRecordRepositoryPort(t mockConstructorTestingTNewSuggestionRecordRepositoryPort) *SuggestionRecordRepositoryPort {
	mock := &SuggestionRecordRepositoryPort{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
