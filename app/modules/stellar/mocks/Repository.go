// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import horizon "github.com/stellar/go/protocols/horizon"
import mock "github.com/stretchr/testify/mock"

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// LoadAccount provides a mock function with given fields: stellarAddress
func (_m *Repository) LoadAccount(stellarAddress string) (*horizon.Account, error) {
	ret := _m.Called(stellarAddress)

	var r0 *horizon.Account
	if rf, ok := ret.Get(0).(func(string) *horizon.Account); ok {
		r0 = rf(stellarAddress)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*horizon.Account)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(stellarAddress)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SubmitTransaction provides a mock function with given fields: txB64
func (_m *Repository) SubmitTransaction(txB64 string) (*horizon.TransactionSuccess, error) {
	ret := _m.Called(txB64)

	var r0 *horizon.TransactionSuccess
	if rf, ok := ret.Get(0).(func(string) *horizon.TransactionSuccess); ok {
		r0 = rf(txB64)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*horizon.TransactionSuccess)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(txB64)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
