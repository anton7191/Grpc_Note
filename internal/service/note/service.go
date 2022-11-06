package note

import (
	"github.com/anton7191/note-server-api/internal/repository/note"
)

type Service struct {
	noteRepository note.Repository
}

func NewService(noteRepository note.Repository) *Service {
	return &Service{
		noteRepository: noteRepository,
	}
}
