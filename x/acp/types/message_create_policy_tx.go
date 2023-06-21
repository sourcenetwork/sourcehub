package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCreatePolicyTx = "create_policy_tx"

var _ sdk.Msg = &MsgCreatePolicyTx{}

func NewMsgCreatePolicyTx(creator string, thing string) *MsgCreatePolicyTx {
	return &MsgCreatePolicyTx{
		Creator: creator,
		Thing:   thing,
	}
}

func (msg *MsgCreatePolicyTx) Route() string {
	return RouterKey
}

func (msg *MsgCreatePolicyTx) Type() string {
	return TypeMsgCreatePolicyTx
}

func (msg *MsgCreatePolicyTx) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreatePolicyTx) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreatePolicyTx) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
