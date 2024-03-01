package relationship

import (
	"context"
	"fmt"

	prototypes "github.com/cosmos/gogoproto/types"

	"github.com/sourcenetwork/sourcehub/x/acp/auth_engine"
	"github.com/sourcenetwork/sourcehub/x/acp/policy"
	"github.com/sourcenetwork/sourcehub/x/acp/types"
)

// RegisterObjectCommand creates an "owner" Relationship for the given object and subject,
// if the object does not have a previous owner.
// If the relationship exists but is archived by the same actor, unarchives it
// if relationship is active this command is a noop
type RegisterObjectCommand struct {
	Registration *types.Registration
	Policy       *types.Policy
	CreationTs   *prototypes.Timestamp
}

func (c *RegisterObjectCommand) Execute(ctx context.Context, engine auth_engine.AuthEngine) (types.RegistrationResult, *types.RelationshipRecord, error) {
	var err error
	var result types.RegistrationResult

	err = c.validate()
	if err != nil {
		return result, nil, fmt.Errorf("failed to register object: %w", err)
	}

	record, err := c.getOwnerRelationship(ctx, engine)
	if err != nil {
		return result, nil, fmt.Errorf("failed to register object: %w", err)
	}

	switch c.resolveObjectStatus(record) {
	case statusUnregistered:
		result, err = c.unregisteredStrategy(ctx, engine)
	case statusArchived:
		result, err = c.archivedObjectStrategy(ctx, engine, record)
	case statusActive:
		result, err = c.activeObjectStrategy(record)
	}

	if err != nil {
		return result, nil, fmt.Errorf("failed to register object: %w", err)
	}

	record, err = c.getOwnerRelationship(ctx, engine)
	if err != nil {
		return result, nil, fmt.Errorf("failed to register object: %w", err)
	}

	return result, record, nil
}

// validates the command input params
func (c *RegisterObjectCommand) validate() error {
	if c.Policy == nil {
		return types.ErrPolicyNil
	}

	err := registrationSpec(c.Registration)
	if err != nil {
		return err
	}

	if c.CreationTs == nil {
		return types.ErrTimestampNil
	}

	return nil
}

func (c *RegisterObjectCommand) getOwnerRelationship(ctx context.Context, engine auth_engine.AuthEngine) (*types.RelationshipRecord, error) {
	selector := &types.RelationshipSelector{
		ObjectSelector: &types.ObjectSelector{
			Selector: &types.ObjectSelector_Object{
				Object: c.Registration.Object,
			},
		},
		RelationSelector: &types.RelationSelector{
			Selector: &types.RelationSelector_Relation{
				Relation: policy.OwnerRelation,
			},
		},
		SubjectSelector: &types.SubjectSelector{
			Selector: &types.SubjectSelector_Wildcard{
				Wildcard: &types.WildcardSelector{},
			},
		},
	}

	records, err := engine.FilterRelationships(ctx, c.Policy, selector)
	if err != nil {
		return nil, err
	}
	if len(records) == 0 {
		return nil, nil
	} else if len(records) > 1 {
		// This is a bad and unexpected condition.
		// If this happens in the system there might be a vulnerability or a protocol rule was updated
		err := fmt.Errorf("object %v has more than one owner: %w", c.Registration.Object, types.ErrAcpProtocolViolation)
		return nil, err
	}

	return records[0], nil
}

func (c *RegisterObjectCommand) resolveObjectStatus(record *types.RelationshipRecord) objectRegistrationStatus {
	if record == nil {
		return statusUnregistered
	}
	if record.Archived == true {
		return statusArchived
	}
	return statusActive
}

func (c *RegisterObjectCommand) unregisteredStrategy(ctx context.Context, engine auth_engine.AuthEngine) (types.RegistrationResult, error) {
	err := c.createOwnerRelationship(ctx, engine)
	if err != nil {
		return types.RegistrationResult_NoOp, err
	}

	return types.RegistrationResult_Registered, nil
}

