package note_v1

import (
	"context"
	"fmt"

	desc "github.com/anton7191/note-server-api/pkg/note_v1"
)

func (n *Implementation) GetNote(ctx context.Context, req *desc.GetNoteRequest) (*desc.GetNoteResponse, error) {
	fmt.Println("GetNote")
	fmt.Println("Id: ", req.GetId())

	return &desc.GetNoteResponse{
		Note: &desc.Note{
			Id:     1,
			Author: "Anton",
			Text:   "Get Note!",
			Title:  "First msg",
		},
	}, nil
}
