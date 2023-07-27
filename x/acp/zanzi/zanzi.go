package zanzi

import (
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	rcdb "github.com/sourcenetwork/raccoondb"
	zanzi "github.com/sourcenetwork/zanzi"
)

func New(kv storetypes.KVStore) zanzi.Zanzi {
	store := rcdb.KvFromCosmosKv(kv)
	return zanzi.New(store)
}
