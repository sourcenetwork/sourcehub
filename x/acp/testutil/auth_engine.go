package testutil

import (
	"testing"

	"cosmossdk.io/log"
	"cosmossdk.io/store"
	"cosmossdk.io/store/metrics"
	storetypes "cosmossdk.io/store/types"
	cmtproto "github.com/cometbft/cometbft/proto/tendermint/types"
	dbm "github.com/cosmos/cosmos-db"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/sourcenetwork/sourcehub/x/acp/auth_engine"
	"github.com/sourcenetwork/sourcehub/x/acp/auth_engine/zanzi"
	"github.com/sourcenetwork/sourcehub/x/acp/types"
)

// GetTestAuthEngine returns an Auth Engine bound to a test
// using the Zanzi Engine implementation
func GetTestAuthEngine(t *testing.T) (auth_engine.AuthEngine, storetypes.CommitMultiStore, sdk.Context) {
	storeKey := storetypes.NewKVStoreKey(types.StoreKey)
	db := dbm.NewMemDB()
	stateStore := store.NewCommitMultiStore(db, log.NewNopLogger(), metrics.NewNoOpMetrics())
	stateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, db)
	require.NoError(t, stateStore.LoadLatestVersion())

	kv := stateStore.GetCommitKVStore(storeKey)
	engine, err := zanzi.NewZanzi(kv, log.NewNopLogger())
	require.Nil(t, err)

	ctx := sdk.NewContext(stateStore, cmtproto.Header{}, false, log.NewNopLogger())

	return engine, stateStore, ctx
}
