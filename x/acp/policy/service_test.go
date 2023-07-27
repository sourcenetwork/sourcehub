package policy

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/sourcenetwork/sourcehub/testutil"
	"github.com/sourcenetwork/sourcehub/x/acp/types"
)

var creator = "cosmos1gue5de6a8fdff0jut08vw5sg9pk6rr00cstakj"
var sequence uint64 = 10
var timestamp = testutil.MustDateTimeToProto("2023-07-26 14:08:30")
var marshalType = types.PolicyMarshalingType_SHORT_YAML

func TestNewPolicyCreatesPolicy(t *testing.T) {
	policy := `
name: policy
description: ok
resources:
  foo:
    relations: 
      owner:
        doc: owner owns
        types:
          - blah
          - ok->that
        manages: 
          - reader
      reader:
    permissions: 
      own:
        expr: owner
        doc: own doc
      read: 
        expr: owner + reader
actor:
  name: actor-resource
  doc: my actor
          `
	marshalType := types.PolicyMarshalingType_SHORT_YAML

	got, err := NewPolicy(policy, marshalType, creator, sequence, timestamp)

	want := &types.Policy{
		Id:           "6157e44d1caf59a6b5627035ab32368f2fdcce418883bfdab1db7a1941f0ece2",
		Name:         "policy",
		Description:  "ok",
		CreationTime: timestamp,
		Creator:      creator,
		Resources: []*types.Resource{
			&types.Resource{
				Name: "foo",
				Relations: []*types.Relation{
					&types.Relation{
						Name: "owner",
						Doc:  "owner owns",
						VrTypes: []*types.Userset{
							&types.Userset{
								Resource: "blah",
								Relation: "",
							},
							&types.Userset{
								Resource: "ok",
								Relation: "that",
							},
						},
						Manages: []string{
							"reader",
						},
					},
					&types.Relation{
						Name: "reader",
					},
				},
				Permissions: []*types.Permission{
					&types.Permission{
						Name:       "own",
						Expression: "owner",
						Doc:        "own doc",
					},
					&types.Permission{
						Name:       "read",
						Expression: "owner + reader",
					},
				},
			},
		},
		ActorResource: &types.ActorResource{
			Name: "actor-resource",
			Doc:  "my actor",
		},
	}
	require.Nil(t, err)
	require.Equal(t, got, want)
}

func TestNewPolicyWithEmptyActorResourceNormalizesActorResourceName(t *testing.T) {
	policy := `
            name: policy
            resources:
              foo:
          `
	marshalType := types.PolicyMarshalingType_SHORT_YAML

	got, err := NewPolicy(policy, marshalType, creator, sequence, timestamp)

	want := &types.Policy{
		Id:           "aed9599dd6ea7e4507db99ace251c35936f89fcb3c954612c3485d62010305ce",
		Name:         "policy",
		CreationTime: timestamp,
		Creator:      creator,
                Resources: []*types.Resource{
                    &types.Resource{
                        Name: "foo",
                    },
                },
		ActorResource: &types.ActorResource{
			Name: DefaultActorResourceName,
			Doc:  "",
		},
	}
	require.Nil(t, err)
	require.Equal(t, got, want)
}

func TestNewPolicyErrorsWhenCreatorIsInvalid(t *testing.T) {
	policy := `
        name: policy
        resources:
          foo:
          `

	creator := "creator"
	got, err := NewPolicy(policy, marshalType, creator, sequence, timestamp)

	require.Nil(t, got)
	require.NotNil(t, err)
	require.True(t, errors.Is(err, ErrInvalidCreator))
}

func TestNewPolicyErrorsWithUnknownMarshalingType(t *testing.T) {
	policy := `
        name: policy
        resources:
          foo:
            relations: 
              owner:
          `

	got, err := NewPolicy(policy, types.PolicyMarshalingType_UNKNOWN, creator, sequence, timestamp)

	require.Nil(t, got)
	require.NotNil(t, err)
	require.True(t, errors.Is(err, ErrUnknownMarshalingType))
}

func TestNewPolicyErrorsWithMalformedManagementGraph(t *testing.T) {
	policy := `
        name: policy
        resources:
          foo:
            relations: 
              owner:
                manages: 
                  - whatever
          `

	got, err := NewPolicy(policy, marshalType, creator, sequence, timestamp)

	require.Nil(t, got)
	require.NotNil(t, err)
	require.True(t, errors.Is(err, ErrMalformedGraph))
}

