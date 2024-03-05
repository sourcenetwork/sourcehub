package relationship

import (
	"context"
	"fmt"

	"github.com/sourcenetwork/sourcehub/x/acp/auth_engine"
	"github.com/sourcenetwork/sourcehub/x/acp/policy"
	"github.com/sourcenetwork/sourcehub/x/acp/types"
)

func QueryOwnerRelationship(ctx context.Context, engine auth_engine.AuthEngine, pol *types.Policy, obj *types.Object) (*types.RelationshipRecord, error) {
	builder := types.RelationshipSelectorBuilder{}
	builder.Object(obj)
	builder.Relation(policy.OwnerRelation)
	builder.AnySubject()
	selector := builder.Build()

	records, err := engine.FilterRelationships(ctx, pol, &selector)
	if err != nil {
		return nil, fmt.Errorf("ObjectOwner: %v", err)
	}

	if len(records) == 0 {
		return nil, nil
	} else if len(records) == 1 {
		return records[0], nil
	} else {
		// Currently only Actors should be Object owners,
		// therefore if an `owner` relationship was found it must be an actor.
		// Nevertheless, in the off chance the owner isn't an Actor,
		// Return an error and log it.

		// TODO Emit metric
		//ctx.Logger().Error("invariant error: object owner isn't type actor", "policyId", req.PolicyId, "object", req.Object, "relationship", records[0])
		return nil, fmt.Errorf("ObjectOwner: object %v has non Actor type as owner: %v", obj.Id, types.ErrAcpProtocolViolation)
	}
}
