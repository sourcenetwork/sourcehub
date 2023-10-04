package zanzi

import (
	"fmt"

	"github.com/sourcenetwork/sourcehub/utils"
	"github.com/sourcenetwork/zanzi/pkg/domain"

	"github.com/sourcenetwork/sourcehub/x/acp/types"
)

func newRelationshipMapper(actorResource string) relationshipMapper {
	return relationshipMapper{
		actorResource,
	}
}

type relationshipMapper struct {
	actorResource string
}

func (m *relationshipMapper) ToZanziRelationship(relationship *types.Relationship) *domain.Relationship {
	return &domain.Relationship{
		Object:   m.MapObject(relationship.Object),
		Relation: relationship.Relation,
		Subject:  m.MapSubject(relationship.Subject),
	}
}

func (m *relationshipMapper) MapObject(object *types.Object) *domain.Entity {
	return &domain.Entity{
		Resource: object.Resource,
		Id:       object.Id,
	}
}

func (m *relationshipMapper) MapSubject(subject *types.Subject) *domain.Subject {
	result := &domain.Subject{}

	switch s := subject.Subject.(type) {
	case *types.Subject_ActorSet:
		result.Subject = &domain.Subject_EntitySet{
			EntitySet: &domain.EntitySet{
				Entity:   m.MapObject(s.ActorSet.Object),
				Relation: s.ActorSet.Relation,
			},
		}
	case *types.Subject_Actor:
		result.Subject = &domain.Subject_Entity{
			Entity: &domain.Entity{
				Resource: m.actorResource,
				Id:       s.Actor.Id,
			},
		}
	case *types.Subject_AllActors:
		result.Subject = &domain.Subject_ResourceSet{
			ResourceSet: &domain.ResourceSet{
				ResourceName: m.actorResource,
			},
		}
            case *types.Subject_Object:
                result.Subject = &domain.Subject_Entity{
                    Entity: m.MapObject(s.Object),
                }
	}

	return result
}

func (m *relationshipMapper) ToZanziRelationshipRecord(record *types.RelationshipRecord) (*domain.RelationshipRecord, error) {
	bytes, err := record.Marshal()
	if err != nil {
		return nil, fmt.Errorf("mapping to zanzi relationship: %w", err)
	}

	return &domain.RelationshipRecord{
		Relationship: m.ToZanziRelationship(record.Relationship),
		AppData:      bytes,
	}, nil
}

func (m *relationshipMapper) FromZanziRelationship(zanziRecord *domain.RelationshipRecord) (*types.RelationshipRecord, error) {
	if zanziRecord == nil {
		return nil, nil
	}

	record := &types.RelationshipRecord{}

	err := record.Unmarshal(zanziRecord.AppData)
	if err != nil {
		return nil, fmt.Errorf("mapping from zanzi relationship: %w", err)
	}

	return record, nil
}

// policyMapper is a middleman between the ACP Policy type and Zanzi's Policy type.
type policyMapper struct{}

func (m *policyMapper) ToZanziRecord(record *types.PolicyRecord) (*domain.PolicyRecord, error) {
	zanziPolicy := m.ToZanzi(record.Policy)

	appData, err := record.Marshal()
	if err != nil {
		return nil, err
	}

	return &domain.PolicyRecord{
		Policy:  zanziPolicy,
		AppData: appData,
	}, nil
}

// Maps the module's Policy into Zanzi's type for storage
func (m *policyMapper) ToZanzi(policy *types.Policy) *domain.Policy {
	actorRes := m.actorToZanziResource(policy.ActorResource)

	resources := utils.MapSlice(policy.Resources, m.toZanziResource)
	resources = append(resources, actorRes)

	return &domain.Policy{
		Id:          policy.Id,
		Name:        policy.Name,
		Description: policy.Description,
		Resources:   resources,
		Attributes:  policy.Attributes,
	}
}

func (m *policyMapper) toZanziResource(resource *types.Resource) *domain.Resource {
	perms := utils.MapSlice(resource.Permissions, m.permToZanziRel)
	rels := utils.MapSlice(resource.Relations, m.toZanziRel)

	rels = append(rels, perms...)

	return &domain.Resource{
		Name:        resource.Name,
		Description: resource.Doc,
		Relations:   rels,
	}
}

// permToZanziRel maps an ACP Permission to a Zanzi Relation.
//
// Permissions are an ACP module specific concept and as far as Zanzi understands it,
// it is just another permission.
func (m *policyMapper) permToZanziRel(permission *types.Permission) *domain.Relation {
	return &domain.Relation{
		Name:        permission.Name,
		Description: permission.Doc,
		RelationExpression: &domain.RelationExpression{
			Expression: &domain.RelationExpression_Expr{
				Expr: permission.Expression,
			},
		},
		SubjectRestriction: &domain.SubjectRestriction{
			SubjectRestriction: &domain.SubjectRestriction_RestrictionSet{
				RestrictionSet: &domain.SubjectRestrictionSet{
					// By design it's impossible to create a relationships whose relation is a permission,
					// therefore ValueRestrictions for Permissions are meaningless.
					Restrictions: nil,
				},
			},
		},
	}
}

// toZanziRel maps an ACP Relation to a Zanzi Relation.
//
// Note the peculiarity regarding the RewriteExpression.
// By design, relations in the ACP module are pure relations,
// without any userset rewrite rules defined within them.
func (m *policyMapper) toZanziRel(relation *types.Relation) *domain.Relation {
	return &domain.Relation{
		Name:        relation.Name,
		Description: relation.Doc,
		RelationExpression: &domain.RelationExpression{
			Expression: &domain.RelationExpression_Tree{
				Tree: domain.ThisNode(),
			},
		},
		SubjectRestriction: &domain.SubjectRestriction{
			SubjectRestriction: &domain.SubjectRestriction_RestrictionSet{
				RestrictionSet: &domain.SubjectRestrictionSet{
					Restrictions: utils.MapSlice(relation.VrTypes, m.toZanziVR),
				},
			},
		},
	}

}

