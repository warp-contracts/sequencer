package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
	"github.com/warp-contracts/sequencer/x/sequencer/types"
)

// TODO: to be removed after the Proposer automatically adds such messages
func CmdArweaveBlockInfo() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "arweave-block-info [height] [timestamp] [hash]",
		Short: "Set last_arweave_block",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			height, err := strconv.ParseUint(args[0], 0, 64)
			if err != nil {
				return err
			}

			timestamp, err := strconv.ParseUint(args[1], 0, 64)
			if err != nil {
				return err
			}

			hash := []byte(args[2])

			msg := types.NewMsgArweaveBlockInfo(clientCtx.GetFromAddress().String(), height, timestamp, hash)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
