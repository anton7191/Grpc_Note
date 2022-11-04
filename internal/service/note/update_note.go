package note

import (
	"context"

	"github.com/anton7191/note-server-api/internal/model"
)

func (s *Service) UpdateNote(ctx context.Context, id int64, updateNoteInfo *model.UpdateNoteInfo) error {
	return s.noteRepository.UpdateNote(ctx, id, updateNoteInfo)
}
