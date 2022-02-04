// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import (
	dynamodb "github.com/aws/aws-sdk-go/service/dynamodb"
	mock "github.com/stretchr/testify/mock"
)

// Stream is an autogenerated mock type for the Stream type
type Stream struct {
	mock.Mock
}

// Count provides a mock function with given fields:
func (_m *Stream) Count() int64 {
	ret := _m.Called()

	var r0 int64
	if rf, ok := ret.Get(0).(func() int64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int64)
	}

	return r0
}

// Error provides a mock function with given fields:
func (_m *Stream) Error() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// HasNext provides a mock function with given fields:
func (_m *Stream) HasNext() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Next provides a mock function with given fields:
func (_m *Stream) Next() map[string]*dynamodb.AttributeValue {
	ret := _m.Called()

	var r0 map[string]*dynamodb.AttributeValue
	if rf, ok := ret.Get(0).(func() map[string]*dynamodb.AttributeValue); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]*dynamodb.AttributeValue)
		}
	}

	return r0
}
