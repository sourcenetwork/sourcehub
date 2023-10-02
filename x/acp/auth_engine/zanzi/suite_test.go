package zanzi

import (
	"context"
	"testing"

	tmdb "github.com/cometbft/cometbft-db"
	"github.com/cometbft/cometbft/libs/log"
	"github.com/cosmos/cosmos-sdk/store"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/sourcenetwork/sourcehub/x/acp/auth_engine"
	"github.com/sourcenetwork/sourcehub/x/acp/types"
)

func setup(t *testing.T) (context.Context, auth_engine.AuthEngine) {
	storeKey := sdk.NewKVStoreKey(types.StoreKey)
	db := tmdb.NewMemDB()
	stateStore := store.NewCommitMultiStore(db)
	stateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, db)
	require.NoError(t, stateStore.LoadLatestVersion())
	kv := stateStore.GetCommitKVStore(storeKey)
	engine, err := NewZanzi(kv, log.NewNopLogger())
	require.Nil(t, err)

	return context.Background(), engine
}

func TestZanzi(t *testing.T) {
	auth_engine.TestAuthEngineImpl(t, setup)
}
