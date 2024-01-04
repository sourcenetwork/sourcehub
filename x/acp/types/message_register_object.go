package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgRegisterObject{}

func NewMsgRegisterObject(creator string, policyId string) *MsgRegisterObject {
	return &MsgRegisterObject{
		Creator:  creator,
		PolicyId: policyId,
	}
}

func (msg *MsgRegisterObject) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
