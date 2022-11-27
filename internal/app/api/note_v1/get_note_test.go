package note_v1

import (
	"context"
	"database/sql"
	"errors"
	"testing"
	"time"

	"github.com/anton7191/note-server-api/internal/model"
	noteMocks "github.com/anton7191/note-server-api/internal/repository/note/mocks"
	"github.com/anton7191/note-server-api/internal/service/note"
	desc "github.com/anton7191/note-server-api/pkg/note_v1"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestGetNote(t *testing.T) {
	var (
		ctx      = context.Background()
		mockCtrl = gomock.NewController(t)

		id          = gofakeit.Int64()
		title       = gofakeit.BeerName()
		text        = gofakeit.BeerStyle()
		author      = gofakeit.Name()
		repoErrText = gofakeit.Phrase()

		req = &desc.GetNoteRequest{
			Id: id,
		}

		createdAt = time.Now()
		updatedAt = sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		}

		repoRes = &model.Note{
			ID:        id,
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
			Info: &model.NoteInfo{
				Title:  title,
				Text:   text,
				Author: author,
			},
		}

		validRes = &desc.GetNoteResponse{
			Note: &desc.Note{
				Id:        id,
				CreatedAt: timestamppb.New(createdAt),
				UpdatedAt: timestamppb.New(updatedAt.Time),
				Note: &desc.NoteInfo{
					Title:  title,
					Text:   text,
					Author: author,
				},
			},
		}

		repoErr = errors.New(repoErrText)
	)

	noteMock := noteMocks.NewMockRepository(mockCtrl)
	gomock.InOrder(
		noteMock.EXPECT().GetNote(ctx, id).Return(repoRes, nil),
		noteMock.EXPECT().GetNote(ctx, id).Return(repoRes, repoErr),
	)

	api := NewMockNoteV1(Implementation{
		noteService: note.NewMockNoteService(noteMock),
	})

	t.Run("get note success", func(t *testing.T) {
		res, err := api.GetNote(ctx, req)
		require.Nil(t, err)
		require.Equal(t, validRes, res)
	})

	t.Run("get note repo err", func(t *testing.T) {
		_, err := api.GetNote(ctx, req)
		require.NotNil(t, err)
		require.Equal(t, repoErrText, err.Error())
	})
}
