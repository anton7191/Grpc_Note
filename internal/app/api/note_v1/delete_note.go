package note_v1

import (
	"context"

	desc "github.com/anton7191/note-server-api/pkg/note_v1"
)

func (i *Implementation) DeleteNote(ctx context.Context, req *desc.DeleteNoteRequest) (*desc.Empty, error) {
	res, err := i.noteService.DeleteNote(ctx, req)
	if err != nil {
		return nil, err
	}

	return res, nil
}
