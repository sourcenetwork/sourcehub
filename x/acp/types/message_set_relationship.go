package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSetRelationship = "set_relationship"

var _ sdk.Msg = &MsgSetRelationship{}

func NewMsgSetRelationship(creator string, policyId string, relationship *Relationship) *MsgSetRelationship {
	return &MsgSetRelationship{
		Creator:      creator,
		PolicyId:     policyId,
		Relationship: relationship,
	}
}

func (msg *MsgSetRelationship) Route() string {
	return RouterKey
}

func (msg *MsgSetRelationship) Type() string {
	return TypeMsgSetRelationship
}

func (msg *MsgSetRelationship) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSetRelationship) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSetRelationship) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
