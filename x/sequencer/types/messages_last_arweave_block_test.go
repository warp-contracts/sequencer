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
			name: "valid address",
			msg: MsgCreateLastArweaveBlock{
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
			name: "valid address",
			msg: MsgDeleteLastArweaveBlock{
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