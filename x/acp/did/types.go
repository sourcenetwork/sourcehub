package did

import (
	"github.com/cosmos/cosmos-sdk/crypto/types"
)

// Registry model a DID method registry
type Registry interface {
	// ResolveKey resolves the did and returns the first key from it
	ResolveKey(did string) (types.PubKey, error)

	// Create prodcues a DID from a pub key
	Create(key types.PubKey) (string, error)
}
