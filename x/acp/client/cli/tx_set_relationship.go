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
		Use:   "set-relationship [policy-id] [resource] [obj-id] [relation] [subject-resource] [subject-id] [subject-rel]",
		Short: "Broadcast message SetRelationship",
		Args:  cobra.ExactArgs(7),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			policyId := args[0]
			resource := args[1]
			objId := args[2]
			relation := args[3]
			subjResource := args[4]
			subjId := args[5]
			subjRel := args[6]

			var relationship *types.Relationship

			if subjRel == "" {
				relationship = types.NewRelationship(resource, objId, relation, subjResource, subjId)
			} else {
				relationship = types.NewActorSetRelationship(resource, objId, relation, subjResource, subjId, subjRel)

			}

			msg := types.NewMsgSetRelationshipNow(
				clientCtx.GetFromAddress().String(),
				policyId,
				relationship,
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
