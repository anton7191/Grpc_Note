package note

import (
	"context"

	desc "github.com/anton7191/note-server-api/pkg/note_v1"
)

func (s *Service) GetNote(ctx context.Context, req *desc.GetNoteRequest) (*desc.GetNoteResponse, error) {
	note, err := s.noteRepository.GetNote(ctx, req)
	if err != nil {
		return nil, err
	}
	return &desc.GetNoteResponse{
		Note: note,
	}, nil

}
