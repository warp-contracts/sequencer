package types

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMsgArweaveBlock_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgArweaveBlock
		err  error
	}{
		{
			name: "invalid arewave block height",
			msg: MsgArweaveBlock{
				BlockInfo: &ArweaveBlockInfo{
					Height:    0,
					Timestamp: 1690809540,
					Hash:      "C5_l8fu5rftA2lQSDgNELuVX7DDVRofDRJ8v3_OaFXE8Ne4pU5loT-Ljd7JiFL4e",
				},
			},
			err: ErrBadArweaveHeight,
		}, {
			name: "invalid arewave block timestamp",
			msg: MsgArweaveBlock{
				BlockInfo: &ArweaveBlockInfo{
					Height:    1431216,
					Timestamp: 0,
					Hash:      "C5_l8fu5rftA2lQSDgNELuVX7DDVRofDRJ8v3_OaFXE8Ne4pU5loT-Ljd7JiFL4e",
				},
			},
			err: ErrBadArweaveTimestamp,
		}, {
			name: "invalid arewave block hash length",
			msg: MsgArweaveBlock{
				BlockInfo: &ArweaveBlockInfo{
					Height:    1431216,
					Timestamp: 0,
					Hash:      "",
				},
			},
			err: ErrBadArweaveHashLength,
		}, {
			name: "valid",
			msg: MsgArweaveBlock{
				BlockInfo: &ArweaveBlockInfo{
					Height:    1431216,
					Timestamp: 1690809540,
					Hash:      "C5_l8fu5rftA2lQSDgNELuVX7DDVRofDRJ8v3_OaFXE8Ne4pU5loT-Ljd7JiFL4e",
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
