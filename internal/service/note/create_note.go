package note

import (
	"context"

	"github.com/anton7191/note-server-api/internal/model"
)

func (s *Service) CreateNote(ctx context.Context, noteInfo *model.NoteInfo) (int64, error) {
	return s.noteRepository.CreateNote(ctx, noteInfo)
}
