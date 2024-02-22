package did

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	"github.com/stretchr/testify/require"
)

func TestKeysecp256k1ToDID(t *testing.T) {
	priv := secp256k1.GenPrivKey()
	pub := priv.PubKey()

	registry := NewKeyRegistry()
	did, err := registry.Create(pub)
	t.Log(did)

	require.Nil(t, err)

	recovered, err := registry.ResolveKey(did)
	require.Nil(t, err)
	require.Equal(t, recovered.Bytes(), pub.Bytes())
}
