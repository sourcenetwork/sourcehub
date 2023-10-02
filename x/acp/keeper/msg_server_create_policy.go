package keeper

import (
	"context"

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

	cmd := policy.CreatePolicyCommand{
		CreatorAddr:  addr,
		Policy:       msg.Policy,
		MarshalType:  msg.MarshalType,
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
