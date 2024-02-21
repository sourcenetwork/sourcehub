package policy

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	prototypes "github.com/cosmos/gogoproto/types"

	"github.com/sourcenetwork/sourcehub/x/acp/auth_engine"
	"github.com/sourcenetwork/sourcehub/x/acp/types"
)

// CreatePolicyCommand models an instruction to createa a new ACP Policy
type CreatePolicyCommand struct {
	// Cosmos Address of the Policy Creator
	Creator string

	// Policy Intermediary Representation
	Policy PolicyIR

	// Timestamp for Policy creation
	CreationTime *prototypes.Timestamp
}

// Execute consumes the data supplied in the command and creates a new ACP Policy and stores it in the given engine.
func (c *CreatePolicyCommand) Execute(ctx sdk.Context, accountKeeper types.AccountKeeper, engine auth_engine.AuthEngine) (*types.Policy, error) {
	err := basicPolicyIRSpec(&c.Policy)
	if err != nil {
		return nil, fmt.Errorf("CreatePolicyCommand: %w", err)
	}

	sequence, err := c.getAccountSequenceNumber(ctx, accountKeeper)
	if err != nil {
		return nil, fmt.Errorf("CreatePolicyCommand: %w", err)
	}

	factory := factory{}
	record, err := factory.Create(c.Policy, c.Creator, sequence, c.CreationTime)
	if err != nil {
		return nil, fmt.Errorf("CreatePolicyCommand: %w", err)
	}

	spec := validPolicySpec{}
	err = spec.Satisfies(record.Policy)
	if err != nil {
		return nil, fmt.Errorf("CreatePolicyCommand: %w", err)
	}

	err = engine.SetPolicy(ctx.Context(), record)
	if err != nil {
		return nil, fmt.Errorf("CreatePolicyCommand: %w", err)
	}

	return record.Policy, nil
}

func (c *CreatePolicyCommand) getAccountSequenceNumber(ctx sdk.Context, accountKeeper types.AccountKeeper) (uint64, error) {
	addr, err := sdk.AccAddressFromBech32(c.Creator)
	if err != nil {
		return 0, fmt.Errorf("%w: %v", ErrInvalidCreator, err)
	}

	acc := accountKeeper.GetAccount(ctx, addr)
	if acc == nil {
		return 0, fmt.Errorf("account %v: %w", c.Creator, types.ErrAccNotFound)
	}

	return acc.GetSequence(), nil
}
