package note_v1

import (
	"context"

	"github.com/anton7191/note-server-api/internal/converter"
	desc "github.com/anton7191/note-server-api/pkg/note_v1"
	_ "github.com/jackc/pgx/stdlib"
)

func (i *Implementation) CreateNote(ctx context.Context, req *desc.CreateNoteRequest) (*desc.CreateNoteResponse, error) {
	id, err := i.noteService.CreateNote(ctx, converter.ToNoteInfo(req.GetNote()))
	if err != nil {
		return nil, err
	}

	return &desc.CreateNoteResponse{
		Id: id,
	}, nil
}
