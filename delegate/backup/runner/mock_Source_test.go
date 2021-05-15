// Code generated by mockery v2.3.0. DO NOT EDIT.

package runner

import (
	model "duchatelle.io/dphoto/dphoto/backup/model"
	mock "github.com/stretchr/testify/mock"
)

// MockSource is an autogenerated mock type for the Source type
type MockSource struct {
	mock.Mock
}

// Execute provides a mock function with given fields: medias
func (_m *MockSource) Execute(medias chan model.FoundMedia) (int, int, error) {
	ret := _m.Called(medias)

	var r0 int
	if rf, ok := ret.Get(0).(func(chan model.FoundMedia) int); ok {
		r0 = rf(medias)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 int
	if rf, ok := ret.Get(1).(func(chan model.FoundMedia) int); ok {
		r1 = rf(medias)
	} else {
		r1 = ret.Get(1).(int)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(chan model.FoundMedia) error); ok {
		r2 = rf(medias)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}