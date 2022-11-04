package repository

import (
	"context"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/anton7191/note-server-api/internal/model"
	"github.com/anton7191/note-server-api/internal/pkg/db"
	"github.com/anton7191/note-server-api/internal/repository/table"
)

type NoteRepository interface {
	CreateNote(ctx context.Context, noteInfo *model.NoteInfo) (int64, error)
	GetNote(ctx context.Context, id int64) (*model.Note, error)
	UpdateNote(ctx context.Context, id int64, updateNoteInfo *model.UpdateNoteInfo) error
	DeleteNote(ctx context.Context, id int64) error
	GetListNote(ctx context.Context) ([]*model.Note, error)
}

type repository struct {
	client db.Client
}

func NewNoteRepository(client db.Client) NoteRepository {
	return &repository{
		client: client,
	}
}

func (r *repository) CreateNote(ctx context.Context, noteInfo *model.NoteInfo) (int64, error) {
	builder := sq.Insert(table.Note).
		PlaceholderFormat(sq.Dollar).
		Columns("title", "text", "author").
		Values(noteInfo.Title, noteInfo.Text, noteInfo.Author).
		Suffix("returning id")

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, err
	}

	q := db.Query{
		Name:     "CreateNote",
		QueryRaw: query,
	}

	row, err := r.client.DB().QueryContext(ctx, q, args...)
	if err != nil {
		return 0, err
	}
	defer row.Close()

	row.Next()
	var id int64
	err = row.Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}
func (r *repository) GetNote(ctx context.Context, id int64) (*model.Note, error) {
	builder := sq.Select("id", "title", "text", "author", "created_at", "updated_at").
		PlaceholderFormat(sq.Dollar).
		From(table.Note).
		Where(sq.Eq{"id": id}).
		Limit(1)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "GetNote",
		QueryRaw: query,
	}

	note := new(model.Note)

	err = r.client.DB().GetContext(ctx, note, q, args...)
	if err != nil {
		return nil, err
	}

	return note, nil
}
func (r *repository) UpdateNote(ctx context.Context, id int64, updateNoteInfo *model.UpdateNoteInfo) error {
	builder := sq.Update(table.Note).
		PlaceholderFormat(sq.Dollar).
		Set("updated_at", time.Now()).
		Where(sq.Eq{"id": id})

	if updateNoteInfo.Title.Valid {
		builder.Set("title", updateNoteInfo.Title.String)
	}
	if updateNoteInfo.Text.Valid {
		builder.Set("text", updateNoteInfo.Text.String)
	}
	if updateNoteInfo.Author.Valid {
		builder.Set("author", updateNoteInfo.Author.String)
	}

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "UpdateNote",
		QueryRaw: query,
	}

	_, err = r.client.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return nil
}
func (r *repository) DeleteNote(ctx context.Context, id int64) error {
	builder := sq.Delete(table.Note).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"id": id})

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "DeleteNote",
		QueryRaw: query,
	}

	_, err = r.client.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return nil
}
func (r *repository) GetListNote(ctx context.Context) ([]*model.Note, error) {
	builder := sq.Select("id", "title", "text", "author", "created_at", "updated_at").
		PlaceholderFormat(sq.Dollar).
		From(table.Note)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "GetListNote",
		QueryRaw: query,
	}

	var notes []*model.Note
	err = r.client.DB().SelectContext(ctx, notes, q, args...)
	if err != nil {
		return nil, err
	}

	return notes, nil
}
