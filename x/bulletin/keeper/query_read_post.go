package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/sourcenetwork/sourcehub/x/bulletin/types"
)

func (k Keeper) ReadPost(goCtx context.Context, req *types.QueryReadPostRequest) (*types.QueryReadPostResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	post, found := k.GetPost(ctx, req.Namespace)
	if !found {
		return nil, status.Error(codes.InvalidArgument, "not found")
	}

	return &types.QueryReadPostResponse{Post: &post}, nil
}
