package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgPolicyCmd{}

func NewMsgPolicyCmd(creator string) *MsgPolicyCmd {
	return &MsgPolicyCmd{
		Creator: creator,
	}
}

func NewMsgPolicyCmdFromJWS(creator string, jws string) *MsgPolicyCmd {
	return &MsgPolicyCmd{
		Creator: creator,
		SignedCmd: &MsgPolicyCmd_SignedCmd{
			Payload: &MsgPolicyCmd_SignedCmd_Jws{
				Jws: jws,
			},
		},
	}
}

func (msg *MsgPolicyCmd) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
