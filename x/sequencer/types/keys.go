package types

const (
	// ModuleName defines the module name
	ModuleName = "sequencer"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_sequencer"
)

var (
	ParamsKey = []byte("p_sequencer")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	LastArweaveBlockKey = "LastArweaveBlock/value/"
)
