package auth_engine

import (
	"context"
	"reflect"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/sourcenetwork/sourcehub/x/acp/types"
)

// SetupFunction is a Test setup function which is used to initialize an
// implementation of some AuthEngine.
// It is run before each individual test case
type SetupFunction func(t *testing.T) (context.Context, AuthEngine)

// TestAuthEngineImpl executes a series of tests for an AuthEngine.
// Uses the SetupFunction f in order to initialize the AuthEngine and execute
// the tests.
func TestAuthEngineImpl(t *testing.T, f SetupFunction) {
	suite := TestSuite{
		setup: f,
	}

	runSuite(t, &suite)
}

var policy = &types.Policy{
	Id: "1",
	Resources: []*types.Resource{
		&types.Resource{
			Name: "test",
			Relations: []*types.Relation{
				&types.Relation{
					Name: "owner",
					VrTypes: []*types.Restriction{
						&types.Restriction{
							ResourceName: "actor",
						},
						&types.Restriction{
							ResourceName: "group",
							RelationName: "owner",
						},
					},
				},
			},
			Permissions: nil,
		},
		&types.Resource{
			Name: "group",
			Relations: []*types.Relation{
				&types.Relation{
					Name: "owner",
					VrTypes: []*types.Restriction{
						&types.Restriction{
							ResourceName: "actor",
						},
					},
				},
			},
			Permissions: nil,
		},
	},
	ActorResource: &types.ActorResource{
		Name: "actor",
	},
}

type TestSuite struct {
	setup SetupFunction
}

func (s *TestSuite) setupPolicy(t *testing.T, ctx context.Context, engine AuthEngine) {
	rec, err := types.NewPolicyRecord(policy)
	require.Nil(t, err)

	err = engine.SetPolicy(ctx, rec)
	require.Nil(t, err)
}

func (s *TestSuite) TestSetAndGetActorRelationship(t *testing.T) {
	ctx, engine := s.setup(t)
	s.setupPolicy(t, ctx, engine)

	relationship := types.NewActorRelationship("test", "archived", "owner", "bob")

	found, err := engine.SetRelationship(ctx, policy, &types.RelationshipRecord{
		Relationship: relationship,
		Archived:     false,
	})
	require.Nil(t, err)
	require.Equal(t, bool(found), false)

	got, err := engine.GetRelationship(ctx, policy, relationship)
	require.Nil(t, err)
	require.Equal(t, got, &types.RelationshipRecord{
		Relationship: relationship,
		Archived:     false,
		PolicyId:     policy.Id,
	})
}

func (s *TestSuite) TestSetAndGetActorSetRelationship(t *testing.T) {
	ctx, engine := s.setup(t)
	s.setupPolicy(t, ctx, engine)

	relationship := types.NewActorSetRelationship("test", "id", "owner", "group", "groupId", "owner")

	found, err := engine.SetRelationship(ctx, policy, &types.RelationshipRecord{
		Relationship: relationship,
		Archived:     false,
	})
	require.Nil(t, err)
	require.Equal(t, bool(found), false)

	got, err := engine.GetRelationship(ctx, policy, relationship)
	require.Nil(t, err)
	require.Equal(t, got, &types.RelationshipRecord{
		Relationship: relationship,
		Archived:     false,
		PolicyId:     policy.Id,
	})
}

func (s *TestSuite) TestSetAndGetAllActorsRelationship(t *testing.T) {
	ctx, engine := s.setup(t)
	s.setupPolicy(t, ctx, engine)

	relationship := types.NewAllActorsRelationship("test", "id", "owner")

	found, err := engine.SetRelationship(ctx, policy, &types.RelationshipRecord{
		Relationship: relationship,
		Archived:     false,
	})
	require.Nil(t, err)
	require.Equal(t, bool(found), false)

	got, err := engine.GetRelationship(ctx, policy, relationship)
	require.Nil(t, err)
	require.Equal(t, got, &types.RelationshipRecord{
		Relationship: relationship,
		Archived:     false,
		PolicyId:     policy.Id,
	})
}

func (s *TestSuite) TestGetUnknownRelationship(t *testing.T) {
	ctx, engine := s.setup(t)
	s.setupPolicy(t, ctx, engine)

	relationship := types.NewActorRelationship("test", "id", "owner", "bob")

	record, err := engine.GetRelationship(ctx, policy, relationship)
	require.Nil(t, err)
	require.Nil(t, record)
}

func (s *TestSuite) TestSetAndGetPolicy(t *testing.T) {
	ctx, engine := s.setup(t)

	rec, err := types.NewPolicyRecord(policy)
	require.Nil(t, err)

	err = engine.SetPolicy(ctx, rec)
	require.Nil(t, err)

	record, err := engine.GetPolicy(ctx, policy.Id)
	require.Nil(t, err)
	require.Equal(t, record.Policy, policy)
}

// runSuite executes all Tests in TestSuite
func runSuite(t *testing.T, suite *TestSuite) {
	suiteVal := reflect.ValueOf(suite)
	suiteT := reflect.TypeOf(suite)

	var tests []reflect.Method
	for i := 0; i < suiteT.NumMethod(); i++ {
		method := suiteT.Method(i)
		if strings.HasPrefix(method.Name, "Test") {
			tests = append(tests, method)
		}
	}

	for _, test := range tests {
		f := test.Func
		t.Run(test.Name, func(t *testing.T) {
			tVal := reflect.ValueOf(t)
			args := []reflect.Value{suiteVal, tVal}
			f.Call(args)
		})
	}
}

// TODO Add Tests for Check, Filter Relationship and yada yada
// (should this be tested directly? seems redundant)
