package access_decision

import (
	"context"
	"fmt"

	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	prototypes "github.com/cosmos/gogoproto/types"

	"github.com/sourcenetwork/sourcehub/x/acp/auth_engine"
	"github.com/sourcenetwork/sourcehub/x/acp/types"
)

// DefaultExpirationDelta sets the number of blocks a Decision is valid for
const DefaultExpirationDelta uint64 = 100

type EvaluateAccessRequestsCommand struct {
	Policy     *types.Policy
	Operations []*types.Operation
	Actor      string

	CreationTime *prototypes.Timestamp

	// Creator is the same as the Tx signer
	Creator authtypes.AccountI

	// Current block height
	CurrentHeight uint64

	params *types.DecisionParams
}

func (c *EvaluateAccessRequestsCommand) Execute(ctx context.Context, engine auth_engine.AuthEngine, repository Repository, paramsRepo ParamsRepository) (*types.AccessDecision, error) {
	err := c.validate()
	if err != nil {
		return nil, fmt.Errorf("EvaluateAccessRequest: %w", err)
	}

	err = c.evaluateRequest(ctx, engine)
	if err != nil {
		return nil, fmt.Errorf("EvaluateAccessRequest: %w", err)
	}

	c.params, err = paramsRepo.GetDefaults(ctx)
	if err != nil {
		return nil, fmt.Errorf("EvaluateAccessRequest: %w", err)
	}

	decision := c.buildDecision()

	err = repository.Set(ctx, decision)
	if err != nil {
		return nil, fmt.Errorf("EvaluateAccessRequest: %w", err)
	}

	return decision, nil
}

func (c *EvaluateAccessRequestsCommand) validate() error {
	if c.Policy == nil {
		return types.ErrPolicyNil
	}

	if c.Operations == nil {
		return types.ErrAccessRequestNil
	}

	if c.CurrentHeight == 0 {
		return types.ErrInvalidHeight
	}

	return nil
}

func (c *EvaluateAccessRequestsCommand) evaluateRequest(ctx context.Context, engine auth_engine.AuthEngine) error {
	actor := types.Actor{
		Id: c.Actor,
	}

	for _, operation := range c.Operations {
		isAllowed, err := engine.Check(ctx, c.Policy, operation, &actor)
		if err != nil {
			return err
		} else if !isAllowed {
			return fmt.Errorf("actor %v: operation %v: %w", actor, operation, types.ErrNotAuthorized)
		}
	}

	return nil
}

func (c *EvaluateAccessRequestsCommand) buildDecision() *types.AccessDecision {
	decision := &types.AccessDecision{
		PolicyId:           c.Policy.Id,
		Params:             c.params,
		CreationTime:       c.CreationTime,
		Operations:         c.Operations,
		IssuedHeight:       c.CurrentHeight,
		Actor:              c.Actor,
		Creator:            c.Creator.GetAddress().String(),
		CreatorAccSequence: c.Creator.GetSequence(),
	}
	decision.Id = decision.ProduceId()
	return decision
}
