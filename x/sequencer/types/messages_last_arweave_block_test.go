package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	"github.com/warp-contracts/sequencer/testutil/sample"
)

func TestMsgCreateLastArweaveBlock_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCreateLastArweaveBlock
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCreateLastArweaveBlock{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid",
			msg: MsgCreateLastArweaveBlock{
				Creator:   sample.AccAddress(),
				Timestamp: 1690809540,
				Height:    1431216,
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

func TestMsgUpdateLastArweaveBlock_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgUpdateLastArweaveBlock
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgUpdateLastArweaveBlock{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgUpdateLastArweaveBlock{
				Creator: sample.AccAddress(),
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

func TestMsgDeleteLastArweaveBlock_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgDeleteLastArweaveBlock
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgDeleteLastArweaveBlock{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid",
			msg: MsgDeleteLastArweaveBlock{
				Creator: sample.AccAddress(),
				Hash:    []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2},
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
