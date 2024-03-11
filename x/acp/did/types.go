package did

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	"github.com/cyware/ssi-sdk/crypto"
	"github.com/cyware/ssi-sdk/did"
	"github.com/cyware/ssi-sdk/did/key"
	didmodel "github.com/hyperledger/aries-framework-go/component/models/did"
)

type DIDDocument struct {
	*did.Document
}

type Resolver interface {
	Resolve(ctx context.Context, did string) (DIDDocument, error)
}

func IsValidDID(did string) error {
	_, err := didmodel.Parse(did)
	if err != nil {
		return fmt.Errorf("did %v: %v", did, err)
	}
	return nil
}

// IssueDID produces a DID for a SourceHub account
func IssueDID(acc sdk.AccountI) (string, error) {
	var keyType crypto.KeyType
	switch acc.GetPubKey().(type) {
	case *secp256k1.PubKey:
		keyType = crypto.SECP256k1
	case *ed25519.PubKey:
		keyType = crypto.Ed25519
	default:
		return "", fmt.Errorf("failed to issue did for %v: account key type must be secp256k1 or ed25519", acc.GetAddress().String())
	}

	did, err := key.CreateDIDKey(keyType, acc.GetPubKey().Bytes())
	if err != nil {
		return "", fmt.Errorf("failed to generated did for %v: %v", acc.GetAddress().String(), err)
	}

	return did.String(), nil
}
