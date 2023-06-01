// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	context "context"

	anypb "google.golang.org/protobuf/types/known/anypb"

	integrations "github.com/chainloop-dev/chainloop/app/controlplane/integrations"

	mock "github.com/stretchr/testify/mock"
)

// Registrable is an autogenerated mock type for the Registrable type
type Registrable struct {
	mock.Mock
}

// PreRegister provides a mock function with given fields: ctx, req
func (_m *Registrable) PreRegister(ctx context.Context, req *anypb.Any) (*integrations.PreRegistration, error) {
	ret := _m.Called(ctx, req)

	var r0 *integrations.PreRegistration
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *anypb.Any) (*integrations.PreRegistration, error)); ok {
		return rf(ctx, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *anypb.Any) *integrations.PreRegistration); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*integrations.PreRegistration)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *anypb.Any) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewRegistrable interface {
	mock.TestingT
	Cleanup(func())
}

// NewRegistrable creates a new instance of Registrable. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRegistrable(t mockConstructorTestingTNewRegistrable) *Registrable {
	mock := &Registrable{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}