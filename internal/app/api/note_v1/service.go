package note_v1

import (
	"github.com/anton7191/note-server-api/internal/service/note"
	desc "github.com/anton7191/note-server-api/pkg/note_v1"
)

type Implementation struct {
	desc.UnimplementedNoteV1Server

	noteService *note.Service
}

func NewNote(noteService *note.Service) *Implementation {
	return &Implementation{
		noteService: noteService,
	}
}

func NewMockNoteV1(i Implementation) *Implementation {
	return &Implementation{
		desc.UnimplementedNoteV1Server{},
		i.noteService,
	}
}
