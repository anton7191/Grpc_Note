package note_v1

import (
	"context"
	"database/sql"
	"errors"
	"testing"

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

func TestUpdateNote(t *testing.T) {
	var (
		ctx      = context.Background()
		mockCtrl = gomock.NewController(t)

		id          = gofakeit.Int64()
		repoErrText = gofakeit.Phrase()

		repoReq = &model.UpdateNoteInfo{
			Title:  sql.NullString{},
			Text:   sql.NullString{},
			Author: sql.NullString{},
		}

		req = &desc.UpdateNoteRequest{
			Id:   id,
			Note: converter.ToDescUpdateNoteInfo(repoReq),
		}

		repoErr = errors.New(repoErrText)
	)

	noteMock := noteMocks.NewMockRepository(mockCtrl)
	gomock.InOrder(
		noteMock.EXPECT().UpdateNote(ctx, id, repoReq).Return(nil),
		noteMock.EXPECT().UpdateNote(ctx, id, repoReq).Return(repoErr),
	)

	api := NewMockNoteV1(Implementation{
		noteService: note.NewMockNoteService(noteMock),
	})

	t.Run("update note success", func(t *testing.T) {
		res, err := api.UpdateNote(ctx, req)
		require.Nil(t, err)
		require.Equal(t, &emptypb.Empty{}, res)
	})

	t.Run("update note repo err", func(t *testing.T) {
		_, err := api.UpdateNote(ctx, req)
		require.NotNil(t, err)
		require.Equal(t, repoErrText, err.Error())
	})
}