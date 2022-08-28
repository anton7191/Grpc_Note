package note_v1

import (
	"context"
	"fmt"

	desc "github.com/anton7191/Note-server-api/pkg/note_v1"
)

func (n *Note) DeleteNote(ctx context.Context, req *desc.DeleteNoteRequest) (*desc.Empty, error) {
	fmt.Println("Delete Note")
	fmt.Println("ID: ", req.GetId())

	return &desc.Empty{}, nil
}
