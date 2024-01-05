package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"

	"github.com/sourcenetwork/sourcehub/x/acp/types"
)

func CmdQueryVerifyAccessRequest() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "verify-access-request [actor] {operations}",
		Short: "verifies an access request against a policy and its relationships",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)
			_ = clientCtx
			_ = queryClient
			return nil
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
