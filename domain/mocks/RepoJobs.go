// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	entity "CareerCenter/domain/entity"
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// RepoJobs is an autogenerated mock type for the RepoJobs type
type RepoJobs struct {
	mock.Mock
}

// GetListJobs provides a mock function with given fields: ctx, filter
func (_m *RepoJobs) GetListJobs(ctx context.Context, filter *entity.Filter) ([]*entity.JobsDTO, error) {
	ret := _m.Called(ctx, filter)

	var r0 []*entity.JobsDTO
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *entity.Filter) ([]*entity.JobsDTO, error)); ok {
		return rf(ctx, filter)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *entity.Filter) []*entity.JobsDTO); ok {
		r0 = rf(ctx, filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entity.JobsDTO)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *entity.Filter) error); ok {
		r1 = rf(ctx, filter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewRepoJobs interface {
	mock.TestingT
	Cleanup(func())
}

// NewRepoJobs creates a new instance of RepoJobs. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRepoJobs(t mockConstructorTestingTNewRepoJobs) *RepoJobs {
	mock := &RepoJobs{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