func (c *RegisterObjectCommand) createOwnerRelationship(ctx context.Context, engine auth_engine.AuthEngine) error {
	record := types.RelationshipRecord{
		Relationship: &types.Relationship{
			Object:   c.Registration.Object,
			Relation: policy.OwnerRelation,
			Subject: &types.Subject{
				Subject: &types.Subject_Actor{
					Actor: c.Registration.Actor,
				},
			},
		},
		Creator:      c.Registration.Actor.Id,
		PolicyId:     c.Policy.Id,
		Archived:     false,
		CreationTime: c.CreationTs,
	}
	_, err := engine.SetRelationship(ctx, c.Policy, &record)
	return err
}

func (c *RegisterObjectCommand) activeObjectStrategy(record *types.RelationshipRecord) (types.RegistrationResult, error) {
	if record.Creator != c.Registration.Actor.Id {
		return types.RegistrationResult_NoOp, types.ErrNotAuthorized
	}

	return types.RegistrationResult_NoOp, nil
}

func (c *RegisterObjectCommand) archivedObjectStrategy(ctx context.Context, engine auth_engine.AuthEngine, record *types.RelationshipRecord) (types.RegistrationResult, error) {
	if record.Creator != c.Registration.Actor.Id {
		return types.RegistrationResult_NoOp, types.ErrNotAuthorized
	}

	err := c.unarchiveRelationship(ctx, engine, record)
	if err != nil {
		return types.RegistrationResult_NoOp, err
	}

	return types.RegistrationResult_Unarchived, nil
}

func (c *RegisterObjectCommand) unarchiveRelationship(ctx context.Context, engine auth_engine.AuthEngine, record *types.RelationshipRecord) error {
	record.Archived = false
	_, err := engine.SetRelationship(ctx, c.Policy, record)
	return err
}

type SetRelationshipCommand struct {
	Policy       *types.Policy
	CreationTs   *prototypes.Timestamp
	Creator      string
	Relationship *types.Relationship
}

func (c *SetRelationshipCommand) Execute(ctx context.Context, engine auth_engine.AuthEngine, authorizer *RelationshipAuthorizer) (auth_engine.RecordFound, error) {
	err := c.validate()
	if err != nil {
		return false, fmt.Errorf("failed to set relationship: %w", err)
	}

	creatorActor := types.Actor{
		Id: c.Creator,
	}
	authorized, err := authorizer.IsAuthorized(ctx, c.Policy, c.Relationship, &creatorActor)
	if err != nil {
		return false, fmt.Errorf("failed to set relationship: %w", err)
	}
	if !authorized {
		return false, fmt.Errorf("failed to set relationship: %w", types.ErrNotAuthorized)
	}

	record, err := engine.GetRelationship(ctx, c.Policy, c.Relationship)
	if err != nil {
		return false, fmt.Errorf("failed to set relationship: %w", err)
	}
	if record != nil {
		return true, nil
	}

	record = &types.RelationshipRecord{
		PolicyId:     c.Policy.Id,
		Relationship: c.Relationship,
		CreationTime: c.CreationTs,
		Creator:      c.Creator,
		Archived:     false,
	}
	_, err = engine.SetRelationship(ctx, c.Policy, record)
	if err != nil {
		return false, fmt.Errorf("failed to set relationship: %w", err)
	}

	return false, nil
}

func (c *SetRelationshipCommand) validate() error {
	err := relationshipSpec(c.Policy, c.Relationship)
	if err != nil {
		return err
	}

	if c.Relationship.Relation == policy.OwnerRelation {
		return ErrSetOwnerRel
	}

	return nil
}

// DeleteRelationshipCommand encapsulates the process of removing a relationship from a Policy
type DeleteRelationshipCommand struct {
	// Policy from which Relationship will be removed
	Policy *types.Policy

	// Relationship to be removed
	Relationship *types.Relationship

	// Id of actor that initiated the deletion
	Actor string
}

func (c *DeleteRelationshipCommand) Execute(ctx context.Context, engine auth_engine.AuthEngine, authorizer *RelationshipAuthorizer) (auth_engine.RecordFound, error) {
	err := c.validate()
	if err != nil {
		return false, fmt.Errorf("failed to delete relationship: %w", err)
	}

	isAuthorized, err := c.isActorAuthorized(ctx, authorizer)
	if err != nil {
		return false, fmt.Errorf("failed to delete relationship: %w", err)
	}

	if !isAuthorized {
		return false, fmt.Errorf("failed to delete relationship: %w", types.ErrNotAuthorized)
	}

	found, err := engine.DeleteRelationship(ctx, c.Policy, c.Relationship)
	if err != nil {
		return false, fmt.Errorf("failed to delete relationship: %w", err)
	}

	return found, nil
}

