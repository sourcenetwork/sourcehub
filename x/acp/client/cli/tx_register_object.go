package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	prototypes "github.com/cosmos/gogoproto/types"
	"github.com/spf13/cobra"

	"github.com/sourcenetwork/sourcehub/x/acp/types"
)

var _ = strconv.Itoa(0)

func CmdRegisterObject() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "register-object [policyId] [resource-name] [objId]",
		Short: "Broadcast message RegisterObject",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			policyId := args[0]
			resource := args[1]
			objId := args[2]

			msg := &types.MsgRegisterObject{
				Creator:  clientCtx.GetFromAddress().String(),
				PolicyId: policyId,
				Object: &types.Object{
					Resource: resource,
					Id:       objId,
				},
				CreationTime: prototypes.TimestampNow(),
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
