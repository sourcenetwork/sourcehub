package policy

import (
	"fmt"

	"github.com/sourcenetwork/sourcehub/x/acp/types"
)

func unmarshal(pol string, t types.PolicyMarshalingType) (*types.Policy, error) {
	var policy *types.Policy
	var err error

	shortUn := shortUnmarshaler{}
	switch t {
	case types.PolicyMarshalingType_SHORT_YAML:
		policy, err = shortUn.UnmarshalYAML(pol)
	case types.PolicyMarshalingType_SHORT_JSON:
		policy, err = shortUn.UnmarshalJSON(pol)
	default:
		err = ErrUnknownMarshalingType
	}
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrUnmarshaling, err)
	}

	return policy, nil
}
