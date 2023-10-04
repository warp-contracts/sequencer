package ante

import (
	"github.com/stretchr/testify/require"
	"testing"

	"github.com/warp-contracts/sequencer/x/sequencer/test"
)

func TestBlockInteractions(t *testing.T) {
	bi := NewBlockInteractions()
	dataItem := test.ArweaveL2Interaction(t)
	height := int64(13)

	bi.NewBlock(height)

	result := bi.Contains(height, &dataItem)
	require.False(t, result)

	bi.Add(height, &dataItem)
	result = bi.Contains(height, &dataItem)
	require.True(t, result)
}

func TestBlockInteractionsResetInvalidHeight(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("The function should panic in case of invalid height")
		}
	}()

	bi := NewBlockInteractions()
	bi.NewBlock(5)
	bi.NewBlock(7)
}

func TestBlockInteractionsAddInvalidHeight(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("The function should panic in case of invalid height")
		}
	}()

	bi := NewBlockInteractions()
	dataItem := test.ArweaveL2Interaction(t)

	bi.NewBlock(5)
	bi.Add(6, &dataItem)
}

func TestBlockInteractionsContainsInvalidHeight(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("The function should panic in case of invalid height")
		}
	}()

	bi := NewBlockInteractions()
	dataItem := test.ArweaveL2Interaction(t)

	bi.NewBlock(5)
	bi.Contains(4, &dataItem)
}

func TestBlockInteractionsNotInitialized(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("The function should panic In the case of no initialization")
		}
	}()

	bi := NewBlockInteractions()
	dataItem := test.ArweaveL2Interaction(t)

	bi.Add(3, &dataItem)
}
