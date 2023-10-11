package cli

import (
	"fmt"
	"github.com/spf13/cobra"
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/cosmos/cosmos-sdk/version"
	authtx "github.com/cosmos/cosmos-sdk/x/auth/tx"

	"github.com/warp-contracts/sequencer/x/sequencer/types"
)

const (
	flagEvents = "events"
	flagType   = "type"

	typeHash       = "hash"
	typeAccSeq     = "acc_seq"
	typeDataItemId = "data_item_id"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string) *cobra.Command {
	// Group sequencer queries under a subcommand
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdQueryParams())
	cmd.AddCommand(CmdShowLastArweaveBlock())
	cmd.AddCommand(CmdListLastSortKey())
	cmd.AddCommand(CmdShowLastSortKey())
	// this line is used by starport scaffolding # 1

	return cmd
}

// QueryTxCmd overrides the default implementation, allowing for transaction search by data item id instead of the signature
func QueryTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "tx --type=[hash|acc_seq|data_item_id] [hash|acc_seq|data_item_id]",
		Short: "Query for a transaction by hash, \"<addr>/<seq>\" combination or data item id in a committed block",
		Long: strings.TrimSpace(fmt.Sprintf(`
Example:
$ %s query tx <hash>
$ %s query tx --%s=%s <addr>/<sequence>
$ %s query tx --%s=%s <data_item_id>
`,
			version.AppName,
			version.AppName, flagType, typeAccSeq,
			version.AppName, flagType, typeDataItemId)),
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			typ, _ := cmd.Flags().GetString(flagType)

			switch typ {
			case typeHash:
				{
					if args[0] == "" {
						return fmt.Errorf("argument should be a tx hash")
					}

					// If hash is given, then query the tx by hash.
					output, err := authtx.QueryTx(clientCtx, args[0])
					if err != nil {
						return err
					}

					if output.Empty() {
						return fmt.Errorf("no transaction found with hash %s", args[0])
					}

					return clientCtx.PrintProto(output)
				}
			case typeAccSeq:
				{
					if args[0] == "" {
						return fmt.Errorf("`acc_seq` type takes an argument '<addr>/<seq>'")
					}

					tmEvents := []string{
						fmt.Sprintf("%s.%s='%s'", sdk.EventTypeTx, sdk.AttributeKeyAccountSequence, args[0]),
					}
					txs, err := authtx.QueryTxsByEvents(clientCtx, tmEvents, query.DefaultPage, query.DefaultLimit, "")
					if err != nil {
						return err
					}
					if len(txs.Txs) == 0 {
						return fmt.Errorf("found no txs matching given address and sequence combination")
					}
					if len(txs.Txs) > 1 {
						// This case means there's a bug somewhere else in the code. Should not happen.
						return fmt.Errorf("found %d txs matching given address and sequence combination", len(txs.Txs))
					}

					return clientCtx.PrintProto(txs.Txs[0])
				}
			case typeDataItemId:
				{
					if args[0] == "" {
						return fmt.Errorf("argument should be a data item id")
					}

					tmEvents := []string{
						fmt.Sprintf("%s.%s='%s'", sdk.EventTypeTx, types.AttributeKeyDataItemId, args[0]),
					}
					txs, err := authtx.QueryTxsByEvents(clientCtx, tmEvents, query.DefaultPage, query.DefaultLimit, "")
					if err != nil {
						return err
					}
					if len(txs.Txs) == 0 {
						return fmt.Errorf("found no txs matching given data item id")
					}
					if len(txs.Txs) > 1 {
						// This case means there's a bug somewhere else in the code. Should not happen.
						return fmt.Errorf("found %d txs matching given data item id", len(txs.Txs))
					}

					return clientCtx.PrintProto(txs.Txs[0])
				}
			default:
				return fmt.Errorf("unknown --%s value %s", flagType, typ)
			}
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	cmd.Flags().String(flagType, typeHash, fmt.Sprintf("The type to be used when querying tx, can be one of \"%s\", \"%s\", \"%s\"", typeHash, typeAccSeq, typeDataItemId))

	return cmd
}
