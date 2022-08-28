package note_v1

import (
	"context"
	"fmt"
	desc "github.com/anton7191/Note-server-api/pkg/note_v1"
)

func (n *Note) GetListNote(ctx context.Context, req *desc.Empty) (*desc.GetListNoteResponse, error) {
	fmt.Println("GetListNote")
	fmt.Println("Request: ", "Empty")
	noteList := desc.GetListNoteResponse{
		Note: []*desc.Note{
			{Id: 1, Title: "1", Text: "first", Author: "Anton"},
			{Id: 2, Title: "2", Text: "second", Author: "Anton"},
		},
	}
	return &noteList, nil
}
