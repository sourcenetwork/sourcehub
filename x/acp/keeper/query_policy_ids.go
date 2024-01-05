package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sourcenetwork/sourcehub/x/acp/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) PolicyIds(goCtx context.Context, req *types.QueryPolicyIdsRequest) (*types.QueryPolicyIdsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	engine, err := k.GetZanziEngine(ctx)
	if err != nil {
		return nil, err
	}

	ids, err := engine.ListPolicyIds(goCtx)
	if err != nil {
		return nil, err
	}

	return &types.QueryPolicyIdsResponse{
		Ids: ids,
	}, nil
}
