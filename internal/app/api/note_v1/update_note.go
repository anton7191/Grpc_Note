package note_v1

import (
	"fmt"
	desc "github.com/anton7191/Grpc_Note/pkg/note_v1"
)

func (n *Note) UpdateNote(ctx context.Context, req *desc.Up) (*desc.CreateNoteResponse, error) {
	fmt.Println("CreateNote")
	fmt.Println("title: ", req.GetTitle())
	fmt.Println("text: ", req.GetText())
	fmt.Println("author: ", req.GetAuthor())

	return &desc.CreateNoteResponse{
		Id: 1,
	}, nil
}
