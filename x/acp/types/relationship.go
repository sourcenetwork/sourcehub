package types

func NewPolicyRecord(policy *Policy) (*PolicyRecord, error) {
	graph := &ManagementGraph{}
	graph.LoadFromPolicy(policy)
	if err := graph.IsWellFormed(); err != nil {
		return nil, err
	}

	return &PolicyRecord{
		Policy:          policy,
		ManagementGraph: graph,
	}, nil
}

func NewRelationship(resource, objId, relation, subjResource, subjId string) *Relationship {
	return &Relationship{
		Object: &Object{
			Resource: resource,
			Id:       objId,
		},
		Relation: relation,
		Subject: &Subject{
			Subject: &Subject_Object{
				Object: &Object{
                                    Resource: subjResource,
					Id: subjId,
				},
			},
		},
	}
}

func NewActorRelationship(resource, objId, relation, actor string) *Relationship {
	return &Relationship{
		Object: &Object{
			Resource: resource,
			Id:       objId,
		},
		Relation: relation,
		Subject: &Subject{
			Subject: &Subject_Actor{
				Actor: &Actor{
					Id: actor,
				},
			},
		},
	}
}

func NewActorSetRelationship(resource, objId, relation, subjResource, subjId, subjRel string) *Relationship {
	return &Relationship{
		Object: &Object{
			Resource: resource,
			Id:       objId,
		},
		Relation: relation,
		Subject: &Subject{
			Subject: &Subject_ActorSet{
				ActorSet: &ActorSet{
					Object: &Object{
						Resource: subjResource,
						Id:       subjId,
					},
					Relation: subjRel,
				},
			},
		},
	}
}

func NewAllActorsRelationship(resource, objId, relation string) *Relationship {
	return &Relationship{
		Object: &Object{
			Resource: resource,
			Id:       objId,
		},
		Relation: relation,
		Subject: &Subject{
			Subject: &Subject_AllActors{
				AllActors: &AllActors{},
			},
		},
	}
}
