package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sourcenetwork/sourcehub/x/acp/policy"
	"github.com/sourcenetwork/sourcehub/x/acp/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ObjectOwner(goCtx context.Context, req *types.QueryObjectOwnerRequest) (*types.QueryObjectOwnerResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	engine, err := k.GetZanziEngine(ctx)
	if err != nil {
		return nil, err
	}

	rec, err := engine.GetPolicy(goCtx, req.PolicyId)
	if err != nil {
		return nil, err
	}
	if rec == nil {
		return nil, fmt.Errorf("policy %v: %w", req.PolicyId, types.ErrPolicyNotFound)
	}

	builder := types.RelationshipSelectorBuilder{}
	builder.Object(req.Object)
	builder.Relation(policy.OwnerRelation)
	builder.AnySubject()
	selector := builder.Build()

	records, err := engine.FilterRelationships(goCtx, rec.Policy, &selector)
	if err != nil {
		return nil, fmt.Errorf("ObjectOwner: %v", err)
	}

	response := &types.QueryObjectOwnerResponse{}

	if len(records) > 0 {
		// Currently only Actors should be Object owners,
		// therefore if an `owner` relationship was found it must be an actor.
		// Nevertheless, in the off chance the owner isn't an Actor,
		// Return an error and log it.
		actor := records[0].Relationship.Subject.GetActor()
		if actor == nil {
			// TODO Emit metric
			ctx.Logger().Error("invariant error: object owner isn't type actor", "policyId", req.PolicyId, "object", req.Object, "relationship", records[0])
			return nil, fmt.Errorf("ObjectOwner: object %v has non Actor type as owner: %v", req.Object.Id, types.ErrAcpProtocolViolation)
		}

		response.OwnerId = actor.Id
		response.IsRegistered = true
	}

	return response, nil
}
