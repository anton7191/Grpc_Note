package note_v1

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	desc "github.com/anton7191/note-server-api/pkg/note_v1"
	"github.com/jmoiron/sqlx"
)

func (n *Implementation) UpdateNote(ctx context.Context, req *desc.UpdateNoteRequest) (*desc.Empty, error) {
	fmt.Println("Update Note")
	dbDsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, port, dbUser, dbPassword, dbName, sslMode,
	)
	db, err := sqlx.Open("pgx", dbDsn)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	builder := sq.Update(noteTable).
		PlaceholderFormat(sq.Dollar).
		SetMap(sq.Eq{
			"title":  req.Note.GetTitle(),
			"text":   req.Note.GetText(),
			"author": req.Note.GetAuthor(),
		}).
		Where(sq.Eq{"id": req.Note.GetId()})

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(query, args...)
	if err != nil {
		return nil, err
	}

	return &desc.Empty{}, nil
}
