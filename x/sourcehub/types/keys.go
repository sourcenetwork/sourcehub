package types

const (
	// ModuleName defines the module name
	ModuleName = "sourcehub"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_sourcehub"
)

var (
	ParamsKey = []byte("p_sourcehub")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
