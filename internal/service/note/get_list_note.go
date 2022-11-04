package note

import (
	"context"

	"github.com/anton7191/note-server-api/internal/model"
)

func (s *Service) GetListNote(ctx context.Context) ([]*model.Note, error) {
	res, err := s.noteRepository.GetListNote(ctx)
	if err != nil {
		return nil, err
	}

	return res, nil
}
