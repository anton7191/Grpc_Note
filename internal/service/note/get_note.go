package note

import (
	"context"

	"github.com/anton7191/note-server-api/internal/model"
)

func (s *Service) GetNote(ctx context.Context, id int64) (*model.Note, error) {
	note, err := s.noteRepository.GetNote(ctx, id)
	if err != nil {
		return nil, err
	}
	return note, nil

}
