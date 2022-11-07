package note_v1

import (
	"context"

	"github.com/anton7191/note-server-api/internal/converter"
	desc "github.com/anton7191/note-server-api/pkg/note_v1"
)

func (i *Implementation) GetNote(ctx context.Context, req *desc.GetNoteRequest) (*desc.GetNoteResponse, error) {
	res, err := i.noteService.GetNote(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	return &desc.GetNoteResponse{
		Note: converter.ToDescNote(res),
	}, nil
}
