package policy_cmd

import (
	"context"
	"fmt"

	"github.com/sourcenetwork/sourcehub/x/acp/did"
	"github.com/sourcenetwork/sourcehub/x/acp/types"
)

// payloadSpec executes validation against a PolicyCmdPayload to ensure it should be accepted
func payloadSpec(params types.Params, currentHeight uint64, payload *types.PolicyCmdPayload) error {
	if payload.ExpirationDelta > params.PolicyCommandMaxExpirationDelta {
		return fmt.Errorf("%w: max %v, given %v", ErrExpirationDeltaTooLarge, params.PolicyCommandMaxExpirationDelta, payload.ExpirationDelta)
	}

	maxHeight := payload.IssuedHeight + payload.ExpirationDelta
	if maxHeight > currentHeight {
		return fmt.Errorf("%v: current %v limit %v", ErrCommandExpired, currentHeight, maxHeight)
	}

	// TODO check payload id is in cache
	return nil
}

// ValidateAndExtractCmd validates a MsgPolicyCmd and return the Cmd payload
func ValidateAndExtractCmd(ctx context.Context, params types.Params, resolver did.Resolver, payload string, contentType types.MsgPolicyCmd_ContentType, currentHeight uint64) (*types.PolicyCmdPayload, error) {
	var cmd *types.PolicyCmdPayload
	var err error

	switch contentType {
	case types.MsgPolicyCmd_JWS:
		verifier := newJWSVerifier(resolver)
		cmd, err = verifier.Verify(ctx, payload)
	default:
		err = fmt.Errorf("invalid signed command: cmd %v: %w", payload, types.ErrInvalidVariant)
	}

	if err != nil {
		return nil, fmt.Errorf("invalid signed command: %w", err)
	}

	err = payloadSpec(params, currentHeight, cmd)
	if err != nil {
		return nil, fmt.Errorf("invalid payload: %v", err)
	}

	return cmd, nil
}
