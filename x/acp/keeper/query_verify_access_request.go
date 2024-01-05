package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/sourcenetwork/sourcehub/x/acp/access_decision"
	"github.com/sourcenetwork/sourcehub/x/acp/types"
)

func (k Keeper) VerifyAccessRequest(goCtx context.Context, req *types.QueryVerifyAccessRequestRequest) (*types.QueryVerifyAccessRequestResponse, error) {
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

	cmd := access_decision.VerifyAccessRequestCommand{
		Policy:        rec.Policy,
		AccessRequest: req.AccessRequest,
	}
	err = cmd.Execute(ctx, engine)
	if err != nil {
		return nil, err
	}

	return &types.QueryVerifyAccessRequestResponse{
		Valid: true,
	}, nil
}
