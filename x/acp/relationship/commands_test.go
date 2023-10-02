package relationship

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/sourcenetwork/sourcehub/testutil"
	authengineutil "github.com/sourcenetwork/sourcehub/testutil/auth_engine"
	"github.com/sourcenetwork/sourcehub/x/acp/auth_engine"
	"github.com/sourcenetwork/sourcehub/x/acp/types"
)

var timestamp = testutil.MustDateTimeToProto("2023-07-26 14:08:30")

var testPolicy = &types.Policy{
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
					},
				},
			},
			Permissions: []*types.Permission{},
		},
	},
	ActorResource: &types.ActorResource{
		Name: "actor",
	},
}

func setup(t *testing.T) (context.Context, auth_engine.AuthEngine) {
	ctx := context.Background()
	engine, _ := authengineutil.GetTestAuthEngine(t)

	rec, err := types.NewPolicyRecord(testPolicy)
	require.Nil(t, err)

	err = engine.SetPolicy(ctx, rec)
	require.Nil(t, err)

	_, err = engine.SetRelationship(ctx, testPolicy, &types.RelationshipRecord{
		Relationship: types.NewActorRelationship("test", "archived", "owner", "bob"),
		Archived:     true,
		Creator:      "bob",
	})
	require.Nil(t, err)

	_, err = engine.SetRelationship(ctx, testPolicy, &types.RelationshipRecord{
		Relationship: types.NewActorRelationship("test", "active", "owner", "bob"),
		Archived:     false,
		Creator:      "bob",
	})
	require.Nil(t, err)

	return ctx, engine
}

func TestRegisterObjectCommand_ValidObjectIsRegistered(t *testing.T) {
	ctx, engine := setup(t)

	cmd := RegisterObjectCommand{
		Registration: &types.Registration{
			Object: &types.Object{
				Resource: "test",
				Id:       "unregistered",
			},
			Actor: &types.Actor{
				Id: "bob",
			},
		},
		Policy:     testPolicy,
		CreationTs: timestamp,
	}

	result, err := cmd.Execute(ctx, engine)

	require.Equal(t, result, types.RegistrationResult_Registered)
	require.Nil(t, err)
}

func TestRegisterObjectCommand_CannotRegisterObjectThatHasBeenArchivedBySomeoneElse(t *testing.T) {
	ctx, engine := setup(t)

	cmd := RegisterObjectCommand{
		Registration: &types.Registration{
			Object: &types.Object{
				Resource: "test",
				Id:       "archived",
			},
			Actor: &types.Actor{
				Id: "alice",
			},
		},
		Policy:     testPolicy,
		CreationTs: timestamp,
	}

	result, err := cmd.Execute(ctx, engine)

	require.Equal(t, types.RegistrationResult_Denied.String(), result.String())
	require.Nil(t, err)
}

func TestRegisterObjectCommand_RegisteringActiveObjectOwnedBySomeoneElseErrors(t *testing.T) {
	ctx, engine := setup(t)

	cmd := RegisterObjectCommand{
		Registration: &types.Registration{
			Object: &types.Object{
				Resource: "test",
				Id:       "active",
			},
			Actor: &types.Actor{
				Id: "alice",
			},
		},
		Policy:     testPolicy,
		CreationTs: timestamp,
	}

	result, err := cmd.Execute(ctx, engine)

	require.Equal(t, result, types.RegistrationResult_Denied)
	require.Nil(t, err)
}

func TestRegisterObjectCommand_RegisteringArchivedObjectByOwnerActivatesIt(t *testing.T) {
	ctx, engine := setup(t)

	cmd := RegisterObjectCommand{
		Registration: &types.Registration{
			Object: &types.Object{
				Resource: "test",
				Id:       "archived",
			},
			Actor: &types.Actor{
				Id: "bob",
			},
		},
		Policy:     testPolicy,
		CreationTs: timestamp,
	}

	result, err := cmd.Execute(ctx, engine)

	require.Equal(t, types.RegistrationResult_Unarchived.String(), result.String())
	require.Nil(t, err)
}
