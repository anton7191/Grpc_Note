package note

import (
	"context"
)

func (s *Service) DeleteNote(ctx context.Context, id int64) error {
	return s.noteRepository.DeleteNote(ctx, id)
}
