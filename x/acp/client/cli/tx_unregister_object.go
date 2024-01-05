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

			polId := args[0]
			resource := args[1]
			objId := args[2]

			msg := &types.MsgUnregisterObject{
				Creator:  clientCtx.GetFromAddress().String(),
				PolicyId: polId,
				Object: &types.Object{
					Resource: resource,
					Id:       objId,
				},
			}
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
