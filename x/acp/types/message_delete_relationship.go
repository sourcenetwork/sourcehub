package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgDeleteRelationship = "delete_relationship"

var _ sdk.Msg = &MsgDeleteRelationship{}

func NewMsgDeleteRelationship(creator string, policyId string, obj *Object, relation string, subject *Subject) *MsgDeleteRelationship {
	return &MsgDeleteRelationship{
		Creator:  creator,
		PolicyId: policyId,
		Object:   obj,
		Relation: relation,
		Subject:  subject,
	}
}

func (msg *MsgDeleteRelationship) Route() string {
	return RouterKey
}

func (msg *MsgDeleteRelationship) Type() string {
	return TypeMsgDeleteRelationship
}

func (msg *MsgDeleteRelationship) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteRelationship) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteRelationship) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
