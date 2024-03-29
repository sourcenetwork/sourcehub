package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

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
	} else if rec == nil {
		return nil, fmt.Errorf("policy %v: %w", msg.PolicyId, types.ErrPolicyNotFound)
	}

	cmd := relationship.RegisterObjectCommand{
		Policy:     rec.Policy,
		CreationTs: msg.CreationTime,
		Registration: &types.Registration{
			Object: msg.Object,
			Actor: &types.Actor{
				Id: msg.Creator,
			},
		},
	}

	result, err := cmd.Execute(goCtx, engine)
	if err != nil {
		return nil, err
	}

	return &types.MsgRegisterObjectResponse{
		Result: result,
	}, nil
}
