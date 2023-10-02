package policy

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/sourcenetwork/sourcehub/x/acp/types"
)

func TestFullUnmarshal(t *testing.T) {
	in := `
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
                  - whatever
            permissions: 
              abc:
                expr: owner
                doc: abc doc
              def: 
                expr: owner + abc
        actor:
          name: actor-resource
          doc: my actor
          `
	unmarshaler := shortUnmarshaler{}
	out, err := unmarshaler.UnmarshalYAML(in)

	want := &types.Policy{
		Name:        "policy",
		Description: "ok",
		Resources: []*types.Resource{
			&types.Resource{
				Name: "foo",
				Relations: []*types.Relation{
					&types.Relation{
						Name: "owner",
						Doc:  "owner owns",
						VrTypes: []*types.Restriction{
							&types.Restriction{
								ResourceName: "blah",
								RelationName: "",
							},
							&types.Restriction{
								ResourceName: "ok",
								RelationName: "that",
							},
						},
						Manages: []string{
							"whatever",
						},
					},
				},
				Permissions: []*types.Permission{
					&types.Permission{
						Name:       "abc",
						Doc:        "abc doc",
						Expression: "owner",
					},
					&types.Permission{
						Name:       "def",
						Doc:        "",
						Expression: "owner + abc",
					},
				},
			},
		},
		ActorResource: &types.ActorResource{
			Name: "actor-resource",
			Doc:  "my actor",
		},
	}
	t.Log(err)
	require.Nil(t, err)
	if !reflect.DeepEqual(want, out) {
		t.Log(want.String())
		t.Log(out.String())
	}
	require.Equal(t, want, out)
}

func TestEmptyResourceMapsToResource(t *testing.T) {
	in := `
    resources:
      foo:
    `

	unmarshaler := shortUnmarshaler{}
	out, err := unmarshaler.UnmarshalYAML(in)

	want := &types.Policy{
		Resources: []*types.Resource{
			&types.Resource{
				Name: "foo",
			},
		},
	}
	require.Nil(t, err)
	require.Equal(t, want, out)
}

func TestResourceWithoutPermsOrRelsMapsToResource(t *testing.T) {
	in := `
    resources:
      foo:
        relations:
        permissions:
    `

	unmarshaler := shortUnmarshaler{}
	out, err := unmarshaler.UnmarshalYAML(in)

	want := &types.Policy{
		Resources: []*types.Resource{
			&types.Resource{
				Name:        "foo",
				Permissions: []*types.Permission{},
				Relations:   []*types.Relation{},
			},
		},
	}
	require.Nil(t, err)
	require.Equal(t, want, out)
}

func TestEmptyRelationMapsToRelation(t *testing.T) {
	in := `
    resources:
      foo:
        relations:
          blah:
    `

	unmarshaler := shortUnmarshaler{}
	out, err := unmarshaler.UnmarshalYAML(in)

	want := &types.Policy{
		Resources: []*types.Resource{
			&types.Resource{
				Name: "foo",
				Relations: []*types.Relation{
					&types.Relation{
						Name: "blah",
					},
				},
				Permissions: []*types.Permission{},
			},
		},
	}
	require.Nil(t, err)
	require.Equal(t, want, out)
}

func TestEmptyPermissionMapsToPermission(t *testing.T) {
	// NOTE The purpose of this test is to assert that the values are
	// correctly unmarshaled.
	// Therefore, even though a permission requires an expression,
	// it's ok because the validation will happen elsewhere.
	// Asserting the type unmarhsals correctly means that the validator -
	// as opposed to the unmarshaler - will error out leading to better error msgs.
	in := `
    resources:
      foo:
        permissions:
          blah:
    `

	unmarshaler := shortUnmarshaler{}
	out, err := unmarshaler.UnmarshalYAML(in)

	want := &types.Policy{
		Resources: []*types.Resource{
			&types.Resource{
				Name: "foo",
				Permissions: []*types.Permission{
					&types.Permission{
						Name: "blah",
					},
				},
				Relations: []*types.Relation{},
			},
		},
	}
	require.Nil(t, err)
	require.Equal(t, want, out)
}

func TestDuplicatedResourceErrors(t *testing.T) {
	in := `
    resources:
      foo:
      foo:
    `

	unmarshaler := shortUnmarshaler{}
	_, err := unmarshaler.UnmarshalYAML(in)

	require.NotNil(t, err)
}

func TestDuplicatedPermissionErrors(t *testing.T) {
	in := `
    resources:
      foo:
        permissions:
          read:
          read:
    `

	unmarshaler := shortUnmarshaler{}
	_, err := unmarshaler.UnmarshalYAML(in)

	require.NotNil(t, err)
}

func TestDuplicatedRelationErrors(t *testing.T) {
	in := `
    resources:
      foo:
        relations:
          reader:
          reader:
    `

	unmarshaler := shortUnmarshaler{}
	_, err := unmarshaler.UnmarshalYAML(in)

	require.NotNil(t, err)
}

func TestRestrictionIdentifierMapsBothForms(t *testing.T) {
	in := `
    resources:
      foo:
        relations:
          blah:
            types:
              - actor
              - book->owner
    `

	unmarshaler := shortUnmarshaler{}
	out, err := unmarshaler.UnmarshalYAML(in)

	want := &types.Policy{
		Resources: []*types.Resource{
			&types.Resource{
				Name: "foo",
				Relations: []*types.Relation{
					&types.Relation{
						Name: "blah",
						VrTypes: []*types.Restriction{
							&types.Restriction{
								ResourceName: "actor",
								RelationName: "",
							},
							&types.Restriction{
								ResourceName: "book",
								RelationName: "owner",
							},
						},
					},
				},
				Permissions: []*types.Permission{},
			},
		},
	}
	require.Nil(t, err)
	require.Equal(t, want, out)
}
