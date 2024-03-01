package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sourcenetwork/sourcehub/x/acp/did"
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
	if rec == nil {
		return nil, fmt.Errorf("MsgDeleteRelationship: policy %v: %w", msg.PolicyId, types.ErrPolicyNotFound)
	}

	addr, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, fmt.Errorf("MsgDeleteRelationship: invalid creator: %v", err)
	}

	account := k.accountKeeper.GetAccount(ctx, addr)
	if account == nil {
		return nil, fmt.Errorf("MsgDeleteRelationship: %w", types.ErrAccNotFound)
	}

	accDID, err := did.IssueDID(account)
	if err != nil {
		return nil, fmt.Errorf("MsgDeleteRelationship: %w", err)
	}

	command := relationship.DeleteRelationshipCommand{
		Policy:       rec.Policy,
		Relationship: msg.Relationship,
		Actor:        accDID,
	}
	found, err := command.Execute(ctx, engine, authorizer)
	if err != nil {
		return nil, err
	}

	return &types.MsgDeleteRelationshipResponse{
		RecordFound: bool(found),
	}, nil
}
