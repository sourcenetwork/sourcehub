package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgRegisterObject = "register_object"

var _ sdk.Msg = &MsgRegisterObject{}

func NewMsgRegisterObject(creator string, creatorDid string, policyId string, obj *Entity) *MsgRegisterObject {
	return &MsgRegisterObject{
		Creator: creator,
                CreatorDid: creatorDid,
                PolicyId: policyId,
                Object: obj,
	}
}

func (msg *MsgRegisterObject) Route() string {
	return RouterKey
}

func (msg *MsgRegisterObject) Type() string {
	return TypeMsgRegisterObject
}

func (msg *MsgRegisterObject) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRegisterObject) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRegisterObject) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
