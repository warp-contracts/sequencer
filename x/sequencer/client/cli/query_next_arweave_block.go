package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"

	"github.com/warp-contracts/sequencer/x/sequencer/types"
)

func CmdListNextArweaveBlock() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-next-arweave-block",
		Short: "list all next_arweave_block",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllNextArweaveBlockRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.NextArweaveBlockAll(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, cmd.Use)
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdShowNextArweaveBlock() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-next-arweave-block [height]",
		Short: "shows a next_arweave_block",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			argHeight := args[0]

			params := &types.QueryGetNextArweaveBlockRequest{
				Height: argHeight,
			}

			res, err := queryClient.NextArweaveBlock(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
