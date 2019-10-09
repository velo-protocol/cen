// Code generated by MockGen. DO NOT EDIT.
// Source: ./app/layers/repositories/stellar/interface.go

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	decimal "github.com/shopspring/decimal"
	horizon "github.com/stellar/go/protocols/horizon"
	entities "gitlab.com/velo-labs/cen/node/app/entities"
	reflect "reflect"
)

// MockStellarRepo is a mock of Repo interface
type MockStellarRepo struct {
	ctrl     *gomock.Controller
	recorder *MockStellarRepoMockRecorder
}

// MockStellarRepoMockRecorder is the mock recorder for MockStellarRepo
type MockStellarRepoMockRecorder struct {
	mock *MockStellarRepo
}

// NewMockStellarRepo creates a new mock instance
func NewMockStellarRepo(ctrl *gomock.Controller) *MockStellarRepo {
	mock := &MockStellarRepo{ctrl: ctrl}
	mock.recorder = &MockStellarRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStellarRepo) EXPECT() *MockStellarRepoMockRecorder {
	return m.recorder
}

// GetAccount mocks base method
func (m *MockStellarRepo) GetAccount(stellarAddress string) (*horizon.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccount", stellarAddress)
	ret0, _ := ret[0].(*horizon.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccount indicates an expected call of GetAccount
func (mr *MockStellarRepoMockRecorder) GetAccount(stellarAddress interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccount", reflect.TypeOf((*MockStellarRepo)(nil).GetAccount), stellarAddress)
}

// GetAccounts mocks base method
func (m *MockStellarRepo) GetAccounts(stellarAddresses ...string) ([]horizon.Account, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range stellarAddresses {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetAccounts", varargs...)
	ret0, _ := ret[0].([]horizon.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccounts indicates an expected call of GetAccounts
func (mr *MockStellarRepoMockRecorder) GetAccounts(stellarAddresses ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccounts", reflect.TypeOf((*MockStellarRepo)(nil).GetAccounts), stellarAddresses...)
}

// GetAccountData mocks base method
func (m *MockStellarRepo) GetAccountData(stellarAddress string) (map[string]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccountData", stellarAddress)
	ret0, _ := ret[0].(map[string]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccountData indicates an expected call of GetAccountData
func (mr *MockStellarRepoMockRecorder) GetAccountData(stellarAddress interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccountData", reflect.TypeOf((*MockStellarRepo)(nil).GetAccountData), stellarAddress)
}

// GetAccountDecodedData mocks base method
func (m *MockStellarRepo) GetAccountDecodedData(stellarAddress string) (map[string]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccountDecodedData", stellarAddress)
	ret0, _ := ret[0].(map[string]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccountDecodedData indicates an expected call of GetAccountDecodedData
func (mr *MockStellarRepoMockRecorder) GetAccountDecodedData(stellarAddress interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccountDecodedData", reflect.TypeOf((*MockStellarRepo)(nil).GetAccountDecodedData), stellarAddress)
}

// GetAccountDecodedDataByKey mocks base method
func (m *MockStellarRepo) GetAccountDecodedDataByKey(stellarAddress, key string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccountDecodedDataByKey", stellarAddress, key)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccountDecodedDataByKey indicates an expected call of GetAccountDecodedDataByKey
func (mr *MockStellarRepoMockRecorder) GetAccountDecodedDataByKey(stellarAddress, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccountDecodedDataByKey", reflect.TypeOf((*MockStellarRepo)(nil).GetAccountDecodedDataByKey), stellarAddress, key)
}

// GetDrsAccountData mocks base method
func (m *MockStellarRepo) GetDrsAccountData() (*entities.DrsAccountData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDrsAccountData")
	ret0, _ := ret[0].(*entities.DrsAccountData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDrsAccountData indicates an expected call of GetDrsAccountData
func (mr *MockStellarRepoMockRecorder) GetDrsAccountData() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDrsAccountData", reflect.TypeOf((*MockStellarRepo)(nil).GetDrsAccountData))
}

// GetMedianPriceFromPriceAccount mocks base method
func (m *MockStellarRepo) GetMedianPriceFromPriceAccount(priceAccountAddress string) (decimal.Decimal, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMedianPriceFromPriceAccount", priceAccountAddress)
	ret0, _ := ret[0].(decimal.Decimal)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMedianPriceFromPriceAccount indicates an expected call of GetMedianPriceFromPriceAccount
func (mr *MockStellarRepoMockRecorder) GetMedianPriceFromPriceAccount(priceAccountAddress interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMedianPriceFromPriceAccount", reflect.TypeOf((*MockStellarRepo)(nil).GetMedianPriceFromPriceAccount), priceAccountAddress)
}

// SubmitTransaction mocks base method
func (m *MockStellarRepo) SubmitTransaction(txB64 string) (*horizon.TransactionSuccess, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubmitTransaction", txB64)
	ret0, _ := ret[0].(*horizon.TransactionSuccess)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SubmitTransaction indicates an expected call of SubmitTransaction
func (mr *MockStellarRepoMockRecorder) SubmitTransaction(txB64 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubmitTransaction", reflect.TypeOf((*MockStellarRepo)(nil).SubmitTransaction), txB64)
}
