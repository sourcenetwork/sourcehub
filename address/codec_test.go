package address

import (
	"testing"

	sdkaddr "github.com/cosmos/cosmos-sdk/codec/address"
	"github.com/stretchr/testify/require"

	"github.com/sourcenetwork/sourcehub/did"
)

func TestAddressCodec_DIDStringGetsMappedToAccount(t *testing.T) {
	resolver := did.NewKeyRegistry()
	didDec := NewDIDDecoder(resolver)

	bech := sdkaddr.Bech32Codec{
		Bech32Prefix: "source",
	}
	cdc := NewAddressCodec(didDec, bech)

	addr, err := cdc.StringToBytes("did:key:zQ3shw2LqCxawpfYWypqSyircWhn56ZCFrzF7uNNZAnZFsA1g")

	require.Nil(t, err)
	require.Nil(t, addr)
}
