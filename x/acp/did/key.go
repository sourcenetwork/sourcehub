package did

import (
	"fmt"
	"strings"

	"github.com/btcsuite/btcutil/base58"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	"github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/hyperledger/aries-framework-go/pkg/vdr/key"
	"github.com/multiformats/go-varint"
)

var _ Registry = (*KeyRegistry)(nil)

const Base58BTCMultibase byte = 'z'
const didKeyPrefix = "did:key:"
const secp256k1PubMulticodec uint64 = 0xe7
const didKeySECPPrefix string = "did:key:zQ3s"

func NewKeyRegistry() *KeyRegistry {
	return &KeyRegistry{
		vdr: key.New(),
	}
}

// KeyRegistry implements a limit DID registry for the "key" method
type KeyRegistry struct {
	vdr *key.VDR
}

func (r *KeyRegistry) ResolveKey(did string) (types.PubKey, error) {
	found := strings.HasPrefix(did, didKeySECPPrefix)
	if !found {
		return nil, fmt.Errorf("invalid did: did must have method key and key SECP")
	}

	encodedPart, _ := strings.CutPrefix(did, "did:key:z")
	encodedBytes := base58.Decode(encodedPart)
	trimLen := varint.UvarintSize(secp256k1PubMulticodec)
	keyBytes := encodedBytes[trimLen:] // strip the first 8 bytes from the multicodec definition

	return &secp256k1.PubKey{
		Key: keyBytes,
	}, nil
}

// Create generates a DID Key identifier from a PubKey
// Currently only cosmos-sdk secp256k1 keys are supported
func (r *KeyRegistry) Create(key types.PubKey) (string, error) {
	switch k := key.(type) {
	case *secp256k1.PubKey:
		break
	default:
		return "", fmt.Errorf("unsuported key type: %v", k)
	}

	// prefix part contains the did:key: prefix and the Base56BTC Multibase indicator (ie 'z')
	prefix := []byte(didKeyPrefix)
	prefix = append(prefix, Base58BTCMultibase)

	// encodingPart contains the base58 encoded string of secp256k1pub multicodec value and
	// the pub key bytes
	var encodingPart []byte
	encodingPart = varint.ToUvarint(secp256k1PubMulticodec)
	encodingPart = append(encodingPart, key.Bytes()...)

	encodedKey := base58.Encode(encodingPart)

	return string(prefix) + encodedKey, nil
}
