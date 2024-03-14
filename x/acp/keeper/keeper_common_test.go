package keeper

import (
	"crypto"
	"testing"
	"time"

	"cosmossdk.io/log"
	"cosmossdk.io/store"
	"cosmossdk.io/store/metrics"
	storetypes "cosmossdk.io/store/types"
	cmtproto "github.com/cometbft/cometbft/proto/tendermint/types"
	dbm "github.com/cosmos/cosmos-db"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	prototypes "github.com/cosmos/gogoproto/types"
	"github.com/stretchr/testify/require"

	"github.com/sourcenetwork/sourcehub/x/acp/did"
	"github.com/sourcenetwork/sourcehub/x/acp/testutil"
	"github.com/sourcenetwork/sourcehub/x/acp/types"
)

var timestamp, _ = prototypes.TimestampProto(time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC))

func setupMsgServer(t *testing.T) (sdk.Context, types.MsgServer, *testutil.AccountKeeperStub) {
	ctx, keeper, accK := setupKeeper(t)
	return ctx, NewMsgServerImpl(keeper), accK
}

func setupKeeper(t *testing.T) (sdk.Context, Keeper, *testutil.AccountKeeperStub) {
	storeKey := storetypes.NewKVStoreKey(types.StoreKey)

	db := dbm.NewMemDB()
	stateStore := store.NewCommitMultiStore(db, log.NewNopLogger(), metrics.NewNoOpMetrics())
	stateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, db)
	require.NoError(t, stateStore.LoadLatestVersion())

	registry := codectypes.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(registry)
	authority := authtypes.NewModuleAddress(govtypes.ModuleName)

	accKeeper := &testutil.AccountKeeperStub{}
	accKeeper.GenAccount()

	keeper := NewKeeper(
		cdc,
		runtime.NewKVStoreService(storeKey),
		log.NewNopLogger(),
		authority.String(),
		accKeeper,
	)

	ctx := sdk.NewContext(stateStore, cmtproto.Header{}, false, log.NewNopLogger())
	ctx = ctx.WithEventManager(sdk.NewEventManager())

	// Initialize params
	keeper.SetParams(ctx, types.DefaultParams())

	return ctx, keeper, accKeeper
}

func mustGenerateActor() (string, crypto.Signer) {
	bob, bobSigner, err := did.ProduceDID()
	if err != nil {
		panic(err)
	}
	return bob, bobSigner
}
