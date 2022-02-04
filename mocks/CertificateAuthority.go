// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	dnsdomain "github.com/thomasduchatelle/dphoto/domain/dnsdomain"
)

// CertificateAuthority is an autogenerated mock type for the CertificateAuthority type
type CertificateAuthority struct {
	mock.Mock
}

// RequestCertificate provides a mock function with given fields: email, domain
func (_m *CertificateAuthority) RequestCertificate(email string, domain string) (*dnsdomain.CompleteCertificate, error) {
	ret := _m.Called(email, domain)

	var r0 *dnsdomain.CompleteCertificate
	if rf, ok := ret.Get(0).(func(string, string) *dnsdomain.CompleteCertificate); ok {
		r0 = rf(email, domain)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dnsdomain.CompleteCertificate)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(email, domain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
