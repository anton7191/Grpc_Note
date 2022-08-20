package note_v1

import (
	"context"
	"fmt"
	desc "github.com/anton7191/Grpc_Note/pkg/note_v1"
)

func (n *Note) GetListNote(ctx context.Context, req *desc.GetListNoteRequest) (*desc.GetListNoteResponse, error) {
	fmt.Println("GetListNote")
	fmt.Println("Request: ", req.Req)
	note_list := desc.GetListNoteResponse{
		Note: []*desc.Note{
			{Id: 1, Title: "1", Text: "first", Author: "Anton"},
			{Id: 2, Title: "2", Text: "second", Author: "Anton"},
		},
	}
	return &note_list, nil
}
