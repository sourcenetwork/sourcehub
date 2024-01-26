package embedded

import (
	"context"
	"fmt"
	"os"

	"cosmossdk.io/log"
	"cosmossdk.io/store"
	"cosmossdk.io/store/metrics"
	storetypes "cosmossdk.io/store/types"
	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	dbm "github.com/cosmos/cosmos-db"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	"github.com/sourcenetwork/sourcehub/x/acp/keeper"
	"github.com/sourcenetwork/sourcehub/x/acp/testutil"
	"github.com/sourcenetwork/sourcehub/x/acp/types"
)

const (
	DefaultDataDir string = ".sourcehub-embedded-acp"
	dataFile       string = "data"
)

// LocalACP wraps the acp module Keeper with a local storage.
// It allows clients to experiment with the ACP module without,
// running the consensus engine and a full node directly.
type LocalACP struct {
	baseCtx   context.Context
	keeper    *keeper.Keeper
	msgServer types.MsgServer
	accKeeper types.AccountKeeper
	db        dbm.DB
}

// GetMsgService returns an implementation of acp's MsgServer
func (l *LocalACP) GetMsgService() types.MsgServer {
	return l.msgServer
}

// GetQueryService returns an implementation of acp's QueryServer
func (l *LocalACP) GetQueryService() types.QueryServer {
	return *l.keeper
}

// GetCtx returns the a go Context which wraps a Cosmos ctx.
// This context MUST be used to interact with the Embedded ACP.
func (l *LocalACP) GetCtx() context.Context {
	return l.baseCtx
}

// closes the underlying DB
func (l *LocalACP) Close() error {
	return l.db.Close()
}

// NewLocalACP creates an instance of LocalACP with the given options.
//
// The default ACP configuration persists data under the user home directory and produces no logs.
func NewLocalACP(options ...Option) (LocalACP, error) {
	o := newDefaultOption()
	o.Apply(options...)

	db, err := o.NewDB()
	if err != nil {
		return LocalACP{}, err
	}

	memStoreKey := storetypes.NewMemoryStoreKey(types.MemStoreKey)
	storeKey := storetypes.NewKVStoreKey(types.StoreKey)

	stateStore := store.NewCommitMultiStore(db, o.logger, o.metrics)
	stateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(memStoreKey, storetypes.StoreTypeMemory, nil)

	err = stateStore.LoadLatestVersion()
	if err != nil {
		return LocalACP{}, err
	}

	registry := codectypes.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(registry)

	accKeeper := &testutil.AccountKeeperStub{}

	storeService := runtime.NewKVStoreService(storeKey)
	authority := authtypes.NewModuleAddress(govtypes.ModuleName)

	k := keeper.NewKeeper(
		cdc,
		storeService,
		log.NewNopLogger(),
		authority.String(),
		accKeeper,
	)

	ctx := sdk.NewContext(stateStore, tmproto.Header{}, false, o.logger)

	k.SetParams(ctx, types.DefaultParams())

	return LocalACP{
		baseCtx:   ctx,
		keeper:    &k,
		msgServer: NewMsgServer(k, stateStore),
		accKeeper: accKeeper,
		db:        db,
	}, nil
}

// Option specifies the Local ACP parameters during its construction
type Option func(o *option)

// WithPersistentStorage configures Embedded ACP's persistent store file system location
func WithPersistentStorage(path string) Option {
	return func(o *option) {
		o.storePath = path
	}
}

// WithInMemStore configures Embeded ACP to use a volatile in memory store
func WithInMemStore() Option {
	return func(o *option) {
		o.storePath = ""
	}
}

// WithLogger configures Embedded ACP's Logger
func WithLogger(logger log.Logger) Option {
	return func(o *option) {
		o.logger = logger
	}
}

// WithMetrics configures the Metric collector for the Embedded ACP
func WithMetrics(metrics metrics.StoreMetrics) Option {
	return func(o *option) {
		o.metrics = metrics
	}
}

// newDefaultOption returns the default option struct
// with no log, metrics and data dir stored in the user's home
func newDefaultOption() option {
	return option{
		storePath: getDefaultStoragePath(),
		logger:    log.NewNopLogger(),
		metrics:   metrics.NewNoOpMetrics(),
	}
}

// options for LocalAcp factory
type option struct {
	storePath string
	logger    log.Logger
	metrics   metrics.StoreMetrics
}

// getDefaultStoragePath returns the default data store path for embedded ACP
func getDefaultStoragePath() string {
	home, err := os.UserHomeDir()
	if err != nil {
		// FIXME I'm not sure how this would behave in a WASM context
		// if this panic ever happens in WASM, figure out an alternative dir for wasm.
		panic(fmt.Errorf("home directory must be set: %v", err))
	}

	return home + "/" + DefaultDataDir
}

func (o *option) Apply(opts ...Option) {
	for _, opt := range opts {
		opt(o)
	}
}

// NewDB returns the DB to be used by Local ACP
func (o *option) NewDB() (dbm.DB, error) {
	backend := dbm.GoLevelDBBackend
	if o.storePath == "" {
		backend = dbm.MemDBBackend
	}

	return dbm.NewDB(dataFile, backend, o.storePath)
}
