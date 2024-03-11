package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sourcenetwork/sourcehub/x/acp/auth_engine"
	"github.com/sourcenetwork/sourcehub/x/acp/policy_cmd"
	"github.com/sourcenetwork/sourcehub/x/acp/relationship"
	"github.com/sourcenetwork/sourcehub/x/acp/types"
)

func (k msgServer) PolicyCmd(goCtx context.Context, msg *types.MsgPolicyCmd) (*types.MsgPolicyCmdResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	engine, err := k.GetZanziEngine(ctx)
	if err != nil {
		return nil, err
	}

	payload, err := policy_cmd.ValidateAndExtractCmd(ctx, *msg.SignedCmd, uint64(ctx.BlockHeight()))
	if err != nil {
		return nil, fmt.Errorf("PolicyCmd: %w", err)
	}

	authorizer := relationship.NewRelationshipAuthorizer(engine)

	rec, err := engine.GetPolicy(goCtx, payload.PolicyId)
	if err != nil {
		return nil, err
	} else if rec == nil {
		return nil, fmt.Errorf("PolcyCmd: policy %v: %w", payload.PolicyId, types.ErrPolicyNotFound)
	}

	result := new(types.PolicyCmdResult)
	policy := rec.Policy

	switch c := payload.Cmd.(type) {
	case *types.PolicyCmdPayload_SetRelationshipCmd:
		var found auth_engine.RecordFound

		cmd := relationship.SetRelationshipCommand{
			Policy:       policy,
			CreationTs:   payload.CreationTime,
			Creator:      payload.Actor,
			Relationship: c.SetRelationshipCmd.Relationship,
		}
		found, _, err = cmd.Execute(ctx, engine, authorizer)

		result.Result = &types.PolicyCmdResult_SetRelationshipResult{
			SetRelationshipResult: &types.SetRelationshipCmdResult{
				RecordExisted: bool(found),
			},
		}
	case *types.PolicyCmdPayload_DeleteRelationshipCmd:
		var found auth_engine.RecordFound

		cmd := relationship.DeleteRelationshipCommand{
			Policy:       policy,
			Actor:        payload.Actor,
			Relationship: c.DeleteRelationshipCmd.Relationship,
		}
		found, err = cmd.Execute(ctx, engine, authorizer)

		result.Result = &types.PolicyCmdResult_DeleteRelationshipResult{
			DeleteRelationshipResult: &types.DeleteRelationshipCmdResult{
				RecordFound: bool(found),
			},
		}
	case *types.PolicyCmdPayload_RegisterObjectCmd:
		var registrationResult types.RegistrationResult
		var record *types.RelationshipRecord

		cmd := relationship.RegisterObjectCommand{
			Policy:     policy,
			CreationTs: payload.CreationTime,
			Registration: &types.Registration{
				Object: c.RegisterObjectCmd.Object,
				Actor: &types.Actor{
					Id: payload.Actor,
				},
			},
		}
		registrationResult, record, err = cmd.Execute(ctx, engine, ctx.EventManager())

		result.Result = &types.PolicyCmdResult_RegisterObjectResult{
			RegisterObjectResult: &types.RegisterObjectCmdResult{
				Result: registrationResult,
				Record: record,
			},
		}
	case *types.PolicyCmdPayload_UnregisterObjectCmd:
		var count uint

		cmd := relationship.UnregisterObjectCommand{
			Policy: policy,
			Object: c.UnregisterObjectCmd.Object,
			Actor:  payload.Actor,
		}
		count, err = cmd.Execute(ctx, engine, authorizer)

		result.Result = &types.PolicyCmdResult_UnregisterObjectResult{
			UnregisterObjectResult: &types.UnregisterObjectCmdResult{
				Found:                true, //TODO true,
				RelationshipsRemoved: uint64(count),
			},
		}

	default:
		err = fmt.Errorf("PolicyCmd: unsuported command %v: %w", c, types.ErrInvalidVariant)
	}

	if err != nil {
		return nil, fmt.Errorf("PolicyCmd failed: %v", err)

	}

	return &types.MsgPolicyCmdResponse{
		Result: result,
	}, nil
}
