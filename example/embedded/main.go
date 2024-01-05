package main

import (
	"log"

	prototypes "github.com/cosmos/gogoproto/types"

	"github.com/sourcenetwork/sourcehub/testutil/sample"
	"github.com/sourcenetwork/sourcehub/x/acp/embedded"
	"github.com/sourcenetwork/sourcehub/x/acp/types"
)

const policy string = `
description: a test policy which mocks a file system with files as resources

resources:
  file:
    permissions:
      read:
        expr: owner + reader
      write:
        expr: owner

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

func main() {
	acp, err := embedded.NewLocalACP()
	if err != nil {
		log.Fatal(err)
	}

	// generate 3 account addresses (from random key pairs) for the 3 actors in the system
	alice, bob, creator := sample.AccAddress(), sample.AccAddress(), sample.AccAddress()
	log.Printf("alice: %v", alice)
	log.Printf("bob: %v", bob)
	log.Printf("creator: %v", creator)

	ctx := acp.GetCtx()
	msgService := acp.GetMsgService()
	queryService := acp.GetQueryService()
	_ = queryService

	polResp, err := msgService.CreatePolicy(ctx, &types.MsgCreatePolicy{
		Creator:      creator,
		Policy:       policy,
		MarshalType:  types.PolicyMarshalingType_SHORT_YAML,
		CreationTime: prototypes.TimestampNow(),
	})
	if err != nil {
		log.Fatalf("failed to create policy: %v", err)
	}
	log.Printf("policy created: %v", polResp.Policy.Id)

	regResp, err := msgService.RegisterObject(ctx, &types.MsgRegisterObject{
		Creator:      alice,
		PolicyId:     polResp.Policy.Id,
		Object:       types.NewObject("file", "readme.txt"),
		CreationTime: prototypes.TimestampNow(),
	})
	if err != nil {
		log.Fatalf("failed to register obj: %v", err)
	}
	log.Printf("alice registered file readme.txt: result %v", regResp.Result)

	_, err = msgService.SetRelationship(ctx, &types.MsgSetRelationship{
		Creator:      alice,
		PolicyId:     polResp.Policy.Id,
		Relationship: types.NewActorRelationship("file", "readme.txt", "reader", bob),
		CreationTime: prototypes.TimestampNow(),
	})
	if err != nil {
		log.Fatalf("failed to set relationship: %v", err)
	}
	log.Printf("alice set bob as reader of file readme.txt")

	checkResult, err := queryService.VerifyAccessRequest(ctx, &types.QueryVerifyAccessRequestRequest{
		PolicyId: polResp.Policy.Id,
		AccessRequest: &types.AccessRequest{
			Operations: []*types.Operation{
				{
					Object:     types.NewObject("file", "readme.txt"),
					Permission: "read",
				},
			},
			Actor: &types.Actor{
				Id: bob,
			},
		},
	})
	if err != nil {
		log.Fatalf("verify access request failed: %v", err)
	}
	log.Printf("is bob reader of readme.txt? %v", checkResult.Valid)
}
