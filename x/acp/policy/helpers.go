package policy

import (
	"github.com/sourcenetwork/sourcehub/x/acp/types"
)

func buildManagementGraph(policy *types.Policy) *types.ManagementGraph {
	graph := &types.ManagementGraph{}
	graph.LoadFromPolicy(policy)
	return graph
}
