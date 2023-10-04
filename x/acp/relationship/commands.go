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

func (c *RegisterObjectCommand) Execute(ctx context.Context, engine auth_engine.AuthEngine) (types.RegistrationResult, error) {
	var err error
	var result types.RegistrationResult

	err = c.validate()
	if err != nil {
		return result, fmt.Errorf("%w: %v", ErrRegisterObject, err)
	}

	record, err := c.getOwnerRelationship(ctx, engine)
	if err != nil {
		return result, fmt.Errorf("%w: %v", ErrRegisterObject, err)
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
		return result, fmt.Errorf("%w: %w", ErrRegisterObject, err)
	}

	return result, nil
}

// validates the command input params
func (c *RegisterObjectCommand) validate() error {
	if c.Policy == nil {
		return fmt.Errorf("Policy is nil")
	}

	if err := c.Registration.Validate(); err != nil {
		return err
	}

	if c.CreationTs == nil {
		return fmt.Errorf("CreationTs is nil")
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
		return nil, fmt.Errorf("invariant violation: object %v has more than one owner", c.Registration.Object)
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
		return types.RegistrationResult_Noop, err
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
		return types.RegistrationResult_Denied, nil
	}

	return types.RegistrationResult_Noop, nil
}

func (c *RegisterObjectCommand) archivedObjectStrategy(ctx context.Context, engine auth_engine.AuthEngine, record *types.RelationshipRecord) (types.RegistrationResult, error) {
	if record.Creator != c.Registration.Actor.Id {
		return types.RegistrationResult_Denied, nil
	}

	err := c.unarchiveRelationship(ctx, engine, record)
	if err != nil {
		return types.RegistrationResult_Noop, err
	}

	return types.RegistrationResult_Unarchived, nil
}

func (c *RegisterObjectCommand) unarchiveRelationship(ctx context.Context, engine auth_engine.AuthEngine, record *types.RelationshipRecord) error {
	record.Archived = false
	_, err := engine.SetRelationship(ctx, c.Policy, record)
	return err
}

type SetRelationshipCommand struct {
    Policy *types.Policy
    CreationTs   *prototypes.Timestamp
    Creator string
    Relationship *types.Relationship
}

func (c *SetRelationshipCommand) Execute(ctx context.Context, engine auth_engine.AuthEngine) (types.SetRelationshipResult, error) {
    err := c.validate()
    if err != nil {
        return types.SetRelationshipResult_SetRelNoOp, err
    }

    authorizer := NewRelationshipAuthorizer(engine)
    creatorActor := types.Actor{
        Id: c.Creator,
    }
    authorized, err := authorizer.IsAuthorized(ctx, c.Policy, c.Relationship, &creatorActor)
    if err != nil {
        return types.SetRelationshipResult_SetRelNoOp, err
    }
    if !authorized {
        return types.SetRelationshipResult_SetRelDenied, nil
    }

    record, err := engine.GetRelationship(ctx, c.Policy, c.Relationship)
    if err != nil {
        return types.SetRelationshipResult_SetRelNoOp, err
    }
    if record != nil {
        return types.SetRelationshipResult_SetRelNoOp, nil
    }

    record = &types.RelationshipRecord{
        PolicyId: c.Policy.Id,
        Relationship: c.Relationship,
        CreationTime: c.CreationTs,
        Creator: c.Creator,
        Archived: false,
    }
    _, err = engine.SetRelationship(ctx, c.Policy, record)
    if err != nil {
        return types.SetRelationshipResult_SetRelNoOp, err
    }

    return types.SetRelationshipResult_SetRelCreated, nil
}

func (c *SetRelationshipCommand) validate() error {
    if c.Relationship.Relation == policy.OwnerRelation {
        return ErrCannotSetOwnerRelationship
    }

    return nil
}
