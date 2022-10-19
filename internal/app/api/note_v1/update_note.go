package note_v1

import (
	"context"

	desc "github.com/anton7191/note-server-api/pkg/note_v1"
)

func (i *Implementation) UpdateNote(ctx context.Context, req *desc.UpdateNoteRequest) (*desc.Empty, error) {
	res, err := i.noteService.UpdateNote(ctx, req)
	if err != nil {
		return nil, err
	}

	return res, nil
}
