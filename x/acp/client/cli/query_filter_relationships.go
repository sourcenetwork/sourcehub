package cli

import (
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"

	"github.com/sourcenetwork/sourcehub/x/acp/types"
)

func CmdQueryFilterRelationships() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "relationships [policy-id] [object] [relation] [subject]",
		Short: "filters through relationships in a policy",
		Long: `Filters thourgh all relationships in a Policy. 
                Performs a lookup using the object, relation and subject filters.
                Uses a mini grammar as describe:
                object := resource:id | *
                relation := name | *
                subject := id | *
                Returns`,
		Args: cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			polId := args[0]
			object := args[1]
			relation := args[2]
			subject := args[3]

			queryClient := types.NewQueryClient(clientCtx)
			req := types.QueryFilterRelationshipsRequest{
				PolicyId: polId,
				Selector: buildSelector(object, relation, subject),
			}

			res, err := queryClient.FilterRelationships(cmd.Context(), &req)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func buildSelector(object, relation, subject string) *types.RelationshipSelector {
	objSelector := &types.ObjectSelector{}
	relSelector := &types.RelationSelector{}
	subjSelector := &types.SubjectSelector{}

	if object == "*" {
		objSelector.Selector = &types.ObjectSelector_Wildcard{
			Wildcard: &types.WildcardSelector{},
		}
	} else {
		res, id, _ := strings.Cut(object, ":")
		objSelector.Selector = &types.ObjectSelector_Object{
			Object: &types.Object{
				Resource: res,
				Id:       id,
			},
		}
	}

	if relation == "*" {
		relSelector.Selector = &types.RelationSelector_Wildcard{
			Wildcard: &types.WildcardSelector{},
		}
	} else {
		relSelector.Selector = &types.RelationSelector_Relation{
			Relation: relation,
		}
	}

	if subject == "*" {
		subjSelector.Selector = &types.SubjectSelector_Wildcard{
			Wildcard: &types.WildcardSelector{},
		}
	} else {
		subjSelector.Selector = &types.SubjectSelector_Subject{
			Subject: &types.Subject{
				Subject: &types.Subject_Actor{
					Actor: &types.Actor{
						Id: subject,
					},
				},
			},
		}
	}

	return &types.RelationshipSelector{
		ObjectSelector:   objSelector,
		RelationSelector: relSelector,
		SubjectSelector:  subjSelector,
	}
}
