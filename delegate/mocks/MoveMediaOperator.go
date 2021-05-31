// Code generated by mockery v2.3.0. DO NOT EDIT.

package mocks

import (
	catalog "duchatelle.io/dphoto/dphoto/catalog"
	mock "github.com/stretchr/testify/mock"
)

// MoveMediaOperator is an autogenerated mock type for the MoveMediaOperator type
type MoveMediaOperator struct {
	mock.Mock
}

// Continue provides a mock function with given fields:
func (_m *MoveMediaOperator) Continue() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Move provides a mock function with given fields: source, dest
func (_m *MoveMediaOperator) Move(source catalog.MediaLocation, dest catalog.MediaLocation) error {
	ret := _m.Called(source, dest)

	var r0 error
	if rf, ok := ret.Get(0).(func(catalog.MediaLocation, catalog.MediaLocation) error); ok {
		r0 = rf(source, dest)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateStatus provides a mock function with given fields: done, total
func (_m *MoveMediaOperator) UpdateStatus(done int, total int) error {
	ret := _m.Called(done, total)

	var r0 error
	if rf, ok := ret.Get(0).(func(int, int) error); ok {
		r0 = rf(done, total)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}