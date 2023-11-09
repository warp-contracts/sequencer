package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"

	"github.com/warp-contracts/sequencer/x/sequencer/types"
)

func CmdListPrevSortKey() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-prev-sort-key",
		Short: "list all prev-sort-key",
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

			params := &types.QueryAllPrevSortKeyRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.PrevSortKeyAll(cmd.Context(), params)
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

func CmdShowPrevSortKey() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-prev-sort-key [contract]",
		Short: "shows a prev-sort-key",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			argContract := args[0]

			params := &types.QueryGetPrevSortKeyRequest{
				Contract: argContract,
			}

			res, err := queryClient.PrevSortKey(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
