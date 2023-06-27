package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgDeletePolicy = "delete_policy"

var _ sdk.Msg = &MsgDeletePolicy{}

func NewMsgDeletePolicy(creator string, id string) *MsgDeletePolicy {
	return &MsgDeletePolicy{
		Creator: creator,
		Id:   id,
	}
}

func (msg *MsgDeletePolicy) Route() string {
	return RouterKey
}

func (msg *MsgDeletePolicy) Type() string {
	return TypeMsgDeletePolicy
}

func (msg *MsgDeletePolicy) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeletePolicy) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeletePolicy) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
