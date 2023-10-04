package testutil

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	"github.com/sourcenetwork/sourcehub/x/acp/types"
)

var _ types.AccountKeeper = (*AccountKeeperStub)(nil)

type AccountKeeperStub struct{}

func (s *AccountKeeperStub) GetAccount(ctx sdk.Context, address sdk.AccAddress) authtypes.AccountI {
	return &authtypes.BaseAccount{
		Address:  string(address),
		Sequence: 1,
	}
}
