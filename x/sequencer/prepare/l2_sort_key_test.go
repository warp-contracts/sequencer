package prepare


import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSortKey(t *testing.T) {
	sortKey := newSortKey(100, 200)

	value1 := sortKey.getNextValue()
	value2 := sortKey.getNextValue()
	value3 := sortKey.getNextValue()

	require.Equal(t, "000000000100,0000000000200,00000000", value1)
	require.Equal(t, "000000000100,0000000000200,00000001", value2)
	require.Equal(t, "000000000100,0000000000200,00000002", value3)
}