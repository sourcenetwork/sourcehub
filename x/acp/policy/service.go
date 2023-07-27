package policy

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"

	gogotypes "github.com/cosmos/gogoproto/types"
	"github.com/sourcenetwork/sourcehub/x/acp/types"
)

const DefaultActorResourceName string = "actor"

// NewPolicy creates a new policy from a marshal policy string.
// The policy is unmarshaled according to the given marshaling type and normalized.
//
// creator, height and sequence are stored as part of the policy or used to build its id.
func NewPolicy(polStr string, t types.PolicyMarshalingType, creator string, sequence uint64, creationTime *gogotypes.Timestamp) (*types.Policy, error) {

	policy, err := unmarshal(polStr, t)
	if err != nil {
		return nil, err
	}

        normalize(policy)

	id := buildId(policy, sequence)
	policy.Id = id
	policy.Creator = creator
	policy.CreationTime = creationTime

	validator := validator{}
	err = validator.Validate(policy)
	if err != nil {
		return nil, err
	}

	return policy, nil
}

// buildId computes the unique id for a policy.
//
// the id is a hash of the policy hash, creator account addr and account sequence number.
func buildId(pol *types.Policy, sequence uint64) string {
	hasher := sha256.New()

	hasher.Write(hashPol(pol))
	hasher.Write([]byte(fmt.Sprintf("%v", sequence)))

	hash := hasher.Sum(nil)
	id := hex.EncodeToString(hash)
	return id
}

// hashPol computes a partial sha256 hash of a policy.
// the hashing algorithm includes a subset of the fields which are deterministic.
func hashPol(pol *types.Policy) []byte {
	hasher := sha256.New()
	hasher.Write([]byte(pol.Name))
	hasher.Write([]byte(pol.Creator))

	for _, resource := range pol.Resources {
		hasher.Write([]byte(resource.Name))

		for _, rel := range resource.Relations {
			hasher.Write([]byte(rel.Name))
		}

		for _, perm := range resource.Permissions {
			hasher.Write([]byte(perm.Name))
			hasher.Write([]byte(perm.Expression))
		}
	}

	return hasher.Sum(nil)
}

func buildManagementGraph(policy *types.Policy) *types.ManagementGraph {
	graph := &types.ManagementGraph{}
	graph.LoadFromPolicy(policy)
	return graph
}

// normalize normalizes a policy by setting default values for optional fields.
func normalize(pol *types.Policy) {
    if pol.ActorResource == nil {
        pol.ActorResource = &types.ActorResource{
            Name: DefaultActorResourceName,
        }
    }
    
    // policy is sorted before building id to ensure determinism
    pol.Sort()
}
