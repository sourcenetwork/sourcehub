package did

import (
	stdcrypto "crypto"
	"crypto/ed25519"

	"github.com/cyware/ssi-sdk/crypto"
	"github.com/cyware/ssi-sdk/did/key"
)

func ProduceDID() (string, stdcrypto.Signer, error) {
	pkey, skey, err := ed25519.GenerateKey(nil)
	if err != nil {
		return "", nil, err
	}

	keyType := crypto.Ed25519
	did, err := key.CreateDIDKey(keyType, []byte(pkey))
	if err != nil {
		return "", nil, err
	}

	return did.String(), skey, nil
}
