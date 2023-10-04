package relationship

import (
	"context"
	"testing"
        "errors"

	"github.com/stretchr/testify/require"
        sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cometbft/cometbft/libs/log"
	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"

	"github.com/sourcenetwork/sourcehub/testutil"
	authengineutil "github.com/sourcenetwork/sourcehub/testutil/auth_engine"
	"github.com/sourcenetwork/sourcehub/x/acp/auth_engine"
	"github.com/sourcenetwork/sourcehub/x/acp/types"
	"github.com/sourcenetwork/sourcehub/x/acp/policy"
	acptestutil "github.com/sourcenetwork/sourcehub/x/acp/testutil"
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
	engine, store := authengineutil.GetTestAuthEngine(t)
	ctx := sdk.NewContext(store, tmproto.Header{}, false, log.NewNopLogger())

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


// setu for SetRelationship tests
// sets alice as the owner of readme.txt and sets bob as an admin of readme.txt
func setupTestSetRelationship(t *testing.T) (context.Context, auth_engine.AuthEngine, *types.Policy) {
    polStr :=  `
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

    engine, store := authengineutil.GetTestAuthEngine(t)
    ctx := sdk.NewContext(store, tmproto.Header{}, false, log.NewNopLogger())

    createCmd := policy.CreatePolicyCommand{
        Policy: polIR,
        CreatorAddr: sdk.AccAddress([]byte("cosmos1gue5de6a8fdff0jut08vw5sg9pk6rr00cstakj")),
        CreationTime: timestamp,
    }
    pol, err := createCmd.Execute(ctx, &acptestutil.AccountKeeperStub{}, engine)
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

    return ctx, engine, pol
}

func TestSetRelationship_ValidRelationshipIsCreated(t *testing.T) {
	ctx, engine, policy := setupTestSetRelationship(t)

        command := SetRelationshipCommand{
            Policy: policy,
            CreationTs: timestamp,
            Creator: "bob",
            Relationship: types.NewActorRelationship("file", "readme.txt", "reader", "charlie"),
        }
        result, err := command.Execute(ctx, engine)

        require.Nil(t, err)
        require.Equal(t, result, types.SetRelationshipResult_SetRelCreated)
}

func  TestSetRelationship_CannotCreateRelationshipWithOwnerRelation(t *testing.T) {
	ctx, engine, policy := setupTestSetRelationship(t)

        command := SetRelationshipCommand{
            Policy: policy,
            CreationTs: timestamp,
            Creator: "bob",
            Relationship: types.NewActorRelationship("file", "any.txt", "owner", "bob"),
        }
        result, err := command.Execute(ctx, engine)

        require.True(t, errors.Is(err, ErrCannotSetOwnerRelationship))
        require.Equal(t, result, types.SetRelationshipResult_SetRelNoOp)
}

func TestSetRelationship_UnauthorizedActorDoesNotCreateRelationship(t *testing.T) {
	ctx, engine, policy := setupTestSetRelationship(t)

        command := SetRelationshipCommand{
            Policy: policy,
            CreationTs: timestamp,
            Creator: "eve",
            Relationship: types.NewActorRelationship("file", "readme.txt", "reader", "charlie"),
        }
        result, err := command.Execute(ctx, engine)

        require.Nil(t, err)
        require.Equal(t, result, types.SetRelationshipResult_SetRelDenied)
}

func  TestSetRelationship_SettingAnExistingRelationshipIsNoop(t *testing.T) {
	ctx, engine, policy := setupTestSetRelationship(t)

        command := SetRelationshipCommand{
            Policy: policy,
            CreationTs: timestamp,
            Creator: "bob",
            Relationship: types.NewActorRelationship("file", "readme.txt", "reader", "charlie"),
        }
        result, err := command.Execute(ctx, engine)
        require.Nil(t, err)
        require.Equal(t, result, types.SetRelationshipResult_SetRelCreated)

        // repeating the same action is a noop
        result, err = command.Execute(ctx, engine)
        require.Nil(t, err)
        require.Equal(t, result, types.SetRelationshipResult_SetRelNoOp)
}

func  TestSetRelationship_CannotCreateRelationshipForUndefinedObject(t *testing.T) {
	ctx, engine, policy := setupTestSetRelationship(t)

        command := SetRelationshipCommand{
            Policy: policy,
            CreationTs: timestamp,
            Creator: "bob",
            Relationship: types.NewActorRelationship("file", "askdfjas", "reader", "charlie"),
        }
        result, err := command.Execute(ctx, engine)
        require.Nil(t, err)
        require.Equal(t, result, types.SetRelationshipResult_SetRelDenied)
}
