package keeper

import (
	"context"
	"encoding/base64"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sourcenetwork/sourcehub/x/bulletin/types"
)

func (k msgServer) CreatePost(goCtx context.Context, msg *types.MsgCreatePost) (*types.MsgCreatePostResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	post := types.Post{
		Namespace: msg.Namespace,
		Payload:   msg.Payload,
		Proof:     msg.Proof,
	}

	k.AddPost(ctx, post)

	b64Payload := base64.StdEncoding.EncodeToString(post.Payload)
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventNewPostValue,
			sdk.NewAttribute(types.AttributeKeyNamespace, post.Namespace),
			sdk.NewAttribute(types.AttributeKeyPayload, b64Payload),
			sdk.NewAttribute(types.AttributeKeyProof, "PLACEHOLDER"),
		),
	)

	return &types.MsgCreatePostResponse{}, nil
}
