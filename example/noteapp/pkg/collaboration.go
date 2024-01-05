package pkg

import (
	"context"
	"fmt"

	"github.com/sourcenetwork/sourcehub/x/acp/types"
)

// NOTE this process of mapping the relationships seems kinda tedious and Error Prone.
// It would be really convenient to generate Go consts or something from a Policy definition.
// The idea is to give developers auto complete and move runtime errors to compilation errors

type PermissionManager struct {
	acp *ACPClient
}

func NewPermissionManager(acp *ACPClient) *PermissionManager {
	return &PermissionManager{
		acp: acp,
	}
}

// RegisterNote registers the session user as the owner of the note object
func (m *PermissionManager) RegisterNote(ctx context.Context, session Session, note *Note) (*Promise[Executed], error) {
	id := fmt.Sprintf("%v", note.ID)
	obj := types.NewObject("notes", id)
	msg := types.NewMsgRegisterObjectNow(session.Actor, PolicyId, obj)
	return m.acp.TxRegisterObject(ctx, session, msg)
}

// ShareNote makes the given actor a `reader` of the given Note
// Creates a Relationship in SourceHub
func (m *PermissionManager) ShareNote(ctx context.Context, session Session, noteId string, actorId string) (*Promise[Executed], error) {
	rel := types.NewActorRelationship("notes", noteId, "reader", actorId)
	msg := types.NewMsgSetRelationshipNow(session.Actor, PolicyId, rel)
	return m.acp.TxSetRelationship(ctx, session, msg)
}

// UnshareNote removes Actor as a reader of Note
// Creates a Relationship in SourceHub
func (m *PermissionManager) UnshareNote(ctx context.Context, session Session, noteId string, actorId string) (*Promise[Executed], error) {
	rel := types.NewActorRelationship("notes", noteId, "reader", actorId)
	msg := types.NewMsgSetRelationshipNow(session.Actor, PolicyId, rel)
	return m.acp.TxSetRelationship(ctx, session, msg)
}
