package pkg

import (
	"encoding/json"
	"log"

	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/bech32"

	"github.com/sourcenetwork/sourcehub/app"
)

type Session struct {
	// SourceHub address of Authenticated user
	Actor   string
	PrivKey cryptotypes.PrivKey
}

func NewSession(key string) Session {
	priv := &secp256k1.PrivKey{}
	json.Unmarshal([]byte(key), priv)

	addr, err := bech32.ConvertAndEncode(app.AccountAddressPrefix, sdk.AccAddress(priv.PubKey().Address()))
	if err != nil {
		panic(err)
	}

	log.Printf("authenticated as user: %v", addr)
	return Session{
		PrivKey: priv,
		Actor:   addr,
	}
}

// Generate a Key Pair, generate a cosmos addr, submit a tx to the Faucet
func NewActor() (string, *secp256k1.PrivKey) {
	key := secp256k1.GenPrivKey()
	addr := key.PubKey().Address()
	accAddr := sdk.AccAddress(addr)
	accAddrEnc, err := bech32.ConvertAndEncode(app.AccountAddressPrefix, accAddr)
	if err != nil {
		panic(err)
	}

	// TODO register in faucet
	return accAddrEnc, key
}

func DumpKey(key *secp256k1.PrivKey) string {
	res, err := json.Marshal(key)
	if err != nil {
		panic(err)
	}
	return string(res)
}
