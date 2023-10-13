package ante

import (
	"github.com/stretchr/testify/require"
	"testing"

	"github.com/warp-contracts/sequencer/x/sequencer/test"
)

func TestBlockInteractions(t *testing.T) {
	bi := NewBlockInteractions()
	dataItem := test.ArweaveL2Interaction(t)

	bi.NewBlock()

	result := bi.Contains(&dataItem)
	require.False(t, result)

	bi.Add(&dataItem)
	result = bi.Contains(&dataItem)
	require.True(t, result)
}
