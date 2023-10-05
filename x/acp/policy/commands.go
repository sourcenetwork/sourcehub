package policy

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	prototypes "github.com/cosmos/gogoproto/types"

	"github.com/sourcenetwork/sourcehub/x/acp/auth_engine"
	"github.com/sourcenetwork/sourcehub/x/acp/types"
)

// CreatePolicyCommand models an instruction to createa a new ACP Policy
type CreatePolicyCommand struct {
	// Cosmos Address of the Policy Creator
	CreatorAddr sdk.AccAddress

	// Policy Intermediary Representation
	Policy PolicyIR

	// Timestamp for Policy creation
	CreationTime *prototypes.Timestamp
}

// Execute consumes the data supplied in the command and creates a new ACP Policy and stores it in the given engine.
func (c *CreatePolicyCommand) Execute(ctx context.Context, accountKeeper types.AccountKeeper, engine auth_engine.AuthEngine) (*types.Policy, error) {
	factory := factory{}

	sequence, err := c.getAccountSequenceNumber(ctx, accountKeeper)
	if err != nil {
		return nil, fmt.Errorf("CreatePolicyCommand: %w", err)
	}

	record, err := factory.Create(c.Policy, string(c.CreatorAddr), sequence, c.CreationTime)
	if err != nil {
		return nil, fmt.Errorf("CreatePolicyCommand: %w", err)
	}

	spec := validPolicySpec{}
	err = spec.Satisfies(record.Policy)
	if err != nil {
		return nil, fmt.Errorf("CreatePolicyCommand: %w", err)
	}

	err = engine.SetPolicy(ctx, record)
	if err != nil {
		return nil, fmt.Errorf("CreatePolicyCommand: %w", err)
	}

	return record.Policy, nil
}

func (c *CreatePolicyCommand) getAccountSequenceNumber(ctx context.Context, accountKeeper types.AccountKeeper) (uint64, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	acc := accountKeeper.GetAccount(sdkCtx, c.CreatorAddr)
	if acc == nil {
		return 0, fmt.Errorf("account %v: %w", c.CreatorAddr, types.ErrAccNotFound)
	}

	return acc.GetSequence(), nil
}
