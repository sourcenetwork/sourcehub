package types

const (
	// ModuleName defines the module name
	ModuleName = "acp"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_acp"
)

var (
	ParamsKey = []byte("p_acp")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
