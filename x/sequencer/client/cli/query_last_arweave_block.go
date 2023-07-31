package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"

	"github.com/warp-contracts/sequencer/x/sequencer/types"
)

func CmdShowLastArweaveBlock() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-last-arweave-block",
		Short: "shows last_arweave_block",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryGetLastArweaveBlockRequest{}

			res, err := queryClient.LastArweaveBlock(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
