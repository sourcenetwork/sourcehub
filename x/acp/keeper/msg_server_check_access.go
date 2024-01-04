package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sourcenetwork/sourcehub/x/acp/types"
)

func (k msgServer) CheckAccess(goCtx context.Context, msg *types.MsgCheckAccess) (*types.MsgCheckAccessResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgCheckAccessResponse{}, nil
}
