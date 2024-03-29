package types

const (
	// ModuleName defines the module name
	ModuleName = "acp"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_acp"

	// AccessDecisionRepositoryKey defines the namespace for Access Decisions
	AccessDecisionRepositoryKey = "access_decision"
)

var (
	ParamsKey = []byte("p_acp")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
