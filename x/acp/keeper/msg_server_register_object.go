package keeper

import (
	"context"

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
	}

	cmd := relationship.RegisterObjectCommand{
		Registration: msg.Registration,
		Policy:       rec.Policy,
		CreationTs:   msg.CreationTime,
	}

	result, err := cmd.Execute(goCtx, engine)
	if err != nil {
		return nil, err
	}

	return &types.MsgRegisterObjectResponse{
		Result: result,
	}, nil
}
