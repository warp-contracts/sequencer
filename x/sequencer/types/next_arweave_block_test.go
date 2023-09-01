package types

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNextArweaveBlock_GetHeightString(t *testing.T) {
	const blockHeight uint64 = 123
	block := NextArweaveBlock {
		BlockInfo: &ArweaveBlockInfo{
			Height: blockHeight,
		},
	}

	heightStr := block.GetHeightString()
	parsedHeight, err := strconv.ParseUint(heightStr, 10, 64)

	require.NoError(t, err)
	require.Equal(t, blockHeight, parsedHeight)
}
