package note_v1

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	sq "github.com/Masterminds/squirrel"
	desc "github.com/anton7191/note-server-api/pkg/note_v1"
	"github.com/jmoiron/sqlx"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (n *Implementation) GetNote(ctx context.Context, req *desc.GetNoteRequest) (*desc.GetNoteResponse, error) {
	dbDsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, port, dbUser, dbPassword, dbName, sslMode,
	)
	db, err := sqlx.Open("pgx", dbDsn)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	builder := sq.Select("id", "title", "text", "author", "created_at", "updated_at").
		PlaceholderFormat(sq.Dollar).
		From(noteTable).
		Where(sq.Eq{"id": req.GetId()}).
		Limit(1)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	row, err := db.QueryContext(ctx, query, args...)
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
	return &desc.GetNoteResponse{
		Note: &desc.Note{
			Id:        id,
			Author:    author,
			Text:      text,
			Title:     title,
			CreatedAt: timestamppb.New(createdAt),
			UpdatedAt: updatedAtPb,
		},
	}, nil
}
