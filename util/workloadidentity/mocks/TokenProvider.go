// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// TokenProvider is an autogenerated mock type for the TokenProvider type
type TokenProvider struct {
	mock.Mock
}

// GetToken provides a mock function with given fields: scope
func (_m *TokenProvider) GetToken(scope string) (string, error) {
	ret := _m.Called(scope)

	if len(ret) == 0 {
		panic("no return value specified for GetToken")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (string, error)); ok {
		return rf(scope)
	}
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(scope)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(scope)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewTokenProvider creates a new instance of TokenProvider. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewTokenProvider(t interface {
	mock.TestingT
	Cleanup(func())
}) *TokenProvider {
	mock := &TokenProvider{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
