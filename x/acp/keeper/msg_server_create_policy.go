package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sourcenetwork/sourcehub/utils"
	"github.com/sourcenetwork/sourcehub/x/acp/policy"
	"github.com/sourcenetwork/sourcehub/x/acp/types"
)

func (k msgServer) CreatePolicy(goCtx context.Context, msg *types.MsgCreatePolicy) (*types.MsgCreatePolicyResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	ctx = utils.WithMsgSpan(ctx)
	defer utils.FinalizeSpan(ctx)

	eventManager := ctx.EventManager()

	engine, err := k.GetZanziEngine(ctx)
	if err != nil {
		return nil, err
	}

	ir, err := policy.Unmarshal(msg.Policy, msg.MarshalType)
	if err != nil {
		return nil, fmt.Errorf("CreatePolicy: %w", err)
	}

	cmd := policy.CreatePolicyCommand{
		Creator:      msg.Creator,
		Policy:       ir,
		CreationTime: msg.CreationTime,
	}
	pol, err := cmd.Execute(ctx, k.accountKeeper, engine)

	if err != nil {
		return nil, err
	}

	event := types.EventPolicyCreated{
		Creator:    msg.Creator,
		PolicyId:   pol.Id,
		PolicyName: pol.Name,
	}
	err = eventManager.EmitTypedEvent(&event)
	if err != nil {
		return nil, err
	}

	return &types.MsgCreatePolicyResponse{
		Policy: pol,
	}, nil
}
