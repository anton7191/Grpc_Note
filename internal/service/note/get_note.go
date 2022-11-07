package note

import (
	"context"

	"github.com/anton7191/note-server-api/internal/model"
)

func (s *Service) GetNote(ctx context.Context, id int64) (*model.Note, error) {
	return s.noteRepository.GetNote(ctx, id)
}
