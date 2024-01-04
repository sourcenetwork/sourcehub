package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCheckAccess{}

func NewMsgCheckAccess(creator string, policyId string) *MsgCheckAccess {
	return &MsgCheckAccess{
		Creator:  creator,
		PolicyId: policyId,
	}
}

func (msg *MsgCheckAccess) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
