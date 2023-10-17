package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// PrevSortKeyKeyPrefix is the prefix to retrieve all PrevSortKey
	PrevSortKeyKeyPrefix = "PrevSortKey/value/"
)

// PrevSortKeyKey returns the store key to retrieve a PrevSortKey from the index fields
func PrevSortKeyKey(
	contract string,
) []byte {
	var key []byte

	contractBytes := []byte(contract)
	key = append(key, contractBytes...)
	key = append(key, []byte("/")...)

	return key
}
