package relationship

import (
	"fmt"

	"github.com/sourcenetwork/sourcehub/x/acp/did"
	"github.com/sourcenetwork/sourcehub/x/acp/types"
)

// relationshipSpec validates a Relationshp according to the expected format.
// For now, we can rely on Zanzi to perform the bulk of the validation,
// however its paramount that a Relationship whose subject type is an Actor must be a DID
func relationshipSpec(policy *types.Policy, relationship *types.Relationship) error {
	switch subj := relationship.Subject.Subject.(type) {
	case *types.Subject_Actor:
		if err := did.IsValidDID(subj.Actor.Id); err != nil {
			return fmt.Errorf("%w: actor must be a valid did: %v", ErrInvalidRelationship, err)
		}
	case *types.Subject_Object:
		err := did.IsValidDID(subj.Object.Id)
		if subj.Object.Resource == policy.ActorResource.Name && err != nil {
			return fmt.Errorf("%w: actor must be a valid did: %v", ErrInvalidRelationship, err)
		}
	}

	if relationship.Object.Id == "" {
		return fmt.Errorf("object id must not be empty: %w", ErrInvalidRelationship)
	}

	return nil
}

func registrationSpec(registration *types.Registration) error {
	if registration == nil {
		return types.ErrRegistrationNil
	}

	if registration.Actor == nil {
		return fmt.Errorf("invalid registration: %w", types.ErrActorNil)
	}

	if registration.Object == nil {
		return fmt.Errorf("invalid registration: %w", types.ErrObjectNil)
	}

	if registration.Object.Id == "" {
		return fmt.Errorf("invalid registration: object id required")
	}

	if err := did.IsValidDID(registration.Actor.Id); err != nil {
		return fmt.Errorf("invalid registration: %v", err)
	}

	return nil
}
