package pkg

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/sourcenetwork/sourcehub/utils"
	"github.com/sourcenetwork/sourcehub/x/acp/types"
)

type NoteCommands struct {
	repo    Repository
	manager *PermissionManager
}

func NewNoteCommands(repo Repository, manager *PermissionManager) NoteCommands {
	return NoteCommands{
		repo:    repo,
		manager: manager,
	}
}

// Create creates a new note with an unique Id
func (c *NoteCommands) Create(ctx context.Context, session Session, title string, body string) (Note, error) {
	note := Note{
		Title:      title,
		Body:       body,
		Creator:    session.Actor,
		LastEditor: session.Actor,
	}

	// TODO Wrap in Db Tx

	// How would I go about this if it were embedded in a collection
	err := c.repo.SetNote(ctx, session, &note)
	if err != nil {
		return Note{}, err
	}
	log.Printf("created note %v", note.ID)

	promise, err := c.manager.RegisterNote(ctx, session, &note)
	if err != nil {
		return Note{}, err
	}
	_, err = promise.Await()
	if err != nil {
		return Note{}, err
	}

	return note, nil
}

type NoteQuerier struct {
	acp   *ACPClient
	repo  Repository
	query types.QueryClient
}

func NewNoteQuerier(acp *ACPClient, repo Repository, querier types.QueryClient) NoteQuerier {
	return NoteQuerier{
		acp:   acp,
		repo:  repo,
		query: querier,
	}
}

// ListAccessibleNotes lists all notes the user has access to (as a reader or editor)
//
// This is kind of an awkward step because we need to query the ACP module to figure this out.
//
// In Defra there is a challenge because doc syncing is an app level concern.
// This creates a chicken and egg problem, because the app can't rely on defra to know which docs
// it has access to but defra can't sync the doc because it doesn't know which docs were shared.
// Ultimately this info is in the Relation graph, as such we could fetch this directly from it.
//
// Another alternative would require dApp devs to implement ad hoc and custom notification systems
// between the peers.
// This would lead to potentially more work for dApp devs and it's something the Relation Graph
// could solve
func (q *NoteQuerier) ListReadableNotesID(ctx context.Context, session Session) ([]string, error) {
	notes, err := q.repo.ListLocalNotes(ctx, session.Actor)
	if err != nil {
		return nil, err
	}
	localIds := utils.MapSlice(notes, func(n Note) string { return fmt.Sprintf("%v", n.ID) })

	// TODO Make builder for this nightmate
	selector := &types.RelationshipSelector{
		ObjectSelector: &types.ObjectSelector{
			Selector: &types.ObjectSelector_Wildcard{
				Wildcard: &types.WildcardSelector{},
			},
		},
		RelationSelector: &types.RelationSelector{
			Selector: &types.RelationSelector_Relation{Relation: "reader"},
		},
		SubjectSelector: &types.SubjectSelector{
			Selector: &types.SubjectSelector_Subject{
				Subject: &types.Subject{
					Subject: &types.Subject_Actor{
						Actor: &types.Actor{
							Id: session.Actor,
						},
					},
				},
			},
		},
	}
	msg := types.QueryFilterRelationshipsRequest{
		PolicyId: PolicyId,
		Selector: selector,
	}
	resp, err := q.query.FilterRelationships(ctx, &msg)
	if err != nil {
		return nil, err
	}

	remoteIds := utils.MapSlice(resp.Records, func(rec *types.RelationshipRecord) string {
		return rec.Relationship.Object.Id
	})

	return append(remoteIds, localIds...), nil
}

// Note: this assumes the need for internet, which is lame
// Local users should be able to check their own notes and sync the remaining notes as it goes.
// Right now that's kinda lame to do
func (q *NoteQuerier) FetchReadableNotes(ctx context.Context, session Session) ([]Note, error) {
	ids, err := q.ListReadableNotesID(ctx, session)
	if err != nil {
		return nil, err
	}

	if len(ids) == 0 {
		return nil, nil
	}

	ticket, err := GenReadablesTicket(ctx, session, q.acp, ids)
	if err != nil {
		return nil, err
	}

	log.Printf("got ticket %v", ticket)

	log.Printf("waiting until block is commited")
	time.Sleep(4000 * time.Millisecond)

	notes := make([]Note, 0, len(ids))

	for _, id := range ids {
		note, err := q.repo.GetNote(ctx, session, ticket, id)
		if err != nil {
			return nil, err
		}
		notes = append(notes, note)
	}

	return notes, nil
}

func (q *NoteQuerier) FetchLocalNodes(ctx context.Context, session Session) ([]Note, error) {
	return q.repo.ListLocalNotes(ctx, session.Actor)
}