func (c *DeleteRelationshipCommand) validate() error {
	if c.Policy == nil {
		return types.ErrPolicyNil
	}

	if c.Relationship == nil {
		return types.ErrRelationshipNil
	}

	if c.Actor == "" {
		return types.ErrActorNil
	}

	if c.Relationship.Relation == policy.OwnerRelation {
		return ErrDeleteOwnerRel
	}

	return nil
}

// verifies whether actor is authorized to remove the specified Relationship
func (c *DeleteRelationshipCommand) isActorAuthorized(ctx context.Context, authorizer *RelationshipAuthorizer) (bool, error) {
	creatorActor := types.Actor{
		Id: c.Actor,
	}
	return authorizer.IsAuthorized(ctx, c.Policy, c.Relationship, &creatorActor)
}

type UnregisterObjectCommand struct {
	// Target Policy
	Policy *types.Policy

	// Object to be unregistered
	Object *types.Object

	// Actor which initiated request
	Actor string
}

func (c *UnregisterObjectCommand) Execute(ctx context.Context, engine auth_engine.AuthEngine, authorizer *RelationshipAuthorizer) (uint, error) {
	// TODO return found or not
	err := c.validate()
	if err != nil {
		return 0, fmt.Errorf("failed to unregister object: %w", err)
	}

	authRelationship := types.NewActorRelationship(c.Object.Resource, c.Object.Id, policy.OwnerRelation, c.Actor)
	actor := types.Actor{Id: c.Actor}
	authorized, err := authorizer.IsAuthorized(ctx, c.Policy, authRelationship, &actor)
	if err != nil {
		return 0, fmt.Errorf("failed to unregister object: %w", err)
	}
	if !authorized {
		return 0, fmt.Errorf("failed to unregister object: %w", types.ErrNotAuthorized)
	}

	ownerRecord, err := engine.GetRelationship(ctx, c.Policy, authRelationship)
	if err != nil {
		return 0, fmt.Errorf("failed to unregister object: %w", err)
	}
	if ownerRecord.Archived {
		return 0, nil
	}

	count, err := c.removeObjectRelationships(ctx, engine)

	err = c.archiveObject(ctx, engine, ownerRecord)
	if err != nil {
		return 0, fmt.Errorf("failed to unregister object: archiving object: %w", err)
	}

	return count, nil
}

func (c *UnregisterObjectCommand) archiveObject(ctx context.Context, engine auth_engine.AuthEngine, ownerRecord *types.RelationshipRecord) error {
	ownerRecord.Archived = true
	_, err := engine.SetRelationship(ctx, c.Policy, ownerRecord)
	return err

}

func (c *UnregisterObjectCommand) removeObjectRelationships(ctx context.Context, engine auth_engine.AuthEngine) (uint, error) {
	selector := &types.RelationshipSelector{
		ObjectSelector: &types.ObjectSelector{
			Selector: &types.ObjectSelector_Object{
				Object: c.Object,
			},
		},
		RelationSelector: &types.RelationSelector{
			Selector: &types.RelationSelector_Wildcard{
				Wildcard: &types.WildcardSelector{},
			},
		},
		SubjectSelector: &types.SubjectSelector{
			Selector: &types.SubjectSelector_Wildcard{
				Wildcard: &types.WildcardSelector{},
			},
		},
	}
	count, err := engine.DeleteRelationships(ctx, c.Policy, selector)
	if err != nil {
		return 0, fmt.Errorf("failed to unregister object: %w", err)
	}

	return count, nil
}

func (c *UnregisterObjectCommand) validate() error {
	if c.Policy == nil {
		return types.ErrPolicyNil
	}

	if c.Object == nil {
		return types.ErrObjectNil
	}

	if c.Actor == "" {
		return types.ErrActorNil
	}

	return nil
}
