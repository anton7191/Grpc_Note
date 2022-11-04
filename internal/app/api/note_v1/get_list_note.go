package note_v1

import (
	"context"

	"github.com/anton7191/note-server-api/internal/converter"
	desc "github.com/anton7191/note-server-api/pkg/note_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) GetListNote(ctx context.Context, req *emptypb.Empty) (*desc.GetListNoteResponse, error) {
	res, err := i.noteService.GetListNote(ctx)
	if err != nil {
		return nil, err
	}

	return &desc.GetListNoteResponse{
		Note: converter.ToDescNoteSlice(res),
	}, nil
}
