package note_v1

import (
	"context"
	"errors"
	"testing"

	"github.com/anton7191/note-server-api/internal/model"
	noteMocks "github.com/anton7191/note-server-api/internal/repository/note/mocks"
	"github.com/anton7191/note-server-api/internal/service/note"
	desc "github.com/anton7191/note-server-api/pkg/note_v1"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestCreateNote(t *testing.T) {
	var (
		ctx      = context.Background()
		mockCtrl = gomock.NewController(t)

		id          = gofakeit.Int64()
		title       = gofakeit.BeerName()
		text        = gofakeit.BeerStyle()
		author      = gofakeit.Name()
		repoErrText = gofakeit.Phrase()

		req = &desc.CreateNoteRequest{
			Note: &desc.NoteInfo{
				Title:  title,
				Text:   text,
				Author: author,
			},
		}

		repoReq = &model.NoteInfo{
			Title:  title,
			Text:   text,
			Author: author,
		}

		validRes = &desc.CreateNoteResponse{
			Id: id,
		}

		repoErr = errors.New(repoErrText)
	)

	noteMock := noteMocks.NewMockRepository(mockCtrl)
	gomock.InOrder(
		noteMock.EXPECT().CreateNote(ctx, repoReq).Return(id, nil),
		noteMock.EXPECT().CreateNote(ctx, repoReq).Return(int64(0), repoErr),
	)

	api := NewMockNoteV1(Implementation{
		noteService: note.NewMockNoteService(noteMock),
	})

	t.Run("create note success", func(t *testing.T) {
		res, err := api.CreateNote(ctx, req)
		require.Nil(t, err)
		require.Equal(t, validRes, res)
	})

	t.Run("note repo err", func(t *testing.T) {
		_, err := api.CreateNote(ctx, req)
		require.NotNil(t, err)
		require.Equal(t, repoErrText, err.Error())
	})
}
