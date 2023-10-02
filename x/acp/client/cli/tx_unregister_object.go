package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/sourcenetwork/sourcehub/x/acp/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdUnregisterObject() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "unregister-object [policy-id] [resource] [id]",
		Short: "Broadcast message UnregisterObject",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			obj := &types.Object{
				Resource: args[2],
				Id:       args[3],
			}

			msg := types.NewMsgUnregisterObject(
				clientCtx.GetFromAddress().String(),
				args[0],
				obj,
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
