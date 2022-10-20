package note_v1

import (
	"context"

	desc "github.com/anton7191/note-server-api/pkg/note_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) DeleteNote(ctx context.Context, req *desc.DeleteNoteRequest) (*emptypb.Empty, error) {
	err := i.noteService.DeleteNote(ctx, req)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
