package vxdr

import (
	"github.com/stellar/go/xdr"
)

type OperationType int32

const (
	OperationTypeWhiteList   OperationType = 0
	OperationTypeSetupCredit OperationType = 1
	OperationTypePriceUpdate OperationType = 2
)

type OperationBody struct {
	Type          OperationType
	WhiteListOp   *WhiteListOp
	SetupCreditOp *SetupCreditOp
	PriceUpdateOp *PriceUpdateOp
}

type WhiteListOp struct {
	Address  xdr.AccountId
	Role     Role
	Currency Currency
}

type SetupCreditOp struct {
	PeggedValue    xdr.Int64
	PeggedCurrency string
	AssetCode      string
}

type PriceUpdateOp struct {
	Asset                       string
	Currency                    Currency
	PriceInCurrencyPerAssetUnit xdr.Int64
}
