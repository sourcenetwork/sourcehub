package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sourcenetwork/sourcehub/x/acp/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) Policy(goCtx context.Context, req *types.QueryPolicyRequest) (*types.QueryPolicyResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	engine, err := k.GetZanziEngine(ctx)
	if err != nil {
		return nil, err
	}

	rec, err := engine.GetPolicy(goCtx, req.Id)
	if err != nil {
		return nil, err
	}
	if rec == nil {
		return nil, fmt.Errorf("id %v: %w", req.Id, types.ErrPolicyNotFound)
	}

	return &types.QueryPolicyResponse{
		Policy: rec.Policy,
	}, nil
}
