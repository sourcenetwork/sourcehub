package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/acp module sentinel errors
var (
	ErrAcpInternal    = sdkerrors.Register(ModuleName, 1000, "internal error")
	ErrPolicyInput    = sdkerrors.Register(ModuleName, 1001, "policy input error")
	ErrInvalidVariant = sdkerrors.Register(ModuleName, 1002, "invalid type variant")
)
