package keeper

import (
	"context"

    "github.com/sourcenetwork/sourcehub/x/acp/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)


func (k msgServer) PolicyCmd(goCtx context.Context,  msg *types.MsgPolicyCmd) (*types.MsgPolicyCmdResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

    // TODO: Handling the message
    _ = ctx

	return &types.MsgPolicyCmdResponse{}, nil
}
