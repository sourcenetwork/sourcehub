package embedded

import (
	"os"
	"testing"

	"cosmossdk.io/log"
	prototypes "github.com/cosmos/gogoproto/types"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"

	"github.com/sourcenetwork/sourcehub/x/acp/types"
)

func TestLocalACPIsPersistentAccrossCalls(t *testing.T) {
	var storePath string
	var polId string

	// Given a previously created Embedded ACP store
	storePath = setupDirectory(t)
	t.Logf("store dir: %v", storePath)

	acp, err := NewLocalACP(
		WithPersistentStorage(storePath),
		WithLogger(log.NewTestLogger(t)),
	)
	require.Nil(t, err)

	srv := acp.GetMsgService()
	ctx := acp.GetCtx()
	createResp, err := srv.CreatePolicy(ctx, getCreatePolMsg())

	require.Nil(t, err)
	polId = createResp.Policy.Id
	err = acp.Close() // close first instance
	require.Nil(t, err)

	// create new ACP with same path
	acp, err = NewLocalACP(
		WithPersistentStorage(storePath),
		WithLogger(log.NewTestLogger(t)),
	)
	ctx = acp.GetCtx()
	require.Nil(t, err)

	// When I query the store for a stored Policy
	query := acp.GetQueryService()
	resp, err := query.Policy(ctx, &types.QueryPolicyRequest{
		Id: polId,
	})

	respIds, errIds := query.PolicyIds(ctx, &types.QueryPolicyIdsRequest{})

	// Then the Policy is fetched
	require.Nil(t, errIds)
	require.True(t, len(respIds.Ids) == 1)
	t.Logf("ids %v", respIds.Ids)

	require.Nil(t, err)
	require.NotNil(t, resp.Policy)

	// cleanup
	err = acp.Close()
	require.Nil(t, err)
	os.RemoveAll(storePath)
}

func setupDirectory(t *testing.T) string {
	path := "/tmp/" + uuid.New().String()
	err := os.Mkdir(path, os.ModePerm)
	require.Nil(t, err)

	return path
}

func getCreatePolMsg() *types.MsgCreatePolicy {
	policyStr := `
name: policy
description: ok
resources:
  file:
    relations: 
      owner:
        doc: owner owns
        types:
          - actor-resource
      reader:
      admin:
        manages:
          - reader
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

	var creator = "cosmos1gue5de6a8fdff0jut08vw5sg9pk6rr00cstakj"
	var timestamp = prototypes.TimestampNow()

	return &types.MsgCreatePolicy{
		Creator:      creator,
		Policy:       policyStr,
		MarshalType:  types.PolicyMarshalingType_SHORT_YAML,
		CreationTime: timestamp,
	}
}
