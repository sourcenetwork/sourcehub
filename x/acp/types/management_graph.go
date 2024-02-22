package types

import (
	"fmt"

	"github.com/sourcenetwork/sourcehub/utils"
)

// BuildManagementGraph builds a Management Graph from a Policy.
func (g *ManagementGraph) LoadFromPolicy(policy *Policy) {
	for _, resource := range policy.Resources {
		// register relation and managed rels
		for _, relation := range resource.Relations {
			g.registerRel(resource.Name, relation.Name)

			for _, managedRel := range relation.Manages {
				g.RegisterManagedRel(resource.Name, relation.Name, managedRel)
			}
		}
	}
}

// IsWellFormed walks through edges in graph and verifies whether the
// source and destination nodes for the edges are defined.
// If any edge is not defined, returns an error with the offending edges.
// If graph is well formed, return nil
func (g *ManagementGraph) IsWellFormed() error {
	for src, edgs := range g.ForwardEdges {
		for dst, _ := range edgs.Edges {

			_, src_ok := g.getNode(src)
			if !src_ok {
				return fmt.Errorf("edge defined from %v to %v: %v not found", src, dst, src)
			}

			_, dst_ok := g.getNode(dst)
			if !dst_ok {
				return fmt.Errorf("edge defined from %v to %v: %v not found", src, dst, dst)
			}
		}
	}
	return nil
}

// RegisterManagement adds a management rule in the Management graph.
// The registered rule states that for resource sourceRel manages managedRel.
func (g *ManagementGraph) RegisterManagedRel(resource, sourceRel, managedRel string) {
	srcId := g.buildNodeId(resource, sourceRel)
	dstId := g.buildNodeId(resource, managedRel)
	g.setEdg(srcId, dstId)
}

func (g *ManagementGraph) registerRel(resource, rel string) {
	node := &ManagerNode{
		Id:   g.buildNodeId(resource, rel),
		Text: rel,
	}
	g.setNode(node)
}

func (g *ManagementGraph) setEdg(src, dst string) {
	if g.ForwardEdges == nil {
		g.ForwardEdges = make(map[string]*ManagerEdges)
	}
	if g.ForwardEdges[src] == nil {
		g.ForwardEdges[src] = &ManagerEdges{}
	}
	if g.ForwardEdges[src].Edges == nil {
		g.ForwardEdges[src].Edges = make(map[string]bool)
	}
	g.ForwardEdges[src].Edges[dst] = true

	if g.BackwardEdges == nil {
		g.BackwardEdges = make(map[string]*ManagerEdges)
	}
	if g.BackwardEdges[dst] == nil {
		g.BackwardEdges[dst] = &ManagerEdges{}
	}
	if g.BackwardEdges[dst].Edges == nil {
		g.BackwardEdges[dst].Edges = make(map[string]bool)
	}
	g.BackwardEdges[dst].Edges[src] = true
}

func (g *ManagementGraph) setNode(node *ManagerNode) {
	if g.Nodes == nil {
		g.Nodes = make(map[string]*ManagerNode)
	}

	g.Nodes[node.Id] = node
}

func (g *ManagementGraph) getNode(id string) (*ManagerNode, bool) {
	if g.Nodes == nil {
		return nil, false
	}
	node, ok := g.Nodes[id]
	return node, ok
}

func (g *ManagementGraph) buildNodeId(resource, rel string) string {
	return resource + "/" + rel
}

// GetManagers return the list of relations which manages the relationName.
// Returns nil if relation wasn't found
func (g *ManagementGraph) GetManagers(resourceName, relationName string) []string {
	// The Management graph points from the manager relation to the managee
	// therefore the backward edges can be used to derive the managing relations

	relId := g.buildNodeId(resourceName, relationName)
	ancestors, ok := g.BackwardEdges[relId]
	if !ok {
		return nil
	}

	managers := make([]string, 0, len(ancestors.Edges))
	for id := range ancestors.Edges {
		node, _ := g.getNode(id)
		managers = append(managers, node.Text)
	}

	utils.SortSlice(managers)
	return managers
}
