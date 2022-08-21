package note_v1

import (
	"context"
	"fmt"
	desc "github.com/anton7191/Note-server-api/pkg/note_v1"
)

func (n *Note) GetNote(ctx context.Context, req *desc.GetNoteRequest) (*desc.GetNoteResponse, error) {
	fmt.Println("GetNote")
	fmt.Println("Id: ", req.Id)

	return &desc.GetNoteResponse{
		Author: "Anton",
		Text:   "Get Note!",
		Title:  "First msg",
	}, nil
}
