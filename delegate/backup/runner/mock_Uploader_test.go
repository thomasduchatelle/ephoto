// Code generated by mockery v2.3.0. DO NOT EDIT.

package runner

import (
	model "duchatelle.io/dphoto/dphoto/backup/model"
	mock "github.com/stretchr/testify/mock"
)

// MockUploader is an autogenerated mock type for the Uploader type
type MockUploader struct {
	mock.Mock
}

// Execute provides a mock function with given fields: buffer
func (_m *MockUploader) Execute(buffer []*model.AnalysedMedia) error {
	ret := _m.Called(buffer)

	var r0 error
	if rf, ok := ret.Get(0).(func([]*model.AnalysedMedia) error); ok {
		r0 = rf(buffer)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
