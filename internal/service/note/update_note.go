package note

import (
	"context"

	desc "github.com/anton7191/note-server-api/pkg/note_v1"
)

func (s *Service) UpdateNote(ctx context.Context, req *desc.UpdateNoteRequest) (*desc.Empty, error) {
	err := s.noteRepository.UpdateNote(ctx, req)
	if err != nil {
		return nil, err
	}

	return &desc.Empty{}, nil
}
