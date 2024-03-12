package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sourcenetwork/sourcehub/x/acp/access_decision"
	"github.com/sourcenetwork/sourcehub/x/acp/types"
)

func (k msgServer) CheckAccess(goCtx context.Context, msg *types.MsgCheckAccess) (*types.MsgCheckAccessResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	eventManager := ctx.EventManager()

	repository := k.GetAccessDecisionRepository(ctx)
	paramsRepository := access_decision.StaticParamsRepository{}
	engine, err := k.GetZanziEngine(ctx)
	if err != nil {
		return nil, err
	}

	record, err := engine.GetPolicy(goCtx, msg.PolicyId)
	if err != nil {
		return nil, err
	}
	if record == nil {
		return nil, fmt.Errorf("policy %v: %w", msg.PolicyId, types.ErrPolicyNotFound)
	}

	creatorAddr, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, types.ErrInvalidAccAddr
	}
	creatorAcc := k.accountKeeper.GetAccount(ctx, creatorAddr)
	if creatorAcc == nil {
		return nil, types.ErrAccNotFound
	}

	cmd := access_decision.EvaluateAccessRequestsCommand{
		Policy:        record.Policy,
		Operations:    msg.AccessRequest.Operations,
		Actor:         msg.AccessRequest.Actor.Id,
		CreationTime:  msg.CreationTime,
		Creator:       creatorAcc,
		CurrentHeight: uint64(ctx.BlockHeight()),
	}
	decision, err := cmd.Execute(goCtx, engine, repository, &paramsRepository)
	if err != nil {
		return nil, err
	}

	err = eventManager.EmitTypedEvent(&types.EventAccessDecisionCreated{
		Creator:    msg.Creator,
		PolicyId:   msg.PolicyId,
		DecisionId: decision.Id,
		Actor:      decision.Actor,
	})
	if err != nil {
		return nil, err
	}

	return &types.MsgCheckAccessResponse{
		Decision: decision,
	}, nil
}
