package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sourcenetwork/sourcehub/x/acp/relationship"
	"github.com/sourcenetwork/sourcehub/x/acp/types"
)

func (k msgServer) DeleteRelationship(goCtx context.Context, msg *types.MsgDeleteRelationship) (*types.MsgDeleteRelationshipResponse, error) {
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

	command := relationship.DeleteRelationshipCommand{
		Policy:       rec.Policy,
		Relationship: msg.Relationship,
		Actor:        msg.Creator,
	}
	found, err := command.Execute(ctx, engine, authorizer)
	if err != nil {
		return nil, err
	}

	return &types.MsgDeleteRelationshipResponse{
		RecordFound: bool(found),
	}, nil
}
