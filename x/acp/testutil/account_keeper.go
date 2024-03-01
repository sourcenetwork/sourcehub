package testutil

import (
	"context"

	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	"github.com/sourcenetwork/sourcehub/x/acp/types"
)

var _ types.AccountKeeper = (*AccountKeeperStub)(nil)

type AccountKeeperStub struct {
	Accounts map[string]sdk.AccountI
}

func (s *AccountKeeperStub) GetAccount(ctx context.Context, address sdk.AccAddress) sdk.AccountI {
	acc := s.Accounts[address.String()]
	return acc
}

func (s *AccountKeeperStub) GenAccount() sdk.AccountI {
	if s.Accounts == nil {
		s.Accounts = make(map[string]sdk.AccountI)
	}

	pubKey := secp256k1.GenPrivKey().PubKey()
	addr := sdk.AccAddress(pubKey.Address())
	acc := authtypes.NewBaseAccount(addr, pubKey, 1, 1)
	s.Accounts[addr.String()] = acc
	return acc
}

func (s *AccountKeeperStub) FirstAcc() sdk.AccountI {
	for _, acc := range s.Accounts {
		return acc
	}
	return nil
}
