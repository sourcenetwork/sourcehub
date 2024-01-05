package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sourcenetwork/sourcehub/x/acp/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) FilterRelationships(goCtx context.Context, req *types.QueryFilterRelationshipsRequest) (*types.QueryFilterRelationshipsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	engine, err := k.GetZanziEngine(ctx)
	if err != nil {
		return nil, err
	}

	rec, err := engine.GetPolicy(goCtx, req.PolicyId)
	if err != nil {
		return nil, err
	}
	if rec == nil {
		return nil, fmt.Errorf("policy %v: %w", req.PolicyId, types.ErrPolicyNotFound)
	}

	records, err := engine.FilterRelationships(goCtx, rec.Policy, req.Selector)
	if err != nil {
		return nil, err
	}

	return &types.QueryFilterRelationshipsResponse{
		Records: records,
	}, nil
}
