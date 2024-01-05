package relationship

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/sourcenetwork/sourcehub/x/acp/auth_engine"
	"github.com/sourcenetwork/sourcehub/x/acp/policy"
	"github.com/sourcenetwork/sourcehub/x/acp/testutil"
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
	engine, _, ctx := testutil.GetTestAuthEngine(t)

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

	require.ErrorIs(t, err, types.ErrNotAuthorized)
	require.Equal(t, result, types.RegistrationResult_NoOp)
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

	require.ErrorIs(t, err, types.ErrNotAuthorized)
	require.Equal(t, result, types.RegistrationResult_NoOp)
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

// setup for DeleteRelationship tests
// sets alice as the owner of readme.txt and sets bob as an admin of readme.txt
func setupTestDeleteRelationship(t *testing.T) (context.Context, auth_engine.AuthEngine, *RelationshipAuthorizer, *types.Policy) {
	polStr := `
    id: set-rel-pol
    resources:
      file:
        permissions:
        relations:
          owner:
            types:
              - actor
          reader:
            types:
              - actor
          admin:
            manages:
              - reader
            types:
              - actor

    actor:
      name: actor
    `
	polIR, err := policy.Unmarshal(polStr, types.PolicyMarshalingType_SHORT_YAML)
	require.Nil(t, err)

	engine, _, ctx := testutil.GetTestAuthEngine(t)

	createCmd := policy.CreatePolicyCommand{
		Policy:       polIR,
		Creator:      "cosmos1gue5de6a8fdff0jut08vw5sg9pk6rr00cstakj",
		CreationTime: timestamp,
	}
	pol, err := createCmd.Execute(ctx, &testutil.AccountKeeperStub{}, engine)
	require.Nil(t, err)

	_, err = engine.SetRelationship(ctx, pol, &types.RelationshipRecord{
		Relationship: types.NewActorRelationship("file", "readme.txt", "owner", "alice"),
		Archived:     false,
		Creator:      "alice",
	})
	require.Nil(t, err)

	_, err = engine.SetRelationship(ctx, pol, &types.RelationshipRecord{
		Relationship: types.NewActorRelationship("file", "readme.txt", "admin", "bob"),
		Archived:     false,
		Creator:      "bob",
	})
	require.Nil(t, err)

	authorizer := NewRelationshipAuthorizer(engine)
	return ctx, engine, authorizer, pol
}

func TestDeleteRelationship_AttemptingToDeleteAnOwnerRelationshipAsOwnerErrors(t *testing.T) {
	ctx, engine, authorizer, pol := setupTestDeleteRelationship(t)

	cmd := DeleteRelationshipCommand{
		Policy:       pol,
		Relationship: types.NewActorRelationship("file", "readme.txt", "owner", "alice"),
		Actor:        "alice",
	}
	result, err := cmd.Execute(ctx, engine, authorizer)

	require.Equal(t, bool(result), false)
	require.ErrorIs(t, err, ErrDeleteOwnerRel)
}

func TestDeleteRelationship_AttemptToDeleteARelationshipForANonExistingObjectReturnsNotAuthorized(t *testing.T) {
	ctx, engine, authorizer, pol := setupTestDeleteRelationship(t)

	cmd := DeleteRelationshipCommand{
		Policy:       pol,
		Relationship: types.NewActorRelationship("file", "none", "reader", "bob"),
		Actor:        "bob",
	}
	result, err := cmd.Execute(ctx, engine, authorizer)

	require.False(t, bool(result))
	require.ErrorIs(t, err, types.ErrNotAuthorized)
}

func TestDeleteRelationship_AttemptingToDeleteARelationshipActorDoesNotManageReturnsUnauthorized(t *testing.T) {
	ctx, engine, authorizer, pol := setupTestDeleteRelationship(t)

	cmd := DeleteRelationshipCommand{
		Policy:       pol,
		Relationship: types.NewActorRelationship("file", "readme.txt", "admin", "bob"),
		Actor:        "bob",
	}
	result, err := cmd.Execute(ctx, engine, authorizer)

	require.False(t, bool(result))
	require.ErrorIs(t, err, types.ErrNotAuthorized)
}

func TestDeleteRelationship_AttemptingToDeleteARelationshipWithUnknownRelationErrors(t *testing.T) {
	ctx, engine, authorizer, pol := setupTestDeleteRelationship(t)

	cmd := DeleteRelationshipCommand{
		Policy:       pol,
		Relationship: types.NewActorRelationship("file", "readme.txt", "unknown-relation", "bob"),
		Actor:        "bob",
	}
	result, err := cmd.Execute(ctx, engine, authorizer)

	require.False(t, bool(result))
	require.NotNil(t, err) // TODO refine error
}

