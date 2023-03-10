// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import (
	partner "capstone-alta1/features/partner"

	mock "github.com/stretchr/testify/mock"
)

// PartnerRepository is an autogenerated mock type for the RepositoryInterface type
type PartnerRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: input
func (_m *PartnerRepository) Create(input partner.Core) error {
	ret := _m.Called(input)

	var r0 error
	if rf, ok := ret.Get(0).(func(partner.Core) error); ok {
		r0 = rf(input)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: partnerID, userID
func (_m *PartnerRepository) Delete(partnerID uint, userID uint) error {
	ret := _m.Called(partnerID, userID)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint, uint) error); ok {
		r0 = rf(partnerID, userID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindUser provides a mock function with given fields: email
func (_m *PartnerRepository) FindUser(email string) (partner.Core, error) {
	ret := _m.Called(email)

	var r0 partner.Core
	if rf, ok := ret.Get(0).(func(string) partner.Core); ok {
		r0 = rf(email)
	} else {
		r0 = ret.Get(0).(partner.Core)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAdditionals provides a mock function with given fields: partnerID
func (_m *PartnerRepository) GetAdditionals(partnerID uint) ([]partner.AdditionalCore, error) {
	ret := _m.Called(partnerID)

	var r0 []partner.AdditionalCore
	if rf, ok := ret.Get(0).(func(uint) []partner.AdditionalCore); ok {
		r0 = rf(partnerID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]partner.AdditionalCore)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(partnerID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAll provides a mock function with given fields: query
func (_m *PartnerRepository) GetAll(query string) ([]partner.Core, error) {
	ret := _m.Called(query)

	var r0 []partner.Core
	if rf, ok := ret.Get(0).(func(string) []partner.Core); ok {
		r0 = rf(query)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]partner.Core)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(query)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetById provides a mock function with given fields: id
func (_m *PartnerRepository) GetById(id uint) (partner.Core, error) {
	ret := _m.Called(id)

	var r0 partner.Core
	if rf, ok := ret.Get(0).(func(uint) partner.Core); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(partner.Core)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetOrders provides a mock function with given fields: partnerID
func (_m *PartnerRepository) GetOrders(partnerID uint) ([]partner.OrderCore, error) {
	ret := _m.Called(partnerID)

	var r0 []partner.OrderCore
	if rf, ok := ret.Get(0).(func(uint) []partner.OrderCore); ok {
		r0 = rf(partnerID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]partner.OrderCore)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(partnerID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPartnerRegisterData provides a mock function with given fields: queryCompanyName, queryPICName, queryPartnerStatus
func (_m *PartnerRepository) GetPartnerRegisterData(queryCompanyName string, queryPICName string, queryPartnerStatus string) ([]partner.Core, error) {
	ret := _m.Called(queryCompanyName, queryPICName, queryPartnerStatus)

	var r0 []partner.Core
	if rf, ok := ret.Get(0).(func(string, string, string) []partner.Core); ok {
		r0 = rf(queryCompanyName, queryPICName, queryPartnerStatus)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]partner.Core)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, string) error); ok {
		r1 = rf(queryCompanyName, queryPICName, queryPartnerStatus)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPartnerRegisterDataByID provides a mock function with given fields: partnerID
func (_m *PartnerRepository) GetPartnerRegisterDataByID(partnerID uint) (partner.Core, error) {
	ret := _m.Called(partnerID)

	var r0 partner.Core
	if rf, ok := ret.Get(0).(func(uint) partner.Core); ok {
		r0 = rf(partnerID)
	} else {
		r0 = ret.Get(0).(partner.Core)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(partnerID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetServices provides a mock function with given fields: partnerID
func (_m *PartnerRepository) GetServices(partnerID uint) ([]partner.ServiceCore, error) {
	ret := _m.Called(partnerID)

	var r0 []partner.ServiceCore
	if rf, ok := ret.Get(0).(func(uint) []partner.ServiceCore); ok {
		r0 = rf(partnerID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]partner.ServiceCore)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(partnerID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: input, partnerID, userID
func (_m *PartnerRepository) Update(input partner.Core, partnerID uint, userID uint) error {
	ret := _m.Called(input, partnerID, userID)

	var r0 error
	if rf, ok := ret.Get(0).(func(partner.Core, uint, uint) error); ok {
		r0 = rf(input, partnerID, userID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateOrderConfirmStatus provides a mock function with given fields: orderID, partnerID
func (_m *PartnerRepository) UpdateOrderConfirmStatus(orderID uint, partnerID uint) error {
	ret := _m.Called(orderID, partnerID)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint, uint) error); ok {
		r0 = rf(orderID, partnerID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdatePartnerVerifyStatus provides a mock function with given fields: verificationLog, verificationStatus, partnerID
func (_m *PartnerRepository) UpdatePartnerVerifyStatus(verificationLog string, verificationStatus string, partnerID uint) error {
	ret := _m.Called(verificationLog, verificationStatus, partnerID)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, uint) error); ok {
		r0 = rf(verificationLog, verificationStatus, partnerID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewPartnerRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewPartnerRepository creates a new instance of PartnerRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewPartnerRepository(t mockConstructorTestingTNewPartnerRepository) *PartnerRepository {
	mock := &PartnerRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
