// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	context "context"

	casclient "github.com/chainloop-dev/chainloop/internal/casclient"

	io "io"

	mock "github.com/stretchr/testify/mock"
)

// DownloaderUploader is an autogenerated mock type for the DownloaderUploader type
type DownloaderUploader struct {
	mock.Mock
}

// Download provides a mock function with given fields: ctx, w, digest
func (_m *DownloaderUploader) Download(ctx context.Context, w io.Writer, digest string) error {
	ret := _m.Called(ctx, w, digest)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, io.Writer, string) error); ok {
		r0 = rf(ctx, w, digest)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Upload provides a mock function with given fields: ctx, r, digest, fileName
func (_m *DownloaderUploader) Upload(ctx context.Context, r io.Reader, digest string, fileName string) (*casclient.UpDownStatus, error) {
	ret := _m.Called(ctx, r, digest, fileName)

	var r0 *casclient.UpDownStatus
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, io.Reader, string, string) (*casclient.UpDownStatus, error)); ok {
		return rf(ctx, r, digest, fileName)
	}
	if rf, ok := ret.Get(0).(func(context.Context, io.Reader, string, string) *casclient.UpDownStatus); ok {
		r0 = rf(ctx, r, digest, fileName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*casclient.UpDownStatus)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, io.Reader, string, string) error); ok {
		r1 = rf(ctx, r, digest, fileName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UploadFile provides a mock function with given fields: ctx, filepath
func (_m *DownloaderUploader) UploadFile(ctx context.Context, filepath string) (*casclient.UpDownStatus, error) {
	ret := _m.Called(ctx, filepath)

	var r0 *casclient.UpDownStatus
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*casclient.UpDownStatus, error)); ok {
		return rf(ctx, filepath)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *casclient.UpDownStatus); ok {
		r0 = rf(ctx, filepath)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*casclient.UpDownStatus)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, filepath)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewDownloaderUploader interface {
	mock.TestingT
	Cleanup(func())
}

// NewDownloaderUploader creates a new instance of DownloaderUploader. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewDownloaderUploader(t mockConstructorTestingTNewDownloaderUploader) *DownloaderUploader {
	mock := &DownloaderUploader{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
