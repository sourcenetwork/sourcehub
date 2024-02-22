package address

import (
	"fmt"

	"cosmossdk.io/core/address"
	sdkaddr "github.com/cosmos/cosmos-sdk/codec/address"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sourcenetwork/sourcehub/did"
)

var _ address.Codec = &AddressCodec{}

func NewDIDDecoder(resolver did.Resolver) DIDDecoder {
	return DIDDecoder{
		resolver: resolver,
	}
}

// DIDDecoder decodes a DID into a Cosmos address
type DIDDecoder struct {
	resolver did.Resolver
}

// StringToBytes resolves a did object into a public key which gets converted to an SDK address
func (c *DIDDecoder) StringToBytes(did string) ([]byte, error) {
	key, err := c.resolver.ResolveKey(did)
	if err != nil {
		return nil, fmt.Errorf("failed decoding did address: %v", err)
	}

	accAddr := sdk.AccAddress(key.Address().Bytes())
	err = sdk.VerifyAddressFormat(accAddr)
	if err != nil {
		return nil, fmt.Errorf("could not generate address from did: %v", err)
	}

	return accAddr, nil
}

func NewAddressCodec(didDec DIDDecoder, bech32Codec sdkaddr.Bech32Codec) *AddressCodec {
	return &AddressCodec{
		didDec:      didDec,
		bech32Codec: bech32Codec,
	}
}

// AddressCodec acts a custom codec for SourceHub which can decode both Bech32 and DIDs
// into SDK Addresses
// Encoding is done using Bech32
type AddressCodec struct {
	didDec      DIDDecoder
	bech32Codec address.Codec
}

// StringToBytes encodes text to bytes
func (c *AddressCodec) StringToBytes(text string) ([]byte, error) {
	if did.IsDID(text) {
		return c.didDec.StringToBytes(text)
	}
	return c.bech32Codec.StringToBytes(text)
}

// BytesToString encodes bytes to text
func (c *AddressCodec) BytesToString(bz []byte) (string, error) {
	return c.bech32Codec.BytesToString(bz)
}
