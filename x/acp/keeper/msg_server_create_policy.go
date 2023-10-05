package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sourcenetwork/sourcehub/x/acp/policy"
	"github.com/sourcenetwork/sourcehub/x/acp/types"
)

func (k msgServer) CreatePolicy(goCtx context.Context, msg *types.MsgCreatePolicy) (*types.MsgCreatePolicyResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	engine, err := k.GetZanziEngine(ctx)
	if err != nil {
		return nil, err
	}

	// We can safely assume Creator exists otherwise
	// antee handler would've errored out
	addr := sdk.MustAccAddressFromBech32(msg.Creator)

	ir, err := policy.Unmarshal(msg.Policy, msg.MarshalType)
	if err != nil {
		return nil, fmt.Errorf("CreatePolicy: %w", err)
	}

	cmd := policy.CreatePolicyCommand{
		CreatorAddr:  addr,
		Policy:       ir,
		CreationTime: msg.CreationTime,
	}
	pol, err := cmd.Execute(goCtx, k.accountKeeper, engine)

	if err != nil {
		return nil, err
	}

	return &types.MsgCreatePolicyResponse{
		Policy: pol,
	}, nil
}
