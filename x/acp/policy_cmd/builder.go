package policy_cmd

import (
	"crypto"

	"github.com/cosmos/gogoproto/jsonpb"
	"github.com/go-jose/go-jose/v3/cryptosigner"

	"github.com/sourcenetwork/sourcehub/x/acp/types"
)

func NewCmdBuilder(params types.Params) CmdBuilder {
	// TODO receive client to fetch current height
	return CmdBuilder{
		params: params,
	}
}

func newCmdBuilderWithHeight(params types.Params, currentHeight uint64) CmdBuilder {
	return CmdBuilder{
		params:        params,
		currentHeight: currentHeight,
	}
}

// CmdBuilder builds PolicyCmdPayloads
type CmdBuilder struct {
	cmd           types.PolicyCmdPayload
	params        types.Params
	currentHeight uint64
	cmdErr        error
}

// Build validates the data provided to the Builder, validates it and returns a PolicyCmdPayload or an error.
func (b *CmdBuilder) Build() (types.PolicyCmdPayload, error) {
	// validate actor did
	// validate delta isn't over the max
	// valid subobjects

	return b.cmd, nil
}

// Actor sets the Actor for the Command
func (b *CmdBuilder) Actor(did string) {
	b.cmd.Actor = did
}

// ExpirationDelta specifies the number of blocks after the issue height for which the Command will be valid.
func (b *CmdBuilder) ExpirationDelta(delta uint64) {
	//b.cmd.
	// take delta, add to current block height
	// max should come from params / default
}

// SetRelationship builds a Payload for a SetRelationship command
func (b *CmdBuilder) SetRelationship(relationship *types.Relationship) {
	b.cmd.Cmd = &types.PolicyCmdPayload_SetRelationshipCmd{
		SetRelationshipCmd: &types.SetRelationshipCmd{
			Relationship: relationship,
		},
	}
}

// DeleteRelationship builds a Payload for a DeleteRelationship command
func (b *CmdBuilder) DeleteRelationship(relationship *types.Relationship) {
	b.cmd.Cmd = &types.PolicyCmdPayload_DeleteRelationshipCmd{
		DeleteRelationshipCmd: &types.DeleteRelationshipCmd{
			Relationship: relationship,
		},
	}
}

// RegisterObject builds a Payload for a RegisterObject command
func (b *CmdBuilder) RegisterObject(obj *types.Object) {
	b.cmd.Cmd = &types.PolicyCmdPayload_RegisterObjectCmd{
		RegisterObjectCmd: &types.RegisterObjectCmd{
			Object: obj,
		},
	}
}

// UnregisterObject builds a Payload for a UnregisterObject command
func (b *CmdBuilder) UnregisterObject(obj *types.Object) {
	b.cmd.Cmd = &types.PolicyCmdPayload_UnregisterObjectCmd{
		UnregisterObjectCmd: &types.UnregisterObjectCmd{
			Object: obj,
		},
	}
}

// SignPayload produces a JWS serialized version of a Payload from a signing key
func SignPayload(cmd types.PolicyCmdPayload, skey crypto.Signer) (string, error) {
	marshaler := jsonpb.Marshaler{}
	payload, err := marshaler.MarshalToString(&cmd)
	if err != nil {
		return "", err
	}

	signer := cryptosigner.Opaque(skey)
	jws, err := signer.SignPayload([]byte(payload), signer.Algs()[0])
	if err != nil {
		return "", err
	}

	return string(jws), nil
}
