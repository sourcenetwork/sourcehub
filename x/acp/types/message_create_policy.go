package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	prototypes "github.com/cosmos/gogoproto/types"
)

var _ sdk.Msg = &MsgCreatePolicy{}

func NewMsgCreatePolicy(creator string, policy string, marshalingType PolicyMarshalingType, creationTime *prototypes.Timestamp) *MsgCreatePolicy {
	return &MsgCreatePolicy{
		Creator:      creator,
		Policy:       policy,
		MarshalType:  marshalingType,
		CreationTime: creationTime,
	}
}

// NewMsgCreatePolicyNow creates a MsgCreatePolicy with CreatedAt set to the current time
func NewMsgCreatePolicyNow(creator string, policy string, marshalingType PolicyMarshalingType) *MsgCreatePolicy {
	return NewMsgCreatePolicy(creator, policy, marshalingType, prototypes.TimestampNow())
}

func (msg *MsgCreatePolicy) ValidateBasic() error {
	// ValidateBasic should probably unmarshal the policy and validate it
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
