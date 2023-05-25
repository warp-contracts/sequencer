package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	"github.com/warp-contracts/sequencer/testutil/sample"
	"github.com/warp-contracts/syncer/src/utils/bundlr"
)

func TestMsgDataItem_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgDataItem
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgDataItem{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "empty data item",
			msg: MsgDataItem{
				Creator: sample.AccAddress(),
			},
			err: bundlr.ErrVerifyIdSignatureMismatch,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
