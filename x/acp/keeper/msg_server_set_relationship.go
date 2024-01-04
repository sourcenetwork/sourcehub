package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sourcenetwork/sourcehub/x/acp/types"
)

func (k msgServer) SetRelationship(goCtx context.Context, msg *types.MsgSetRelationship) (*types.MsgSetRelationshipResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgSetRelationshipResponse{}, nil
}