func (m *policyMapper) toZanziVR(vr *types.Restriction) *domain.SubjectRestrictionSet_Restriction {
	restriction := &domain.SubjectRestrictionSet_Restriction{}
	if vr.RelationName != "" {
		restriction.Entry = &domain.SubjectRestrictionSet_Restriction_EntitySet{
			EntitySet: &domain.EntitySetRestriction{
				ResourceName: vr.ResourceName,
				RelationName: vr.RelationName,
			},
		}
	} else {
		restriction.Entry = &domain.SubjectRestrictionSet_Restriction_Entity{
			Entity: &domain.EntityRestriction{
				ResourceName: vr.ResourceName,
			},
		}
	}

	return restriction
}

// actorToZanziResource maps the ActorResource to a Zanzi Resource.
//
// The actor resource is a convenient way of giving users flexibility
// to name their user namespace as they see fit but internall, it is treated
// as any other resource.
// Note a particularity, the Actor resource has no Permission defined within it.
// That makes sense as the actor namespace are terminal nodes in the relation graph.
func (m *policyMapper) actorToZanziResource(actor *types.ActorResource) *domain.Resource {
	rels := utils.MapSlice(actor.Relations, m.toZanziRel)
	return &domain.Resource{
		Name:        actor.Name,
		Description: actor.Doc,
		Relations:   rels,
	}
}

// Converts a Zanzi Policy into an ACP module Policy.
//
// The Policy is fetched from Zanzi's Policy AppData field.
func (m *policyMapper) FromZanzi(zanziRecord *domain.PolicyRecord) (*types.PolicyRecord, error) {
	record := &types.PolicyRecord{}

	err := record.Unmarshal(zanziRecord.AppData)
	if err != nil {
		return nil, err
	}

	return record, nil
}

func newSelectorMapper(mapper relationshipMapper) *selectorMapper {
	return &selectorMapper{
		relationshipMapper: mapper,
	}
}

// selectorMapper maps ACP RelationshipSelector to zanzi's RelationshipSelector
type selectorMapper struct {
	relationshipMapper relationshipMapper
}

// ToZanziSelector maps a selector to Zanzi's representation
func (m *selectorMapper) ToZanziSelector(selector *types.RelationshipSelector) (*domain.RelationshipSelector, error) {
	objSelector, err := m.MapObjectSelector(selector.ObjectSelector)
	if err != nil {
		return nil, err
	}

	relSelector, err := m.mapRelationSelector(selector.RelationSelector)
	if err != nil {
		return nil, err
	}

	subjSelector, err := m.mapSubjectSelector(selector.SubjectSelector)
	if err != nil {
		return nil, err
	}

	return &domain.RelationshipSelector{
		ObjectSelector:   objSelector,
		RelationSelector: relSelector,
		SubjectSelector:  subjSelector,
	}, nil
}

func (m *selectorMapper) MapObjectSelector(selector *types.ObjectSelector) (*domain.ObjectSelector, error) {
	zanziSelector := &domain.ObjectSelector{}

	switch selectorType := selector.Selector.(type) {
	case *types.ObjectSelector_Object:
		zanziSelector.Selector = &domain.ObjectSelector_ObjectSpec{
			ObjectSpec: &domain.Entity{
				Resource: selectorType.Object.Resource,
				Id:       selectorType.Object.Id,
			},
		}
	case *types.ObjectSelector_Wildcard:
		zanziSelector.Selector = &domain.ObjectSelector_Wildcard{
			Wildcard: &domain.WildcardSelector{},
		}
	default:
		return nil, fmt.Errorf("ObjectSelector %v: %w", selectorType, types.ErrInvalidVariant)
	}

	return zanziSelector, nil
}

func (m *selectorMapper) mapRelationSelector(selector *types.RelationSelector) (*domain.RelationSelector, error) {
	zanziSelector := &domain.RelationSelector{}

	switch selectorType := selector.Selector.(type) {
	case *types.RelationSelector_Relation:
		zanziSelector.Selector = &domain.RelationSelector_RelationName{
			RelationName: selectorType.Relation,
		}
	case *types.RelationSelector_Wildcard:
		zanziSelector.Selector = &domain.RelationSelector_Wildcard{
			Wildcard: &domain.WildcardSelector{},
		}
	default:
		return nil, fmt.Errorf("RelationSelector %v: %w", selectorType, types.ErrInvalidVariant)
	}
	return zanziSelector, nil
}

func (m *selectorMapper) mapSubjectSelector(selector *types.SubjectSelector) (*domain.SubjectSelector, error) {
	zanziSelector := &domain.SubjectSelector{}

	switch selectorType := selector.Selector.(type) {
	case *types.SubjectSelector_Subject:
		zanziSelector.Selector = &domain.SubjectSelector_SubjectSpec{
			SubjectSpec: m.relationshipMapper.MapSubject(selectorType.Subject),
		}
	case *types.SubjectSelector_Wildcard:
		zanziSelector.Selector = &domain.SubjectSelector_Wildcard{
			Wildcard: &domain.WildcardSelector{},
		}
	default:
		return nil, fmt.Errorf("SubjectSelector %v: %w", selectorType, types.ErrInvalidVariant)
	}

	return zanziSelector, nil
}
