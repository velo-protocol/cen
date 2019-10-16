package grpc

import (
	"context"
	"fmt"
	spec "gitlab.com/velo-labs/cen/grpc"
	"gitlab.com/velo-labs/cen/libs/txnbuild"
	"gitlab.com/velo-labs/cen/libs/xdr"
	"gitlab.com/velo-labs/cen/node/app/constants"
	"gitlab.com/velo-labs/cen/node/app/errors"
)

func (handler *handler) SubmitVeloTx(ctx context.Context, req *spec.VeloTxRequest) (*spec.VeloTxReply, error) {
	veloTx, err := vtxnbuild.TransactionFromXDR(req.GetSignedVeloTxXdr())
	if err != nil {
		return nil, nerrors.ErrInvalidArgument{
			Message: err.Error(),
		}.GRPCError()
	}

	switch veloTx.TxEnvelope().VeloTx.VeloOp.Body.Type {
	case vxdr.OperationTypeWhitelist:
		return handler.handleWhitelistOperation(ctx, &veloTx)
	case vxdr.OperationTypeSetupCredit:
		return handler.handleSetupCreditOperation(ctx, &veloTx)
	case vxdr.OperationTypePriceUpdate:
		return handler.handlePriceUpdateOperation(ctx, &veloTx)
	case vxdr.OperationTypeMintCredit:
		return handler.handleMintCreditOperation(ctx, &veloTx)
	case vxdr.OperationTypeRedeemCredit:
		return handler.handlerRedeemCreditOperation(ctx, &veloTx)
	default: // this case should never occur, if the cen/libs and cen/node is aligned
		return nil, nerrors.ErrInvalidArgument{
			Message: constants.ErrUnknownVeloOperationType,
		}.GRPCError()
	}

}

func (handler *handler) handleWhitelistOperation(ctx context.Context, veloTx *vtxnbuild.VeloTx) (*spec.VeloTxReply, error) {
	op := veloTx.TxEnvelope().VeloTx.VeloOp.Body.WhitelistOp
	if op == nil {
		return nil, nerrors.ErrInvalidArgument{
			Message: fmt.Sprintf(constants.ErrFormatMissingOperation, constants.VeloOpWhitelist),
		}.GRPCError()
	}

	signedStellarTxXdr, err := handler.UseCase.CreateWhitelist(ctx, veloTx)
	if err != nil {
		return nil, err.GRPCError()
	}

	return &spec.VeloTxReply{
		SignedStellarTxXdr: *signedStellarTxXdr,
		Message:            fmt.Sprintf(constants.ReplyWhitelistSuccess, op.Address.Address(), vxdr.RoleMap[op.Role]),
	}, nil
}

func (handler *handler) handleSetupCreditOperation(ctx context.Context, veloTx *vtxnbuild.VeloTx) (*spec.VeloTxReply, error) {
	op := veloTx.TxEnvelope().VeloTx.VeloOp.Body.SetupCreditOp
	if op == nil {
		return nil, nerrors.ErrInvalidArgument{
			Message: fmt.Sprintf(constants.ErrFormatMissingOperation, constants.VeloOpSetupCredit),
		}.GRPCError()
	}

	signedStellarTxXdr, err := handler.UseCase.SetupCredit(ctx, veloTx)
	if err != nil {
		return nil, err.GRPCError()
	}

	return &spec.VeloTxReply{
		SignedStellarTxXdr: *signedStellarTxXdr,
		Message:            constants.ReplySetupCreditSuccess,
	}, nil
}

func (handler *handler) handlePriceUpdateOperation(ctx context.Context, veloTx *vtxnbuild.VeloTx) (*spec.VeloTxReply, error) {
	op := veloTx.TxEnvelope().VeloTx.VeloOp.Body.PriceUpdateOp
	if op == nil {
		return nil, nerrors.ErrInvalidArgument{
			Message: fmt.Sprintf(constants.ErrFormatMissingOperation, constants.VeloOpPriceUpdate),
		}.GRPCError()
	}

	signedStellarTxXdr, err := handler.UseCase.UpdatePrice(ctx, veloTx)
	if err != nil {
		return nil, err.GRPCError()
	}

	return &spec.VeloTxReply{
		SignedStellarTxXdr: *signedStellarTxXdr,
		Message:            constants.ReplyPriceUpdateSuccess,
	}, nil
}

func (handler *handler) handleMintCreditOperation(ctx context.Context, veloTx *vtxnbuild.VeloTx) (*spec.VeloTxReply, error) {
	op := veloTx.TxEnvelope().VeloTx.VeloOp.Body.MintCreditOp
	if op == nil {
		return nil, nerrors.ErrInvalidArgument{
			Message: fmt.Sprintf(constants.ErrFormatMissingOperation, constants.VeloOpMintCredit),
		}.GRPCError()
	}

	mintCreditOutput, err := handler.UseCase.MintCredit(ctx, veloTx)
	if err != nil {
		return nil, err.GRPCError()
	}

	return &spec.VeloTxReply{
		SignedStellarTxXdr: mintCreditOutput.SignedStellarTxXdr,
		Message: fmt.Sprintf(
			constants.ReplyMintCreditSuccess,
			mintCreditOutput.MintAmount.Truncate(7).StringFixed(7),
			mintCreditOutput.MintCurrency,
			mintCreditOutput.CollateralAmount.Truncate(7).StringFixed(7),
			mintCreditOutput.CollateralAsset,
		),
	}, nil
}

func (handler *handler) handlerRedeemCreditOperation(ctx context.Context, veloTx *vtxnbuild.VeloTx) (*spec.VeloTxReply, error) {
	op := veloTx.TxEnvelope().VeloTx.VeloOp.Body.RedeemCreditOp
	if op == nil {
		return nil, nerrors.ErrInvalidArgument{
			Message: fmt.Sprintf(constants.ErrFormatMissingOperation, constants.VeloOpRedeemCredit),
		}.GRPCError()
	}

	redeemCreditOutput, err := handler.UseCase.RedeemCredit(ctx, veloTx)
	if err != nil {
		return nil, err.GRPCError()
	}

	return &spec.VeloTxReply{
		SignedStellarTxXdr: redeemCreditOutput.SignedStellarTxXdr,
		Message:            constants.ReplyRedeemCreditSuccess,
	}, nil
}
