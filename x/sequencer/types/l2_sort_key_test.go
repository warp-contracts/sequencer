package types

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSortKeyNextValues(t *testing.T) {
	sortKey := NewSortKey(100, 200)

	value1 := sortKey.GetNextValue()
	value2 := sortKey.GetNextValue()
	value3 := sortKey.GetNextValue()

	require.Equal(t, "000000000100,0000000000200,00000000", value1)
	require.Equal(t, "000000000100,0000000000200,00000001", value2)
	require.Equal(t, "000000000100,0000000000200,00000002", value3)
}
