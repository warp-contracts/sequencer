package prepare

import "fmt"

// A structure useful for assigning SortKey values to subsequent interactions in a block.
type SortKey struct {
	arweaveHeight uint64
	sequencerHeight int64
	index int
}

func newSortKey(arweaveHeight uint64, sequencerHeight int64) *SortKey {
	return &SortKey{
		arweaveHeight,
		sequencerHeight,
		0,
	}
}

func (key *SortKey) getNextValue() string {
	value := fmt.Sprintf("%.12d,%.13d,%.8d", key.arweaveHeight, key.sequencerHeight, key.index)
	key.index++
	return value
}
