package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	prototypes "github.com/cosmos/gogoproto/types"
)

var _ sdk.Msg = &MsgRegisterObject{}

func NewMsgRegisterObject(creator string, policyId string, object *Object, creationTime *prototypes.Timestamp) *MsgRegisterObject {
	return &MsgRegisterObject{
		Creator:      creator,
		PolicyId:     policyId,
		Object:       object,
		CreationTime: creationTime,
	}
}

func NewMsgRegisterObjectNow(creator string, policyId string, object *Object) *MsgRegisterObject {
	return NewMsgRegisterObject(creator, policyId, object, prototypes.TimestampNow())
}

func (msg *MsgRegisterObject) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
