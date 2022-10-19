package note

import (
	"context"

	desc "github.com/anton7191/note-server-api/pkg/note_v1"
)

func (s *Service) GetListNote(ctx context.Context, req *desc.Empty) (*desc.GetListNoteResponse, error) {
	res, err := s.noteRepository.GetListNote(ctx, req)
	if err != nil {
		return nil, err
	}

	return &desc.GetListNoteResponse{
		Note: res,
	}, nil
}
