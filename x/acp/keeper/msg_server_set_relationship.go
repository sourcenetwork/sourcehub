package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sourcenetwork/sourcehub/x/acp/relationship"
	"github.com/sourcenetwork/sourcehub/x/acp/types"
)

func (k msgServer) SetRelationship(goCtx context.Context, msg *types.MsgSetRelationship) (*types.MsgSetRelationshipResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	engine, err := k.GetZanziEngine(ctx)
	if err != nil {
		return nil, err
	}

	rec, err := engine.GetPolicy(goCtx, msg.PolicyId)
	if err != nil {
		return nil, err
	} else if rec == nil {
		return nil, fmt.Errorf("policy %v: %w", msg.PolicyId, types.ErrPolicyNotFound)
	}

	authorizer := relationship.NewRelationshipAuthorizer(engine)

	command := relationship.SetRelationshipCommand{
		Policy:       rec.Policy,
		Creator:      msg.Creator,
		CreationTs:   msg.CreationTime,
		Relationship: msg.Relationship,
	}
	found, err := command.Execute(goCtx, engine, authorizer)
	if err != nil {
		return nil, err
	}

	return &types.MsgSetRelationshipResponse{
		RecordExisted: bool(found),
	}, nil
}
