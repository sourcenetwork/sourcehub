package relationship

import (
	"github.com/sourcenetwork/sourcehub/x/acp/types"
)

var (
	ErrDeleteOwnerRel      = types.ErrAcpProtocolViolation.Wrapf("cannot delete an owner relationship")
	ErrSetOwnerRel         = types.ErrAcpProtocolViolation.Wrapf("cannot set an owner relationship")
	ErrInvalidRelationship = types.ErrAcpInput.Wrapf("invalid relationship")
)
