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

func CmdDeleteRelationship() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-relationship [policy-id] [resource] [obj-id] [relation] [subj-id] [subj-resource] [subj-rel]",
		Short: "Broadcast message DeleteRelationship",
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
			subjId := args[4]
			subjRes := args[5]
			subjRel := args[6]
			var relationship *types.Relationship

			if subjRel != "" {
				relationship = types.NewActorSetRelationship(resource, objId, relation, subjRes, subjId, subjRel)
			} else if subjRes != "" {
				relationship = types.NewRelationship(resource, objId, relation, subjRes, subjId)
			} else {
				relationship = types.NewActorRelationship(resource, objId, relation, subjId)
			}

			msg := &types.MsgDeleteRelationship{
				Creator:      clientCtx.GetFromAddress().String(),
				PolicyId:     policyId,
				Relationship: relationship,
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
