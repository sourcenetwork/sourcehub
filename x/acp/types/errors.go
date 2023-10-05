package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/acp module sentinel errors
var (
	// ErrAcpInternal is a general base error for IO or unexpected system errors
	ErrAcpInternal = sdkerrors.Register(ModuleName, 1000, "internal error")

	// ErrAcpInput is a general base error for input errors
	ErrAcpInput = sdkerrors.Register(ModuleName, 1001, "input error")

	// ErrAcpProtocolViolation is a general base error for operations forbidden by the protocol
	ErrAcpProtocolViolation = sdkerrors.Register(ModuleName, 1002, "acp protocol violation")

	// ErrAcpInvariantViolation indicates that an important condition of the protocol
	// has been violated, either by a bug or a successful exploit.
	// These are bad.
	ErrAcpInvariantViolation = sdkerrors.Register(ModuleName, 1003, "invariant violation")

	ErrPolicyNil       = ErrAcpInput.Wrapf("policy must not be nil")
	ErrRelationshipNil = ErrAcpInput.Wrapf("relationship must not be nil")
	ErrActorNil        = ErrAcpInput.Wrapf("actor must not be nil")
	ErrRegistrationNil = ErrAcpInput.Wrapf("registration must not be nil")
	ErrInvalidVariant  = ErrAcpInput.Wrapf("invalid type variant")
	ErrTimestampNil    = ErrAcpInput.Wrapf("timestamp must not be nil")
	ErrAccNotFound     = ErrAcpInput.Wrapf("account not found")

	ErrNotAuthorized = ErrAcpProtocolViolation.Wrapf("actor not authorized")
)
