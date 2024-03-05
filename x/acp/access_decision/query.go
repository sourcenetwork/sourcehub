package access_decision

import (
	"context"
	"fmt"

	"github.com/sourcenetwork/sourcehub/x/acp/auth_engine"
	"github.com/sourcenetwork/sourcehub/x/acp/types"
)

// VerifyAccessRequest verifies whether the given AccessRequest is valid for Policy.
// An AccessRequest is valid if the Request's Actor is authorized to
// execute all the Operations within it.
type VerifyAccessRequestQuery struct {
	Policy        *types.Policy
	AccessRequest *types.AccessRequest
}

// Execute runs the Comand for the given context and engine
func (c *VerifyAccessRequestQuery) Execute(ctx context.Context, engine auth_engine.AuthEngine) (bool, error) {
	err := c.validate()
	if err != nil {
		return false, fmt.Errorf("verify access request: %w", err)
	}

	actor := c.AccessRequest.Actor
	for _, operation := range c.AccessRequest.Operations {
		isAllowed, err := engine.Check(ctx, c.Policy, operation, actor)
		if err != nil {
			return false, err
		} else if !isAllowed {
			return false, nil
		}
	}
	return true, nil
}

func (c *VerifyAccessRequestQuery) validate() error {
	if c.Policy == nil {
		return types.ErrPolicyNil
	}

	if c.AccessRequest == nil {
		return types.ErrAccessRequestNil
	}

	return nil
}
