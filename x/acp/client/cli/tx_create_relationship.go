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

func CmdCreateRelationship() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-relationship [policy-id] [resource] [obj-id] [relation] [actor-resource] [actor-id] [actor-rel]",
		Short: "Broadcast message CreateRelationship",
		Args:  cobra.ExactArgs(7),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			obj := &types.Object{
				Resource: args[1],
				Id:       args[2],
			}

			subject := &types.Subject{}
			/*
			                            &types.Object{
							Resource: args[5],
							Id:       args[6],
						}
			*/

			msg := types.NewMsgCreateRelationship(
				clientCtx.GetFromAddress().String(),
				args[0],
				obj,
				args[3],
				subject,
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
