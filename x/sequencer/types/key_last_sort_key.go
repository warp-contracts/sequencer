package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// LastSortKeyKeyPrefix is the prefix to retrieve all LastSortKey
	LastSortKeyKeyPrefix = "LastSortKey/value/"
)

// LastSortKeyKey returns the store key to retrieve a LastSortKey from the index fields
func LastSortKeyKey(
	contract string,
) []byte {
	var key []byte

	contractBytes := []byte(contract)
	key = append(key, contractBytes...)
	key = append(key, []byte("/")...)

	return key
}
