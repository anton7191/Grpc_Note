package note_v1

import (
	"context"

	desc "github.com/anton7191/note-server-api/pkg/note_v1"
	_ "github.com/jackc/pgx/stdlib"
)

func (i *Implementation) CreateNote(ctx context.Context, req *desc.CreateNoteRequest) (*desc.CreateNoteResponse, error) {
	res, err := i.noteService.CreateNote(ctx, req)
	if err != nil {
		return nil, err
	}

	return res, nil
}
