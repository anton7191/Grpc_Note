package note_v1

import (
	"context"
	"fmt"

	desc "github.com/anton7191/Note-server-api/pkg/note_v1"
)

func (n *Note) UpdateNote(ctx context.Context, req *desc.UpdateNoteRequest) (*desc.Empty, error) {
	fmt.Println("Update Note")
	fmt.Println("ID: ", req.Note.GetId())
	fmt.Println("text: ", req.Note.GetText())
	fmt.Println("Title: ", req.Note.GetTitle())
	fmt.Println("Author: ", req.Note.GetAuthor())

	return &desc.Empty{}, nil
}
