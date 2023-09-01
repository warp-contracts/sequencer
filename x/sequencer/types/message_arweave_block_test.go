package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	"github.com/warp-contracts/sequencer/testutil/sample"
	"github.com/warp-contracts/sequencer/x/sequencer/types"
)

func TestMsgArweaveBlock_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgArweaveBlock
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgArweaveBlock{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "invalid arewave block height",
			msg: MsgArweaveBlock{
				Creator: sample.AccAddress(),
				BlockInfo: &types.ArweaveBlockInfo{
					Height:    0,
					Timestamp: 1690809540,
					Hash:      []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2},
				},
			},
			err: ErrBadArweaveHeight,
		}, {
			name: "invalid arewave block timestamp",
			msg: MsgArweaveBlock{
				Creator: sample.AccAddress(),
				BlockInfo: &types.ArweaveBlockInfo{
					Height:    1431216,
					Timestamp: 0,
					Hash:      []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2},
				},
			},
			err: ErrBadArweaveTimestamp,
		}, {
			name: "invalid arewave block hash length",
			msg: MsgArweaveBlock{
				Creator: sample.AccAddress(),
				BlockInfo: &types.ArweaveBlockInfo{
					Height:    1431216,
					Timestamp: 0,
					Hash:      []byte{},
				},
			},
			err: ErrBadArweaveHashLength,
		}, {
			name: "valid",
			msg: MsgArweaveBlock{
				Creator: sample.AccAddress(),
				BlockInfo: &types.ArweaveBlockInfo{
					Height:    1431216,
					Timestamp: 1690809540,
					Hash:      []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2},
				},
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
