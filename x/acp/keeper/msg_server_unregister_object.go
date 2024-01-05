package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sourcenetwork/sourcehub/x/acp/relationship"
	"github.com/sourcenetwork/sourcehub/x/acp/types"
)

func (k msgServer) UnregisterObject(goCtx context.Context, msg *types.MsgUnregisterObject) (*types.MsgUnregisterObjectResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	engine, err := k.GetZanziEngine(ctx)
	if err != nil {
		return nil, err
	}
	authorizer := relationship.NewRelationshipAuthorizer(engine)

	rec, err := engine.GetPolicy(goCtx, msg.PolicyId)
	if err != nil {
		return nil, err
	}

	command := relationship.UnregisterObjectCommand{
		Policy: rec.Policy,
		Object: msg.Object,
		Actor:  msg.Creator,
	}
	_, err = command.Execute(goCtx, engine, authorizer)
	if err != nil {
		return nil, err
	}

	return &types.MsgUnregisterObjectResponse{
		Found: true,
	}, nil
}
