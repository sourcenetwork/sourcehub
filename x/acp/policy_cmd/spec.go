package policy_cmd

import (
	"github.com/sourcenetwork/sourcehub/x/acp/types"
)

// JWSCmdSpec validates a JWS encoding of a SignedPolicyCmd
type JWSCmdSpec struct{}

func (s *JWSCmdSpec) Satisfies() error { panic("todo") }

// ValidateAndExtractCmd validates a MsgPolicyCmd and return the Cmd payload
func ValidateAndExtractCmd(msg types.MsgPolicyCmd_SignedCmd) (*types.PolicyCmdPayload, error) {
	panic("todo")
	return nil, nil
}

// SignedPolicyCmdSpec validates a SignedPolicyCmd
type SignedPolicyCmdSpec struct{}

func (s *SignedPolicyCmdSpec) Satisfies() error { panic("todo") }
