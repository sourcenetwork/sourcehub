package policy_cmd

import (
	"crypto"
	"fmt"

	"github.com/cosmos/gogoproto/jsonpb"
	prototypes "github.com/cosmos/gogoproto/types"
	"github.com/go-jose/go-jose/v3"
	"github.com/go-jose/go-jose/v3/cryptosigner"

	"github.com/sourcenetwork/sourcehub/x/acp/did"
	"github.com/sourcenetwork/sourcehub/x/acp/types"
)

func NewCmdBuilder() (CmdBuilder, error) {
	// TODO receive client to fetch current height and params
	return CmdBuilder{
		//params: params,
	}, nil
}

// CmdBuilder builds PolicyCmdPayloads
type CmdBuilder struct {
	cmd           types.PolicyCmdPayload
	params        types.Params
	currentHeight uint64
	cmdErr        error
	signer        crypto.Signer
}

// BuildJWS produces a signed JWS for the specified Cmd
func (b *CmdBuilder) BuildJWS() (string, error) {
	if b.signer == nil {
		return "", fmt.Errorf("CmdBuilder failed: %w", ErrSignerRequired)
	}

	payload, err := b.Build()
	if err != nil {
		return "", err
	}

	return SignPayload(payload, b.signer)
}

// SetSigner sets the Signer for the Builder, which will be used to produce a JWS
func (b *CmdBuilder) SetSigner(signer crypto.Signer) {
	b.signer = signer
}

// GetSigner returns the currently set Signer
func (b *CmdBuilder) GetSigner() crypto.Signer {
	return b.signer
}

// Build validates the data provided to the Builder, validates it and returns a PolicyCmdPayload or an error.
func (b *CmdBuilder) Build() (types.PolicyCmdPayload, error) {
	b.cmd.IssuedHeight = b.currentHeight

	if b.cmd.CreationTime == nil {
		b.cmd.CreationTime = prototypes.TimestampNow()
	}

	if b.cmd.ExpirationDelta == 0 {
		b.cmd.ExpirationDelta = b.params.PolicyCommandMaxExpirationDelta
	}

	if b.cmd.PolicyId == "" {
		return types.PolicyCmdPayload{}, fmt.Errorf("CmdBuilder: policy id: %w", ErrBuilderMissingArgument)
	}

	if b.cmd.ExpirationDelta > b.params.PolicyCommandMaxExpirationDelta {
		return types.PolicyCmdPayload{}, fmt.Errorf("CmdBuilder: %v", ErrExpirationDeltaTooLarge)
	}

	if err := did.IsValidDID(b.cmd.Actor); err != nil {
		return types.PolicyCmdPayload{}, fmt.Errorf("CmdBuilder: invalid actor: %v", err)
	}

	if b.cmd.Cmd == nil {
		return types.PolicyCmdPayload{}, fmt.Errorf("CmdBuilder: Command not specified: %v", ErrBuilderMissingArgument)
	}

	if b.cmdErr != nil {
		// TODO validate commands
		return types.PolicyCmdPayload{}, fmt.Errorf("CmdBuilder: Command invalid: %v", b.cmdErr)
	}

	return b.cmd, nil
}

// CreationTimestamp sets the creation timestamp
func (b *CmdBuilder) CreationTimestamp(ts *prototypes.Timestamp) {
	b.cmd.CreationTime = ts
}

// Actor sets the Actor for the Command
func (b *CmdBuilder) Actor(did string) {
	b.cmd.Actor = did
}

// ExpirationDelta specifies the number of blocks after the issue height for which the Command will be valid.
func (b *CmdBuilder) ExpirationDelta(delta uint64) {
	b.cmd.ExpirationDelta = delta
}

// PolicyID sets the Policy ID for the payload
func (b *CmdBuilder) PolicyID(id string) {
	b.cmd.PolicyId = id
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

	opaque := cryptosigner.Opaque(skey)
	key := jose.SigningKey{
		Algorithm: opaque.Algs()[0],
		Key:       opaque,
	}
	var opts *jose.SignerOptions
	signer, err := jose.NewSigner(key, opts)
	if err != nil {
		return "", err
	}

	obj, err := signer.Sign([]byte(payload))
	if err != nil {
		return "", err
	}

	return obj.FullSerialize(), nil
}