func TestDeleteRelationship_AuthorizedActorRemovesRelationship(t *testing.T) {
	ctx, engine, authorizer, pol := setupTestDeleteRelationship(t)

	cmd := DeleteRelationshipCommand{
		Policy:       pol,
		Relationship: types.NewActorRelationship("file", "readme.txt", "admin", "bob"),
		Actor:        "alice",
	}
	result, err := cmd.Execute(ctx, engine, authorizer)

	require.True(t, bool(result))
	require.Nil(t, err)
}

// setup for DeleteRelationship tests
// sets alice as the owner of readme.txt and sets bob as an admin of readme.txt
func setupUnregisterObjectTests(t *testing.T) (context.Context, auth_engine.AuthEngine, *RelationshipAuthorizer, *types.Policy) {
	polStr := `
    id: unregister-pol
    resources:
      file:
        permissions:
        relations:
          owner:
            types:
              - actor
          reader:
            types:
              - actor
              - file->reader
          admin:
            manages:
              - reader
            types:
              - actor

    actor:
      name: actor
    `
	polIR, err := policy.Unmarshal(polStr, types.PolicyMarshalingType_SHORT_YAML)
	require.Nil(t, err)

	engine, _, ctx := testutil.GetTestAuthEngine(t)

	createCmd := policy.CreatePolicyCommand{
		Policy:       polIR,
		Creator:      "cosmos1gue5de6a8fdff0jut08vw5sg9pk6rr00cstakj",
		CreationTime: timestamp,
	}
	pol, err := createCmd.Execute(ctx, &testutil.AccountKeeperStub{}, engine)
	require.Nil(t, err)

	_, err = engine.SetRelationship(ctx, pol, &types.RelationshipRecord{
		Relationship: types.NewActorRelationship("file", "readme.txt", "owner", "alice"),
		Archived:     false,
		Creator:      "alice",
	})
	require.Nil(t, err)

	_, err = engine.SetRelationship(ctx, pol, &types.RelationshipRecord{
		Relationship: types.NewActorRelationship("file", "readme.txt", "admin", "bob"),
		Archived:     false,
		Creator:      "bob",
	})
	require.Nil(t, err)

	_, err = engine.SetRelationship(ctx, pol, &types.RelationshipRecord{
		Relationship: types.NewActorSetRelationship("file", "delagated", "reader", "file", "readme.txt", "reader"),
		Archived:     false,
		Creator:      "charlie",
	})
	require.Nil(t, err)

	authorizer := NewRelationshipAuthorizer(engine)
	return ctx, engine, authorizer, pol
}

func TestUnregisterObject_RegisteredObjectCanBeUnregisteredByAuthor(t *testing.T) {
	ctx, engine, authorizer, policy := setupUnregisterObjectTests(t)

	command := UnregisterObjectCommand{
		Policy: policy,
		Object: &types.Object{"file", "readme.txt"},
		Actor:  "alice",
	}
	count, err := command.Execute(ctx, engine, authorizer)

	require.Nil(t, err)
	require.Equal(t, count, uint(2))
}

func TestUnregisterObject_ActorCannotUnregisterObjectTheyDoNotOwn(t *testing.T) {
	ctx, engine, authorizer, policy := setupUnregisterObjectTests(t)

	command := UnregisterObjectCommand{
		Policy: policy,
		Object: &types.Object{"file", "readme.txt"},
		Actor:  "bob",
	}
	count, err := command.Execute(ctx, engine, authorizer)

	require.ErrorIs(t, err, types.ErrNotAuthorized)
	require.Equal(t, count, uint(0))
}

func TestUnregisterObject_UnregisteringAnObjectThatDoesNotExistReturnsUnauthorized(t *testing.T) {
	ctx, engine, authorizer, policy := setupUnregisterObjectTests(t)

	command := UnregisterObjectCommand{
		Policy: policy,
		Object: &types.Object{"file", "nothing.txt"},
		Actor:  "bob",
	}
	count, err := command.Execute(ctx, engine, authorizer)

	require.ErrorIs(t, err, types.ErrNotAuthorized)
	require.Equal(t, count, uint(0))
}

func TestUnregisterObject_UnregisteringAnAlreadyArchivedObjectDoesNothing(t *testing.T) {
	ctx, engine, authorizer, policy := setupUnregisterObjectTests(t)

	command := UnregisterObjectCommand{
		Policy: policy,
		Object: &types.Object{"file", "nothing.txt"},
		Actor:  "bob",
	}
	count, err := command.Execute(ctx, engine, authorizer)

	require.ErrorIs(t, err, types.ErrNotAuthorized)
	require.Equal(t, count, uint(0))
}
