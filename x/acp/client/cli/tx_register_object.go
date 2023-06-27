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
		Use:   "register-object [actor-did] [policyId] [namespace] [id]",
		Short: "Broadcast message RegisterObject",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

                        obj := &types.Entity {
                            Resource: args[2],
                            Id: args[3],
                        }
			msg := types.NewMsgRegisterObject(
				clientCtx.GetFromAddress().String(),
                                args[0],
                                args[1],
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
