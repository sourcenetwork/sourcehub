package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sourcenetwork/sourcehub/x/acp/types"
)

func (k msgServer) UnregisterObject(goCtx context.Context, msg *types.MsgUnregisterObject) (*types.MsgUnregisterObjectResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgUnregisterObjectResponse{}, nil
}
