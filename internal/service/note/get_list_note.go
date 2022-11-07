package note

import (
	"context"

	"github.com/anton7191/note-server-api/internal/model"
)

func (s *Service) GetListNote(ctx context.Context) ([]*model.Note, error) {
	return s.noteRepository.GetListNote(ctx)
}
