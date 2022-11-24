package note_v1

import (
	"context"
	"errors"
	"testing"

	noteMocks "github.com/anton7191/note-server-api/internal/repository/note/mocks"
	"github.com/anton7191/note-server-api/internal/service/note"
	desc "github.com/anton7191/note-server-api/pkg/note_v1"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/emptypb"
)

func TestDeleteNote(t *testing.T) {
	var (
		ctx      = context.Background()
		mockCtrl = gomock.NewController(t)

		id          = gofakeit.Int64()
		repoErrText = gofakeit.Phrase()

		req = &desc.DeleteNoteRequest{
			Id: id,
		}

		repoErr = errors.New(repoErrText)
	)

	noteMock := noteMocks.NewMockRepository(mockCtrl)
	api := NewMockNoteV1(Implementation{
		noteService: note.NewMockNoteService(noteMock),
	})

	noteMock.EXPECT().DeleteNote(ctx, id).Return(nil)
	noteMock.EXPECT().DeleteNote(ctx, id).Return(repoErr)

	t.Run("delete note success", func(t *testing.T) {
		res, err := api.DeleteNote(ctx, req)
		require.Nil(t, err)
		require.Equal(t, &emptypb.Empty{}, res)
	})

	t.Run("delete note repo err", func(t *testing.T) {
		_, err := api.DeleteNote(ctx, req)
		require.NotNil(t, err)
		require.Equal(t, repoErrText, err.Error())
	})
}
