package auth_engine

import (
	"context"

	"github.com/sourcenetwork/sourcehub/x/acp/types"
)

type RecordFound bool

// AuthEngine models an Authorization engine service provider.
// The engine is responsible for storing Policies and Relationships, as well as evaluating queries.
type AuthEngine interface {

	// Reurn a Relationship from a Policy, returns nil if Relationship does not exist
	GetRelationship(ctx context.Context, policy *types.Policy, rel *types.Relationship) (*types.RelationshipRecord, error)

	// Sets a Relationship within a Policy
	SetRelationship(ctx context.Context, policy *types.Policy, rec *types.RelationshipRecord) (RecordFound, error)

	// Returns all Relationships which matches selector
	FilterRelationships(ctx context.Context, policy *types.Policy, selector *types.RelationshipSelector) ([]*types.RelationshipRecord, error)

	// GetPolicy returns a PolicyRecord for the given id
	GetPolicy(ctx context.Context, policyId string) (*types.PolicyRecord, error)

	// SetPolicy stores a new Policy with the given Id
	SetPolicy(ctx context.Context, pol *types.PolicyRecord) error

	// ListPolicyIds returns the IDs of all known Policies
	ListPolicyIds(ctx context.Context) ([]string, error)

	// Check verifies whether an Acccess Request is allowed within a certain Policy
	Check(ctx context.Context, policy *types.Policy, request *types.Operation, actor *types.Actor) (bool, error)

	// DeleteRelationship removes a Relationship from a Policy
	DeleteRelationship(ctx context.Context, policy *types.Policy, relationship *types.Relationship) (RecordFound, error)

	// DeleteRelationships removes all Relationships matching the given selector
	DeleteRelationships(ctx context.Context, policy *types.Policy, selector *types.RelationshipSelector) (uint, error)
}
