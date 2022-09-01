package note_v1

import desc "github.com/anton7191/note-server-api/pkg/note_v1"

type Implementation struct {
	desc.UnimplementedNoteV1Server
}

func NewNote() *Implementation {
	return &Implementation{}
}
