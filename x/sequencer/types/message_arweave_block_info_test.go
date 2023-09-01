package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	"github.com/warp-contracts/sequencer/testutil/sample"
)

func TestMsgArweaveBlockInfo_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgArweaveBlockInfo
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgArweaveBlockInfo{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "invalid arewave block height",
			msg: MsgArweaveBlockInfo{
				Creator:   sample.AccAddress(),
				Height:    0,
				Timestamp: 1690809540,
				Hash:      []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2},
			},
			err: ErrBadArweaveHeight,
		}, {
			name: "invalid arewave block timestamp",
			msg: MsgArweaveBlockInfo{
				Creator:   sample.AccAddress(),
				Height:    1431216,
				Timestamp: 0,
				Hash:      []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2},
			},
			err: ErrBadArweaveTimestamp,
		}, {
			name: "invalid arewave block hash length",
			msg: MsgArweaveBlockInfo{
				Creator:   sample.AccAddress(),
				Height:    1431216,
				Timestamp: 0,
				Hash:      []byte{},
			},
			err: ErrBadArweaveHashLength,
		}, {
			name: "valid",
			msg: MsgArweaveBlockInfo{
				Creator:   sample.AccAddress(),
				Height:    1431216,
				Timestamp: 1690809540,
				Hash:      []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2},
			},
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
