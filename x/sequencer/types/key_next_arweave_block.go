package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// NextArweaveBlockKeyPrefix is the prefix to retrieve all NextArweaveBlock
	NextArweaveBlockKeyPrefix = "NextArweaveBlock/value/"
)

// NextArweaveBlockKey returns the store key to retrieve a NextArweaveBlock from the index fields
func NextArweaveBlockKey(
	height string,
) []byte {
	var key []byte

	heightBytes := []byte(height)
	key = append(key, heightBytes...)
	key = append(key, []byte("/")...)

	return key
}
