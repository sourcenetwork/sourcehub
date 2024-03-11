package policy_cmd

import (
	"github.com/sourcenetwork/sourcehub/x/acp/types"
)

type CmdBuilder struct {
	cmd types.PolicyCmdPayload
}

func (b *CmdBuilder) Build() (types.PolicyCmdPayload, error) {
	return b.cmd, nil
}

func (b *CmdBuilder) Actor(did string) {
	// set did
	// validate
}

// take delta, add to current block height
func (b *CmdBuilder) ExpirationDelta() {}

func (b *CmdBuilder) SetRelationship()    {}
func (b *CmdBuilder) DeleteRelationship() {}
func (b *CmdBuilder) RegisterObject()     {}
func (b *CmdBuilder) UnregisterObject()   {}

// cmd signer
// receives a keyring, keyname and bingo bango

// serializer - receives signed command and output it serialized somehow
