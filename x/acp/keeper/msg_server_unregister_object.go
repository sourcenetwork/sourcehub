package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sourcenetwork/sourcehub/x/acp/did"
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
	if rec == nil {
		return nil, fmt.Errorf("MsgUnregisterObject: policy %v: %w", msg.PolicyId, types.ErrPolicyNotFound)
	}

	addr, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, fmt.Errorf("MsgUnregisterObject: invalid creator: %v: %w", err, types.ErrAcpInput)
	}

	account := k.accountKeeper.GetAccount(ctx, addr)
	if account == nil {
		return nil, fmt.Errorf("MsgUnregisterObject: %w", types.ErrAccNotFound)
	}

	accDID, err := did.IssueDID(account)
	if err != nil {
		return nil, fmt.Errorf("MsgUnregisterObject: %w", err)
	}

	cmd := relationship.UnregisterObjectCommand{
		Policy: rec.Policy,
		Object: msg.Object,
		Actor:  accDID,
	}
	_, err = cmd.Execute(goCtx, engine, authorizer)
	if err != nil {
		return nil, fmt.Errorf("MsgUnregisterObject: %w", err)
	}

	return &types.MsgUnregisterObjectResponse{
		Found: true,
	}, nil
}
