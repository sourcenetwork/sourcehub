package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	prototypes "github.com/cosmos/gogoproto/types"
)

var _ sdk.Msg = &MsgSetRelationship{}

func NewMsgSetRelationship(creator string, policyId string, relationship *Relationship, creationTime *prototypes.Timestamp) *MsgSetRelationship {
	return &MsgSetRelationship{
		Creator:      creator,
		PolicyId:     policyId,
		CreationTime: creationTime,
		Relationship: relationship,
	}
}

func NewMsgSetRelationshipNow(creator string, policyId string, relationship *Relationship) *MsgSetRelationship {
	return NewMsgSetRelationship(creator, policyId, relationship, prototypes.TimestampNow())
}

func (msg *MsgSetRelationship) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
