package policy

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sourcenetwork/sourcehub/x/acp/types"
)

// basicPolicyIRSpec performs basic initial validation of a PolicyIR.
// Returns nil if the initial validation is accepted
func basicPolicyIRSpec(pol *PolicyIR) error {
	for _, resource := range pol.Resources {
		found := false
		for _, relation := range resource.Relations {
			if relation.Name == OwnerRelation {
				found = true
			}
		}
		if !found {
			return fmt.Errorf("resource %v: %w", resource.Name, ErrResourceMissingOwnerRelation)
		}
	}

	return nil
}

type validPolicySpec struct{}

// Validates a local Policy type, returns nil if Policy is ok
//
// The ACP module can delegate most of the Policy validation to Zanzi itself.
// The exception is the Manage Graph which is local to the acp system.
func (v *validPolicySpec) Satisfies(pol *types.Policy) error {
	g := buildManagementGraph(pol)
	err := g.IsWellFormed()
	if err != nil {
		return fmt.Errorf("%w: %w", ErrMalformedGraph, err)
	}

	_, err = sdk.AccAddressFromBech32(pol.Creator)
	if err != nil {
		return fmt.Errorf("%w: %v", ErrInvalidCreator, err)
	}

	err = v.resourcesContainOwner(pol)
	if err != nil {
		return err
	}

	return nil
}

func (v *validPolicySpec) resourcesContainOwner(pol *types.Policy) error {
	for _, resource := range pol.Resources {
		found := false
		for _, relation := range resource.Relations {
			if relation.Name == OwnerRelation {
				found = true
			}
		}
		if !found {
			return fmt.Errorf("resource %v: %w", resource.Name, ErrResourceMissingOwnerRelation)
		}
	}

	return nil
}
