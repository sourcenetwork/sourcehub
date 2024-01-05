package pkg

import (
	"context"

	prototypes "github.com/cosmos/gogoproto/types"

	"github.com/sourcenetwork/sourcehub/utils"
	"github.com/sourcenetwork/sourcehub/x/acp/access_ticket"
	"github.com/sourcenetwork/sourcehub/x/acp/types"
)

func GenReadablesTicket(ctx context.Context, session Session, acp *ACPClient, ids []string) (string, error) {
	operations := utils.MapSlice(ids, func(id string) *types.Operation {
		return &types.Operation{
			Object:     types.NewObject("notes", id),
			Permission: "read",
		}
	})

	msg := &types.MsgCheckAccess{
		Creator:      session.Actor,
		PolicyId:     PolicyId,
		CreationTime: prototypes.TimestampNow(),
		AccessRequest: &types.AccessRequest{
			Operations: operations,
			Actor: &types.Actor{
				Id: session.Actor,
			},
		},
	}

	promise, err := acp.TxCheckAccess(ctx, session, msg)
	if err != nil {
		return "", err
	}

	checkId, err := promise.Await()
	if err != nil {
		return "", err
	}

	abciServ, err := access_ticket.NewABCIService("tcp://127.0.0.1:26657")
	if err != nil {
		return "", err
	}

	issuer := access_ticket.NewTicketIssuer(&abciServ)
	ticket, err := issuer.Issue(ctx, checkId, session.PrivKey)
	if err != nil {
		return "", err
	}

	return ticket, nil
}
