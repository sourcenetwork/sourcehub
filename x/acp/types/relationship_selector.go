package types

// RelationshipSelectorBuilder builds RelationshipSelector filters
type RelationshipSelectorBuilder struct {
	objSelector      *ObjectSelector
	relationSelector *RelationSelector
	subjSelector     *SubjectSelector
}

// Object sets the Selector to filter Relationships whose Object are obj
func (b *RelationshipSelectorBuilder) Object(obj *Object) *RelationshipSelectorBuilder {
	b.objSelector = &ObjectSelector{
		Selector: &ObjectSelector_Object{
			Object: obj,
		},
	}
	return b
}

// AnyObject configures the Selector to not filter Relationship's Objects
func (b *RelationshipSelectorBuilder) AnyObject() *RelationshipSelectorBuilder {
	b.objSelector = &ObjectSelector{
		Selector: &ObjectSelector_Wildcard{
			Wildcard: &WildcardSelector{},
		},
	}
	return b
}

// Relation configures the Selector to filter for Relationships whose Relation are rel
func (b *RelationshipSelectorBuilder) Relation(rel string) *RelationshipSelectorBuilder {
	b.relationSelector = &RelationSelector{
		Selector: &RelationSelector_Relation{
			Relation: rel,
		},
	}
	return b
}

// AnyRelation configures the Selector to not filter Relationship's Relations
func (b *RelationshipSelectorBuilder) AnyRelation(rel string) *RelationshipSelectorBuilder {
	b.relationSelector = &RelationSelector{
		Selector: &RelationSelector_Wildcard{
			Wildcard: &WildcardSelector{},
		},
	}
	return b
}

// Actor configures the Selector to filter for Relationships whose Subjects match the given Actor
func (b *RelationshipSelectorBuilder) Actor(actorId string) *RelationshipSelectorBuilder {
	b.subjSelector = &SubjectSelector{
		Selector: &SubjectSelector_Subject{
			Subject: &Subject{
				&Subject_Actor{
					Actor: &Actor{
						Id: actorId,
					},
				},
			},
		},
	}
	return b
}

// Subject configures the Selector to filter for Relationships whose Subjects match subject
func (b *RelationshipSelectorBuilder) Subject(subject *Subject) *RelationshipSelectorBuilder {
	b.subjSelector = &SubjectSelector{
		Selector: &SubjectSelector_Subject{
			Subject: subject,
		},
	}
	return b
}

// AnySubject configures the Selector to not filter Relationship's Subjects
func (b *RelationshipSelectorBuilder) AnySubject() *RelationshipSelectorBuilder {
	b.subjSelector = &SubjectSelector{
		Selector: &SubjectSelector_Wildcard{
			Wildcard: &WildcardSelector{},
		},
	}
	return b
}

func (b *RelationshipSelectorBuilder) Build() RelationshipSelector {
	return RelationshipSelector{
		ObjectSelector:   b.objSelector,
		RelationSelector: b.relationSelector,
		SubjectSelector:  b.subjSelector,
	}
}
