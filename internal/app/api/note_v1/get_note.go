package note_v1

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	desc "github.com/anton7191/note-server-api/pkg/note_v1"
	"github.com/jmoiron/sqlx"
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

	builder := sq.Select("id", "title", "text", "author").
		PlaceholderFormat(sq.Dollar).
		From(noteTable).
		Where(sq.Eq{"id": req.GetId()})

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
	var idFromDb int64
	var titleFromDb string
	var textFromDb string
	var authorFromDb string

	err = row.Scan(&idFromDb, &titleFromDb, &textFromDb, &authorFromDb)
	if err != nil {
		return nil, err
	}

	return &desc.GetNoteResponse{
		Note: &desc.Note{
			Id:     idFromDb,
			Author: authorFromDb,
			Text:   textFromDb,
			Title:  titleFromDb,
		},
	}, nil
}
