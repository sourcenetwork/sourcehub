package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sourcenetwork/sourcehub/x/acp/did"
	"github.com/sourcenetwork/sourcehub/x/acp/relationship"
	"github.com/sourcenetwork/sourcehub/x/acp/types"
)

func (k msgServer) RegisterObject(goCtx context.Context, msg *types.MsgRegisterObject) (*types.MsgRegisterObjectResponse, error) {
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
		return nil, fmt.Errorf("MsgRegisterObject: policy %v: %w", msg.PolicyId, types.ErrPolicyNotFound)
	}

	addr, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, fmt.Errorf("MsgRegisterObject: invalid creator: %v", err)
	}

	account := k.accountKeeper.GetAccount(ctx, addr)
	if account == nil {
		return nil, fmt.Errorf("MsgRegisterObject: %w", types.ErrAccNotFound)
	}

	accDID, err := did.IssueDID(account)
	if err != nil {
		return nil, fmt.Errorf("MsgRegisterObject: %w", err)
	}

	cmd := relationship.RegisterObjectCommand{
		Policy:     rec.Policy,
		CreationTs: msg.CreationTime,
		Creator:    msg.Creator,
		Registration: &types.Registration{
			Object: msg.Object,
			Actor: &types.Actor{
				Id: accDID,
			},
		},
	}

	result, record, err := cmd.Execute(goCtx, engine, ctx.EventManager())
	if err != nil {
		return nil, err
	}

	return &types.MsgRegisterObjectResponse{
		Result: result,
		Record: record,
	}, nil
}
