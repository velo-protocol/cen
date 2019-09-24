package usecases

import (
	"context"
	"fmt"
	"github.com/AlekSi/pointer"
	vconvert "gitlab.com/velo-labs/cen/libs/convert"
	vtxnbuild "gitlab.com/velo-labs/cen/libs/txnbuild"
	vxdr "gitlab.com/velo-labs/cen/libs/xdr"
	"gitlab.com/velo-labs/cen/node/app/constants"
	"gitlab.com/velo-labs/cen/node/app/entities"
	nerrors "gitlab.com/velo-labs/cen/node/app/errors"
)

func (useCase *useCase) UpdatePrice(ctx context.Context, veloTx *vtxnbuild.VeloTx) nerrors.NodeError {

	if err := veloTx.VeloOp.Validate(); err != nil {
		return nerrors.ErrInvalidArgument{Message: err.Error()}
	}

	txSenderPublicKey := veloTx.TxEnvelope().VeloTx.SourceAccount.Address()
	txSenderKeyPair, err := vconvert.PublicKeyToKeyPair(txSenderPublicKey)
	if err != nil {
		return nerrors.ErrInvalidArgument{Message: err.Error()}
	}
	if veloTx.TxEnvelope().Signatures == nil {
		return nerrors.ErrUnAuthenticated{Message: constants.ErrSignatureNotFound}
	}
	if txSenderKeyPair.Hint() != veloTx.TxEnvelope().Signatures[0].Hint {
		return nerrors.ErrUnAuthenticated{Message: constants.ErrSignatureNotMatchSourceAccount}
	}

	trustedPartnerEntity, err := useCase.WhitelistRepo.FindOneWhitelist(entities.WhiteListFilter{
		StellarPublicKey: pointer.ToString(txSenderPublicKey),
		RoleCode:         pointer.ToString(string(vxdr.RoleTrustedPartner)),
	})
	if err != nil {
		return nerrors.ErrInternal{Message: err.Error()}
	}
	if trustedPartnerEntity == nil {
		return nerrors.ErrPermissionDenied{
			Message: fmt.Sprintf(constants.ErrFormatSignerNotHavePermission, constants.VeloOpSetupCredit),
		}
	}

	return nil
}
