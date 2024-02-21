package policy

import (
	"fmt"

	"github.com/sourcenetwork/sourcehub/x/acp/types"
)

var (
	ErrInvalidPolicy = types.ErrAcpInput.Wrap("invalid policy")

	ErrUnknownMarshalingType = types.ErrAcpInput.Wrap("unknown marshaling type")
	ErrUnmarshaling          = types.ErrAcpInput.Wrap("unmarshaling error")

	ErrInvalidShortPolicy           = fmt.Errorf("invalid short policy: %w", ErrInvalidPolicy)
	ErrInvalidCreator               = fmt.Errorf("invalid creator: %w", ErrInvalidPolicy)
	ErrResourceMissingOwnerRelation = fmt.Errorf("resource missing owner relation: %w", ErrInvalidPolicy)
	ErrInvalidManagementRule        = fmt.Errorf("invalid relation managamente definition: %w", ErrInvalidPolicy)
)
