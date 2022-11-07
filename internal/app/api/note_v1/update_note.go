package note_v1

import (
	"context"

	"github.com/anton7191/note-server-api/internal/converter"
	desc "github.com/anton7191/note-server-api/pkg/note_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) UpdateNote(ctx context.Context, req *desc.UpdateNoteRequest) (*emptypb.Empty, error) {
	err := i.noteService.UpdateNote(ctx, req.GetId(), converter.ToUpdateNoteInfo(req.GetNote()))
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
