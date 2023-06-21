package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCreatePolicy = "create_policy"

var _ sdk.Msg = &MsgCreatePolicy{}

func NewMsgCreatePolicy(creator string, thing string) *MsgCreatePolicy {
	return &MsgCreatePolicy{
		Creator: creator,
		Thing:   thing,
	}
}

func (msg *MsgCreatePolicy) Route() string {
	return RouterKey
}

func (msg *MsgCreatePolicy) Type() string {
	return TypeMsgCreatePolicy
}

func (msg *MsgCreatePolicy) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreatePolicy) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreatePolicy) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
