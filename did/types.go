package did

import (
	"strings"

	"github.com/cosmos/cosmos-sdk/crypto/types"
)

const didPrefix = "did"

// Registry model a DID method registry
type Registry interface {
	// ResolveKey resolves the did and returns the first key from it
	ResolveKey(did string) (types.PubKey, error)

	// Create prodcues a DID from a pub key
	Create(key types.PubKey) (string, error)
}

type Resolver interface {
	ResolveKey(did string) (types.PubKey, error)
}

// IsDID checks whether `did` is a valid did string
func IsDID(did string) bool {
	return strings.HasPrefix(did, didPrefix)
}
