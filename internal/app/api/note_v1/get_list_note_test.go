package note_v1

import (
	"context"
	"database/sql"
	"errors"
	"testing"
	"time"

	"github.com/anton7191/note-server-api/internal/converter"
	"github.com/anton7191/note-server-api/internal/model"
	noteMocks "github.com/anton7191/note-server-api/internal/repository/note/mocks"
	"github.com/anton7191/note-server-api/internal/service/note"
	desc "github.com/anton7191/note-server-api/pkg/note_v1"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/emptypb"
)

func TestListNote(t *testing.T) {
	var (
		ctx      = context.Background()
		mockCtrl = gomock.NewController(t)

		repoErrText = gofakeit.Phrase()

		repoErr = errors.New(repoErrText)
	)

	repoRes := []*model.Note{}

	for i := 0; i < 5; i++ {
		repoRes = append(repoRes, &model.Note{
			ID:        gofakeit.Int64(),
			CreatedAt: time.Now(),
			UpdatedAt: sql.NullTime{},
			Info: &model.NoteInfo{
				Title:  gofakeit.BeerName(),
				Text:   gofakeit.BeerName(),
				Author: gofakeit.BeerName(),
			},
		})
	}

	validRes := &desc.GetListNoteResponse{
		Note: converter.ToDescNoteSlice(repoRes),
	}

	noteMock := noteMocks.NewMockRepository(mockCtrl)
	gomock.InOrder(
		noteMock.EXPECT().GetListNote(ctx).Return(repoRes, nil),
		noteMock.EXPECT().GetListNote(ctx).Return(repoRes, repoErr),
	)

	api := NewMockNoteV1(Implementation{
		noteService: note.NewMockNoteService(noteMock),
	})

	t.Run("get list note success", func(t *testing.T) {
		res, err := api.GetListNote(ctx, &emptypb.Empty{})
		require.Nil(t, err)
		require.Equal(t, validRes, res)
	})

	t.Run("get list note repo err", func(t *testing.T) {
		_, err := api.GetListNote(ctx, &emptypb.Empty{})
		require.NotNil(t, err)
		require.Equal(t, repoErrText, err.Error())
	})
}
