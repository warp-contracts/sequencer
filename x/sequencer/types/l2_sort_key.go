package types

import "fmt"

// A structure useful for assigning SortKey values to subsequent interactions in a block.
type SortKey struct {
	ArweaveHeight   uint64
	SequencerHeight int64
	Index           int64
}

func NewSortKey(arweaveHeight uint64, sequencerHeight int64) *SortKey {
	return &SortKey{
		arweaveHeight,
		sequencerHeight,
		0,
	}
}

func (key *SortKey) GetNextValue() string {
	value := fmt.Sprintf("%.12d,%.13d,%.8d", key.ArweaveHeight, key.SequencerHeight, key.Index)
	key.Index++
	return value
}
