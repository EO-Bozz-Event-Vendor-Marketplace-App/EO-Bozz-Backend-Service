// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import (
	review "capstone-alta1/features/review"

	mock "github.com/stretchr/testify/mock"
)

// ReviewRepository is an autogenerated mock type for the RepositoryInterface type
type ReviewRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: input
func (_m *ReviewRepository) Create(input review.Core) error {
	ret := _m.Called(input)

	var r0 error
	if rf, ok := ret.Get(0).(func(review.Core) error); ok {
		r0 = rf(input)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: id
func (_m *ReviewRepository) Delete(id uint) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAll provides a mock function with given fields:
func (_m *ReviewRepository) GetAll() ([]review.Core, error) {
	ret := _m.Called()

	var r0 []review.Core
	if rf, ok := ret.Get(0).(func() []review.Core); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]review.Core)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetById provides a mock function with given fields: id
func (_m *ReviewRepository) GetById(id uint) (review.Core, error) {
	ret := _m.Called(id)

	var r0 review.Core
	if rf, ok := ret.Get(0).(func(uint) review.Core); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(review.Core)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: input, id
func (_m *ReviewRepository) Update(input review.Core, id uint) error {
	ret := _m.Called(input, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(review.Core, uint) error); ok {
		r0 = rf(input, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewReviewRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewReviewRepository creates a new instance of ReviewRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewReviewRepository(t mockConstructorTestingTNewReviewRepository) *ReviewRepository {
	mock := &ReviewRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
