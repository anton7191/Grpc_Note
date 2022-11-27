package converter

import (
	"database/sql"

	"github.com/anton7191/note-server-api/internal/model"
	desc "github.com/anton7191/note-server-api/pkg/note_v1"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func ToNoteInfo(noteInfo *desc.NoteInfo) *model.NoteInfo {
	return &model.NoteInfo{
		Title:  noteInfo.GetTitle(),
		Text:   noteInfo.GetText(),
		Author: noteInfo.GetAuthor(),
	}
}

func ToDescNoteInfo(noteInfo *model.NoteInfo) *desc.NoteInfo {
	if noteInfo == nil {
		return nil
	}
	return &desc.NoteInfo{
		Title:  noteInfo.Title,
		Text:   noteInfo.Text,
		Author: noteInfo.Author,
	}
}

func ToUpdateNoteInfo(updateNoteInfo *desc.UpdateNoteInfo) *model.UpdateNoteInfo {
	var title, text, author sql.NullString
	if updateNoteInfo.GetTitle() != nil {
		title = sql.NullString{
			String: updateNoteInfo.GetTitle().GetValue(),
			Valid:  updateNoteInfo.GetTitle() != nil,
		}
	}

	if updateNoteInfo.GetText() != nil {
		text = sql.NullString{
			String: updateNoteInfo.GetText().GetValue(),
			Valid:  updateNoteInfo.GetText() != nil,
		}
	}

	if updateNoteInfo.GetAuthor() != nil {
		author = sql.NullString{
			String: updateNoteInfo.GetAuthor().GetValue(),
			Valid:  updateNoteInfo.GetAuthor() != nil,
		}
	}

	return &model.UpdateNoteInfo{
		Title:  title,
		Text:   text,
		Author: author,
	}
}

func ToDescUpdateNoteInfo(updateNoteInfo *model.UpdateNoteInfo) *desc.UpdateNoteInfo {
	var title, text, author *wrapperspb.StringValue
	if updateNoteInfo.Title.Valid {
		title = &wrapperspb.StringValue{
			Value: updateNoteInfo.Title.String,
		}
	}
	if updateNoteInfo.Text.Valid {
		text = &wrapperspb.StringValue{
			Value: updateNoteInfo.Text.String,
		}
	}
	if updateNoteInfo.Author.Valid {
		author = &wrapperspb.StringValue{
			Value: updateNoteInfo.Author.String,
		}
	}

	return &desc.UpdateNoteInfo{
		Title:  title,
		Text:   text,
		Author: author,
	}
}

func ToDescNote(note *model.Note) *desc.Note {
	var updatedAt *timestamppb.Timestamp
	if note.UpdatedAt.Valid {
		updatedAt = timestamppb.New(note.UpdatedAt.Time)
	}

	return &desc.Note{
		Id:        note.ID,
		Note:      ToDescNoteInfo(note.Info),
		CreatedAt: timestamppb.New(note.CreatedAt),
		UpdatedAt: updatedAt,
	}
}

func ToDescNoteSlice(notes []*model.Note) []*desc.Note {
	var noteSlice []*desc.Note
	for _, v := range notes {
		noteSlice = append(noteSlice, ToDescNote(v))
	}

	return noteSlice
}
