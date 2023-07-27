package policy

import (
	zanzi "github.com/sourcenetwork/zanzi/pkg/core"

	"github.com/sourcenetwork/sourcehub/utils"
	"github.com/sourcenetwork/sourcehub/x/acp/types"
)

// relationRewriteExpr stores the RewriteExpression for ACP Relation objects.
//
// By design relations are always simple expression, meaning they don't expand
// to any other relations in Zanzibar, thefore their rewrite expr is always "_this"
const relationRewriteExpr string = "_this"

// mapper is a middleman between the ACP Policy type and Zanzi's Policy type.
type mapper struct{}

// Maps the module's Policy into Zanzi's type for storage
func (m *mapper) ToZanzi(policy *types.Policy) (*zanzi.Policy, error) {
	actorRes := m.actorToZanziResource(policy.ActorResource)

	resources := utils.MapSlice(policy.Resources, m.toZanziResource)
	resources = append(resources, actorRes)

	polData := m.buildPolicyData(policy)

	polDataBytes, err := polData.Marshal()
	if err != nil {
		return nil, err
	}

	return &zanzi.Policy{
		Id:          policy.Id,
		Name:        policy.Name,
		Description: policy.Description,
		Resources:   resources,
		Attributes:  policy.Attributes,
		AppData:     polDataBytes,
	}, err
}

func (m *mapper) buildPolicyData(pol *types.Policy) *types.PolicyData {
	graph := buildManagementGraph(pol)

	return &types.PolicyData{
		AcpPolicy:       pol,
		ManagementGraph: graph,
	}
}

func (m *mapper) toZanziResource(resource *types.Resource) *zanzi.Resource {
	perms := utils.MapSlice(resource.Permissions, m.permToZanziRel)
	rels := utils.MapSlice(resource.Relations, m.toZanziRel)

	rels = append(rels, perms...)

	return &zanzi.Resource{
		Name:        resource.Name,
		Description: resource.Doc,
		Relations:   rels,
	}
}

// permToZanziRel maps an ACP Permission to a Zanzi Relation.
//
// Permissions are an ACP module specific concept and as far as Zanzi understands it,
// it is just another permission.
func (m *mapper) permToZanziRel(permission *types.Permission) *zanzi.Relation {
	return &zanzi.Relation{
		Name:              permission.Name,
		Description:       permission.Doc,
		RewriteExpression: permission.Expression,
		// By design it's impossible to create a relationships whose relation is a permission,
		// therefore ValueRestrictions for Permissions are meaningless.
		ValueRestrictions: nil,
	}
}

// toZanziRel maps an ACP Relation to a Zanzi Relation.
//
// Note the peculiarity regarding the RewriteExpression.
// By design, relations in the ACP module are pure relations,
// without any userset rewrite rules defined within them.
func (m *mapper) toZanziRel(relation *types.Relation) *zanzi.Relation {
	return &zanzi.Relation{
		Name:              relation.Name,
		Description:       relation.Doc,
		RewriteExpression: relationRewriteExpr,
		ValueRestrictions: utils.MapSlice(relation.VrTypes, m.toZanziVR),
	}
}

func (m *mapper) toZanziVR(vr *types.Userset) *zanzi.ValueRestriction {
	return &zanzi.ValueRestriction{
		ResourceName: vr.Resource,
		RelationName: vr.Relation,
	}
}

// actorToZanziResource maps the ActorResource to a Zanzi Resource.
//
// The actor resource is a convenient way of giving users flexibility
// to name their user namespace as they see fit but internall, it is treated
// as any other resource.
// Note a particularity, the Actor resource has no Permission defined within it.
// That makes sense as the actor namespace are terminal nodes in the relation graph.
func (m *mapper) actorToZanziResource(actor *types.ActorResource) *zanzi.Resource {
	rels := utils.MapSlice(actor.Relations, m.toZanziRel)
	return &zanzi.Resource{
		Name:        actor.Name,
		Description: actor.Doc,
		Relations:   rels,
	}
}

// Converts a Zanzi Policy into an ACP module Policy.
//
// The Policy is fetched from Zanzi's Policy AppData field.
func (m *mapper) FromZanzi(policy *zanzi.Policy) (*types.Policy, error) {
	data := types.PolicyData{}

	err := data.Unmarshal(policy.AppData)
	if err != nil {
		return nil, err
	}

	return data.AcpPolicy, nil
}
