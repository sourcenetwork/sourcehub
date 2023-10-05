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

	Check(ctx context.Context, policy *types.Policy, request *types.AuthorizationRequest, actor *types.Actor) (bool, error)

	DeleteRelationship(ctx context.Context, policy *types.Policy, relationship *types.Relationship) (RecordFound, error)
}
