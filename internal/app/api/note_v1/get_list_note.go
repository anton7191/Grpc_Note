package note_v1

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	desc "github.com/anton7191/note-server-api/pkg/note_v1"
	"github.com/jmoiron/sqlx"
)

func (n *Implementation) GetListNote(ctx context.Context, req *desc.Empty) (*desc.GetListNoteResponse, error) {
	fmt.Println("GetListNote")

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

	noteList := &desc.GetListNoteResponse{}

	for row.Next() {
		var noteFromDb desc.Note
		err = row.Scan(&noteFromDb.Id, &noteFromDb.Title, &noteFromDb.Text, &noteFromDb.Author)
		if err != nil {
			return nil, err
		}
		noteList.Note = append(noteList.Note, &noteFromDb)
	}

	return noteList, nil
}
