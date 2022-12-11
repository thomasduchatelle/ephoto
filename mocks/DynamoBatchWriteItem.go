// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import (
	dynamodb "github.com/aws/aws-sdk-go/service/dynamodb"

	mock "github.com/stretchr/testify/mock"
)

// DynamoBatchWriteItem is an autogenerated mock type for the DynamoBatchWriteItem type
type DynamoBatchWriteItem struct {
	mock.Mock
}

// BatchWriteItem provides a mock function with given fields: _a0
func (_m *DynamoBatchWriteItem) BatchWriteItem(_a0 *dynamodb.BatchWriteItemInput) (*dynamodb.BatchWriteItemOutput, error) {
	ret := _m.Called(_a0)

	var r0 *dynamodb.BatchWriteItemOutput
	if rf, ok := ret.Get(0).(func(*dynamodb.BatchWriteItemInput) *dynamodb.BatchWriteItemOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dynamodb.BatchWriteItemOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*dynamodb.BatchWriteItemInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewDynamoBatchWriteItem interface {
	mock.TestingT
	Cleanup(func())
}

// NewDynamoBatchWriteItem creates a new instance of DynamoBatchWriteItem. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewDynamoBatchWriteItem(t mockConstructorTestingTNewDynamoBatchWriteItem) *DynamoBatchWriteItem {
	mock := &DynamoBatchWriteItem{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
