package types

// DONTCOVER

import (
	sdkerrors "cosmossdk.io/errors"
)

// x/acp module sentinel errors
var (
	ErrInvalidSigner = sdkerrors.Register(ModuleName, 1100, "expected gov account as only signer for proposal message")
	ErrSample        = sdkerrors.Register(ModuleName, 1101, "sample error")

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

	ErrPolicyNil        = ErrAcpInput.Wrapf("policy must not be nil")
	ErrRelationshipNil  = ErrAcpInput.Wrapf("relationship must not be nil")
	ErrActorNil         = ErrAcpInput.Wrapf("actor must not be nil")
	ErrRegistrationNil  = ErrAcpInput.Wrapf("registration must not be nil")
	ErrAccessRequestNil = ErrAcpInput.Wrapf("AccessRequest must not be nil")
	ErrInvalidVariant   = ErrAcpInput.Wrapf("invalid type variant")
	ErrObjectNil        = ErrAcpInput.Wrapf("object must not be nil")
	ErrTimestampNil     = ErrAcpInput.Wrapf("timestamp must not be nil")
	ErrAccNotFound      = ErrAcpInput.Wrapf("account not found")
	ErrPolicyNotFound   = ErrAcpInput.Wrapf("policy not found")
	ErrObjectNotFound   = ErrAcpInput.Wrapf("object not found")
	ErrInvalidHeight    = ErrAcpInput.Wrapf("invalid block height")
	ErrInvalidAccAddr   = ErrAcpInput.Wrapf("invalid account address")
	ErrInvalidDID       = ErrAcpInput.Wrapf("invalid DID")

	ErrNotAuthorized = ErrAcpProtocolViolation.Wrapf("actor not authorized")
)
