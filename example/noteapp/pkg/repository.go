// Repository models a Faux Defra which returns documents based on the result of
// evaluating an AccessTicket

package pkg

import (
	"context"
	"fmt"
	"log"
	"slices"
	"strconv"

	"github.com/sourcenetwork/sourcehub/x/acp/access_ticket"
	"github.com/sourcenetwork/sourcehub/x/acp/did"
	"github.com/sourcenetwork/sourcehub/x/acp/types"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var _ Repository = (*SQLite)(nil)

type Repository interface {
	SetNote(ctx context.Context, session Session, note *Note) error

	GetNote(ctx context.Context, session Session, ticket string, id string) (Note, error)

	ListLocalNotes(ctx context.Context, actorId string) ([]Note, error)
}

func NewSQLite(path string) (*SQLite, error) {
	db, err := gorm.Open(sqlite.Open(path), &gorm.Config{})

	if err != nil {
		return nil, nil
	}

	repo := &SQLite{
		db: db,
	}
	err = repo.Init()
	if err != nil {
		return nil, err
	}

	return repo, nil
}

type SQLite struct {
	db *gorm.DB
}

func (s *SQLite) SetNote(ctx context.Context, session Session, note *Note) error {
	result := s.db.Create(note)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *SQLite) GetNote(ctx context.Context, session Session, ticket string, id string) (Note, error) {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return Note{}, err
	}

	note := Note{}
	result := s.db.First(&note, uint(idInt))
	if result.Error != nil {
		return Note{}, result.Error
	}

	if note.Creator == session.Actor {
		return note, nil
	}

	ops, err := s.getTicketOperations(ticket)
	if err != nil {
		return Note{}, err
	}

	op := types.Operation{
		Object:     types.NewObject("notes", fmt.Sprintf("%v", id)),
		Permission: "read",
	}

	if !slices.ContainsFunc(ops, func(o *types.Operation) bool {
		return o.Object.Id == op.Object.Id && o.Object.Resource == op.Object.Resource && o.Permission == op.Permission
	}) {
		return Note{}, fmt.Errorf("unauthorized: actor %v cannot read note %v", session.Actor, id)
	}

	return note, nil
}

func (s *SQLite) Init() error {
	return s.db.AutoMigrate(&Note{})
}

func (s *SQLite) getTicketOperations(ticket string) ([]*types.Operation, error) {
	ctx := context.Background()
	service, err := access_ticket.NewABCIService("tcp://127.0.0.1:26657")
	if err != nil {
		return nil, err
	}

	registry := did.NewKeyRegistry()
	spec := access_ticket.NewAccessTicketSpec(&service, registry)

	err = spec.Satisfies(ctx, ticket)
	if err != nil {
		return nil, err
	}

	log.Printf("Ticket valid!")

	marshaler := access_ticket.Marshaler{}
	decision, _ := marshaler.Unmarshal(ticket)
	return decision.Decision.Operations, nil
}

func (s *SQLite) ListLocalNotes(ctx context.Context, actorId string) ([]Note, error) {
	var notes []Note
	result := s.db.Where("creator = ?", actorId).Find(&notes)

	if result.Error != nil {
		return nil, result.Error
	}

	return notes, nil
}
