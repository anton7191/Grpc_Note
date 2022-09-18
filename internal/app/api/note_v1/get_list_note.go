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

func (n *Implementation) GetListNote(ctx context.Context, req *desc.Empty) (*desc.GetListNoteResponse, error) {
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
		From(noteTable)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	row, err := db.QueryContext(ctx, query, args...)
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

	return &desc.GetListNoteResponse{Note: noteList}, nil
}
