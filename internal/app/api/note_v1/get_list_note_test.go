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
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestListNote(t *testing.T) {
	var (
		ctx      = context.Background()
		mockCtrl = gomock.NewController(t)

		repoErrText = gofakeit.Phrase()

		repoErr = errors.New(repoErrText)
	)

	repoRes := []*model.Note{}
	validResSlice := []*desc.Note{}
	var (
		id                  int64
		title, text, author string
		createdAt           time.Time
		updatedAt           sql.NullTime
	)

	for i := 0; i < 5; i++ {
		id = gofakeit.Int64()
		title = gofakeit.BeerName()
		text = gofakeit.BeerName()
		author = gofakeit.BeerName()
		createdAt = time.Now()
		updatedAt = sql.NullTime{Time: time.Now(), Valid: true}

		repoRes = append(repoRes, &model.Note{
			ID:        id,
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
			Info: &model.NoteInfo{
				Title:  title,
				Text:   text,
				Author: author,
			},
		})

		validResSlice = append(validResSlice, &desc.Note{
			Id:        id,
			CreatedAt: timestamppb.New(createdAt),
			UpdatedAt: timestamppb.New(updatedAt.Time),
			Note: &desc.NoteInfo{
				Title:  title,
				Text:   text,
				Author: author,
			},
		})
	}

	validRes := &desc.GetListNoteResponse{
		Note: validResSlice,
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
