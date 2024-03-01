package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgUnregisterObject{}

func NewMsgUnregisterObject(creator string, policyId string, object *Object) *MsgUnregisterObject {
	return &MsgUnregisterObject{
		Creator:  creator,
		PolicyId: policyId,
		Object:   object,
	}
}

func (msg *MsgUnregisterObject) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
