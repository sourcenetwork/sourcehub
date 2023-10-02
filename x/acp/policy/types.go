package policy

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"

	"github.com/sourcenetwork/sourcehub/x/acp/types"
)

// policyIder builds Policy ids
type policyIder struct{}

// buildId computes the unique id for a policy.
//
// the id is a hash of the policy hash, creator account addr and account sequence number.
func (i *policyIder) Id(pol *types.Policy, sequence uint64) string {
	hasher := sha256.New()

	hasher.Write(i.hashPol(pol))
	hasher.Write([]byte(fmt.Sprintf("%v", sequence)))

	hash := hasher.Sum(nil)
	id := hex.EncodeToString(hash)
	return id
}

// hashPol computes a partial sha256 hash of a policy.
// the hashing algorithm includes a subset of the fields which are deterministic.
func (i *policyIder) hashPol(pol *types.Policy) []byte {
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
