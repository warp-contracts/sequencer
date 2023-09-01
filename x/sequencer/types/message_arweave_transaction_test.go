package types

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMsgArweaveTransaction_ValidateBasic(t *testing.T) {
	msg := MsgArweaveTransaction{}

	err := msg.ValidateBasic()

	require.NotNil(t, err)
}
