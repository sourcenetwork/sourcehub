package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sourcenetwork/sourcehub/x/acp/policy"
	"github.com/sourcenetwork/sourcehub/x/acp/types"
)

func (k msgServer) CreatePolicy(goCtx context.Context, msg *types.MsgCreatePolicy) (*types.MsgCreatePolicyResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// We can safely assume Creator exists otherwise
	// antee handler would've errored out
	addr := sdk.MustAccAddressFromBech32(msg.Creator)
	acc := k.accountKeeper.GetAccount(ctx, addr)
	sequence := acc.GetSequence()

	pol, err := policy.NewPolicy(msg.Policy, msg.MarshalType, msg.Creator,
		sequence, msg.CreationTime)
	if err != nil {
		return nil, types.ErrPolicyInput.Wrapf("failed to create policy: %v", err)
	}

	err = k.polRepo.Set(goCtx, pol)
	if err != nil {
		return nil, types.ErrPolicyInput.Wrapf("failed to create policy: %v", err)
	}

	return &types.MsgCreatePolicyResponse{
		Id: pol.Id,
	}, nil
}
