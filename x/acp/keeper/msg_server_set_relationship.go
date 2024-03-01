package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sourcenetwork/sourcehub/x/acp/did"
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
	}
	if rec == nil {
		return nil, fmt.Errorf("MsgSetRelationship: policy %v: %w", msg.PolicyId, types.ErrPolicyNotFound)
	}

	addr, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, fmt.Errorf("MsgSetRelationship: invalid creator: %v", err)
	}

	account := k.accountKeeper.GetAccount(ctx, addr)
	if account == nil {
		return nil, fmt.Errorf("MsgSetRelationship: %w", types.ErrAccNotFound)
	}

	accDID, err := did.IssueDID(account)
	if err != nil {
		return nil, fmt.Errorf("MsgSetRelationship: %w", err)
	}

	authorizer := relationship.NewRelationshipAuthorizer(engine)

	cmd := relationship.SetRelationshipCommand{
		Policy:       rec.Policy,
		Creator:      accDID,
		CreationTs:   msg.CreationTime,
		Relationship: msg.Relationship,
	}
	found, err := cmd.Execute(goCtx, engine, authorizer)
	if err != nil {
		return nil, err
	}

	return &types.MsgSetRelationshipResponse{
		RecordExisted: bool(found),
		// TODO return record
	}, nil
}
