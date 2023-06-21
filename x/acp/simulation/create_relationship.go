package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/sourcenetwork/sourcehub/x/acp/keeper"
	"github.com/sourcenetwork/sourcehub/x/acp/types"
)

func SimulateMsgCreateRelationship(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgCreateRelationship{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the CreateRelationship simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "CreateRelationship simulation not implemented"), nil, nil
	}
}
