package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCreateRelationship = "create_relationship"

var _ sdk.Msg = &MsgCreateRelationship{}

func NewMsgCreateRelationship(creator string, creatorDid string, policyId string, obj *Entity, relation string, actor *Entity, actorRel string) *MsgCreateRelationship {
	return &MsgCreateRelationship{
		Creator:       creator,
		CreatorDid:    creatorDid,
		PolicyId:      policyId,
		Object:        obj,
		Relation:      relation,
		Actor:         actor,
		ActorRelation: actorRel,
	}
}

func (msg *MsgCreateRelationship) Route() string {
	return RouterKey
}

func (msg *MsgCreateRelationship) Type() string {
	return TypeMsgCreateRelationship
}

func (msg *MsgCreateRelationship) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateRelationship) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateRelationship) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
