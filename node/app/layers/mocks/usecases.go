// Code generated by MockGen. DO NOT EDIT.
// Source: ./app/layers/usecases/interface.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	txnbuild "gitlab.com/velo-labs/cen/libs/txnbuild"
	entities "gitlab.com/velo-labs/cen/node/app/entities"
	errors "gitlab.com/velo-labs/cen/node/app/errors"
	reflect "reflect"
)

// MockUseCase is a mock of UseCase interface
type MockUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockUseCaseMockRecorder
}

// MockUseCaseMockRecorder is the mock recorder for MockUseCase
type MockUseCaseMockRecorder struct {
	mock *MockUseCase
}

// NewMockUseCase creates a new mock instance
func NewMockUseCase(ctrl *gomock.Controller) *MockUseCase {
	mock := &MockUseCase{ctrl: ctrl}
	mock.recorder = &MockUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUseCase) EXPECT() *MockUseCaseMockRecorder {
	return m.recorder
}

// SetupCredit mocks base method
func (m *MockUseCase) SetupCredit(ctx context.Context, veloTx *txnbuild.VeloTx) (*entities.SetupCreditOutput, errors.NodeError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetupCredit", ctx, veloTx)
	ret0, _ := ret[0].(*entities.SetupCreditOutput)
	ret1, _ := ret[1].(errors.NodeError)
	return ret0, ret1
}

// SetupCredit indicates an expected call of SetupCredit
func (mr *MockUseCaseMockRecorder) SetupCredit(ctx, veloTx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetupCredit", reflect.TypeOf((*MockUseCase)(nil).SetupCredit), ctx, veloTx)
}

// Whitelist mocks base method
func (m *MockUseCase) Whitelist(ctx context.Context, veloTx *txnbuild.VeloTx) (*entities.WhitelistOutput, errors.NodeError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Whitelist", ctx, veloTx)
	ret0, _ := ret[0].(*entities.WhitelistOutput)
	ret1, _ := ret[1].(errors.NodeError)
	return ret0, ret1
}

// Whitelist indicates an expected call of Whitelist
func (mr *MockUseCaseMockRecorder) Whitelist(ctx, veloTx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Whitelist", reflect.TypeOf((*MockUseCase)(nil).Whitelist), ctx, veloTx)
}

// UpdatePrice mocks base method
func (m *MockUseCase) UpdatePrice(ctx context.Context, veloTx *txnbuild.VeloTx) (*string, errors.NodeError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePrice", ctx, veloTx)
	ret0, _ := ret[0].(*string)
	ret1, _ := ret[1].(errors.NodeError)
	return ret0, ret1
}

// UpdatePrice indicates an expected call of UpdatePrice
func (mr *MockUseCaseMockRecorder) UpdatePrice(ctx, veloTx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePrice", reflect.TypeOf((*MockUseCase)(nil).UpdatePrice), ctx, veloTx)
}

// MintCredit mocks base method
func (m *MockUseCase) MintCredit(ctx context.Context, veloTx *txnbuild.VeloTx) (*entities.MintCreditOutput, errors.NodeError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MintCredit", ctx, veloTx)
	ret0, _ := ret[0].(*entities.MintCreditOutput)
	ret1, _ := ret[1].(errors.NodeError)
	return ret0, ret1
}

// MintCredit indicates an expected call of MintCredit
func (mr *MockUseCaseMockRecorder) MintCredit(ctx, veloTx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MintCredit", reflect.TypeOf((*MockUseCase)(nil).MintCredit), ctx, veloTx)
}

// RedeemCredit mocks base method
func (m *MockUseCase) RedeemCredit(ctx context.Context, veloTx *txnbuild.VeloTx) (*entities.RedeemCreditOutput, errors.NodeError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RedeemCredit", ctx, veloTx)
	ret0, _ := ret[0].(*entities.RedeemCreditOutput)
	ret1, _ := ret[1].(errors.NodeError)
	return ret0, ret1
}

// RedeemCredit indicates an expected call of RedeemCredit
func (mr *MockUseCaseMockRecorder) RedeemCredit(ctx, veloTx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RedeemCredit", reflect.TypeOf((*MockUseCase)(nil).RedeemCredit), ctx, veloTx)
}

// GetExchangeRate mocks base method
func (m *MockUseCase) GetExchangeRate(ctx context.Context, input *entities.GetExchangeRateInput) (*entities.GetExchangeRateOutPut, errors.NodeError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetExchangeRate", ctx, input)
	ret0, _ := ret[0].(*entities.GetExchangeRateOutPut)
	ret1, _ := ret[1].(errors.NodeError)
	return ret0, ret1
}

// GetExchangeRate indicates an expected call of GetExchangeRate
func (mr *MockUseCaseMockRecorder) GetExchangeRate(ctx, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetExchangeRate", reflect.TypeOf((*MockUseCase)(nil).GetExchangeRate), ctx, input)
}

// GetCollateralHealthCheck mocks base method
func (m *MockUseCase) GetCollateralHealthCheck(ctx context.Context) (*entities.GetCollateralHealthCheckOutput, errors.NodeError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCollateralHealthCheck", ctx)
	ret0, _ := ret[0].(*entities.GetCollateralHealthCheckOutput)
	ret1, _ := ret[1].(errors.NodeError)
	return ret0, ret1
}

// GetCollateralHealthCheck indicates an expected call of GetCollateralHealthCheck
func (mr *MockUseCaseMockRecorder) GetCollateralHealthCheck(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCollateralHealthCheck", reflect.TypeOf((*MockUseCase)(nil).GetCollateralHealthCheck), ctx)
}

// RebalanceReserve mocks base method
func (m *MockUseCase) RebalanceReserve(ctx context.Context, veloTx *txnbuild.VeloTx) (*entities.RebalanceOutput, errors.NodeError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RebalanceReserve", ctx, veloTx)
	ret0, _ := ret[0].(*entities.RebalanceOutput)
	ret1, _ := ret[1].(errors.NodeError)
	return ret0, ret1
}

// RebalanceReserve indicates an expected call of RebalanceReserve
func (mr *MockUseCaseMockRecorder) RebalanceReserve(ctx, veloTx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RebalanceReserve", reflect.TypeOf((*MockUseCase)(nil).RebalanceReserve), ctx, veloTx)
}
