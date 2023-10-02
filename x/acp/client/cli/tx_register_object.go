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

func CmdRegisterObject() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "register-object [creator] [policyId] [resource-name] [objId]",
		Short: "Broadcast message RegisterObject",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			_ = args[0]
			policyId := args[1]
			resource := args[2]
			objId := args[3]

			registration := &types.Registration{
				Object: &types.Object{
					Resource: resource,
					Id:       objId,
				},
				Actor: &types.Actor{
					Id: clientCtx.GetFromAddress().String(),
				},
			}

			msg := types.NewMsgRegisterObject(policyId, registration)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
