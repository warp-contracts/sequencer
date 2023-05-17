package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
	"github.com/warp-contracts/sequencer/x/sequencer/types"
)

var _ = strconv.Itoa(0)

func CmdArweave() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "arweave [format] [id] [last-tx] [owner] [tags] [target] [quantity] [data-root] [data-size] [data] [reward] [signature]",
		Short: "Broadcast message arweave",
		Args:  cobra.ExactArgs(12),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argFormat := args[0]
			argId := args[1]
			argLastTx := args[2]
			argOwner := args[3]
			argTags := args[4]
			argTarget := args[5]
			argQuantity := args[6]
			argDataRoot := args[7]
			argDataSize := args[8]
			argData := args[9]
			argReward := args[10]
			argSignature := args[11]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgArweave(
				clientCtx.GetFromAddress().String(),
				argFormat,
				argId,
				argLastTx,
				argOwner,
				argTags,
				argTarget,
				argQuantity,
				argDataRoot,
				argDataSize,
				argData,
				argReward,
				argSignature,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
