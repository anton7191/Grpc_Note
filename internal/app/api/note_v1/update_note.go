package note_v1

import (
	"context"
	"fmt"
	desc "github.com/anton7191/Grpc_Note/pkg/note_v1"
)

func (n *Note) UpdateNote(ctx context.Context, req *desc.UpdateNoteRequest) (*desc.UpdateNoteResponse, error) {
	fmt.Println("Update Note")
	fmt.Println("ID: ", req.GetId())
	fmt.Println("text: ", req.GetText())

	return &desc.UpdateNoteResponse{
		Status: true,
	}, nil
}
