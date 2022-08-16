package note_v1

import (
	"context"
	"fmt"
	desc "github.com/anton7191/Grpc_Note/pkg/note_v1"
)

func (n *Note) CreateNote(ctx context.Context, req *desc.CreateNoteRequest) (*desc.CreateNoteResponse, error) {
	fmt.Println("CreateNote")
	fmt.Println("title: ", req.GetTitle())
	fmt.Println("text: ", req.GetText())
	fmt.Println("author: ", req.GetAuthor())

	return &desc.CreateNoteResponse{
		Id: 1,
	}, nil
}

func (*Note) GetNote(ctx context.Context, req *desc.GetNoteRequest) (*desc.GetNoteResponse, error) {
	fmt.Println("GetNote")
	fmt.Println("Id: ", req.Id)

	return &desc.GetNoteResponse, nil
}
