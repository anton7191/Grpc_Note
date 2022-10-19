package repository

import (
	"context"
	"database/sql"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/anton7191/note-server-api/internal/repository/table"
	desc "github.com/anton7191/note-server-api/pkg/note_v1"
	"github.com/jmoiron/sqlx"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type NoteRepository interface {
	CreateNote(ctx context.Context, req *desc.CreateNoteRequest) (int64, error)
	GetNote(ctx context.Context, req *desc.GetNoteRequest) (*desc.Note, error)
	UpdateNote(ctx context.Context, req *desc.UpdateNoteRequest) error
	DeleteNote(ctx context.Context, req *desc.DeleteNoteRequest) error
	GetListNote(ctx context.Context, req *desc.Empty) ([]*desc.Note, error)
}

type repository struct {
	db *sqlx.DB
}

func NewNoteRepository(db *sqlx.DB) NoteRepository {
	return &repository{
		db: db,
	}
}

func (r *repository) CreateNote(ctx context.Context, req *desc.CreateNoteRequest) (int64, error) {
	builder := sq.Insert(table.Note).
		PlaceholderFormat(sq.Dollar).
		Columns("title", "text", "author").
		Values(req.GetTitle(), req.GetText(), req.GetAuthor()).
		Suffix("returning id")

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, err
	}

	row, err := r.db.QueryContext(ctx, query, args...)
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
func (r *repository) GetNote(ctx context.Context, req *desc.GetNoteRequest) (*desc.Note, error) {
	builder := sq.Select("id", "title", "text", "author", "created_at", "updated_at").
		PlaceholderFormat(sq.Dollar).
		From(table.Note).
		Where(sq.Eq{"id": req.GetId()}).
		Limit(1)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	row, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer row.Close()

	row.Next()
	var id int64
	var title string
	var text string
	var author string
	var createdAt time.Time
	var updatedAt sql.NullTime

	err = row.Scan(&id, &title, &text, &author, &createdAt, &updatedAt)
	if err != nil {
		return nil, err
	}

	var updatedAtPb *timestamppb.Timestamp
	if updatedAt.Valid {
		updatedAtPb = timestamppb.New(updatedAt.Time)
	}

	return &desc.Note{
		Id:        id,
		Author:    author,
		Text:      text,
		Title:     title,
		CreatedAt: timestamppb.New(createdAt),
		UpdatedAt: updatedAtPb,
	}, nil
}
func (r *repository) UpdateNote(ctx context.Context, req *desc.UpdateNoteRequest) error {
	builder := sq.Update(table.Note).
		PlaceholderFormat(sq.Dollar).
		SetMap(sq.Eq{
			"title":      req.Note.GetTitle(),
			"text":       req.Note.GetText(),
			"author":     req.Note.GetAuthor(),
			"updated_at": time.Now(),
		}).
		Where(sq.Eq{"id": req.Note.GetId()})

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	_, err = r.db.Exec(query, args...)
	if err != nil {
		return err
	}

	return nil
}
func (r *repository) DeleteNote(ctx context.Context, req *desc.DeleteNoteRequest) error {
	builder := sq.Delete(table.Note).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"id": req.GetId()})

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	_, err = r.db.Exec(query, args...)
	if err != nil {
		return err
	}

	return nil
}
func (r *repository) GetListNote(ctx context.Context, req *desc.Empty) ([]*desc.Note, error) {
	builder := sq.Select("id", "title", "text", "author", "created_at", "updated_at").
		PlaceholderFormat(sq.Dollar).
		From(table.Note)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	row, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer row.Close()

	noteList := []*desc.Note{}
	var createdAt time.Time
	var updatedAt sql.NullTime

	for row.Next() {
		var note desc.Note
		var updatedAtPb *timestamppb.Timestamp
		err = row.Scan(&note.Id, &note.Title, &note.Text, &note.Author, &createdAt, &updatedAt)
		if err != nil {
			return nil, err
		}
		note.CreatedAt = timestamppb.New(createdAt)
		if updatedAt.Valid {
			updatedAtPb = timestamppb.New(updatedAt.Time)
		}
		note.UpdatedAt = updatedAtPb
		noteList = append(noteList, &note)
	}
	return noteList, nil
}
