package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgUnregisterObject = "unregister_object"

var _ sdk.Msg = &MsgUnregisterObject{}

func NewMsgUnregisterObject(creator string, thing string) *MsgUnregisterObject {
	return &MsgUnregisterObject{
		Creator: creator,
		Thing:   thing,
	}
}

func (msg *MsgUnregisterObject) Route() string {
	return RouterKey
}

func (msg *MsgUnregisterObject) Type() string {
	return TypeMsgUnregisterObject
}

func (msg *MsgUnregisterObject) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUnregisterObject) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUnregisterObject) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
