package main

import (
	"context"
	"log"

	"github.com/anton7191/note-server-api/internal/converter"
	"github.com/anton7191/note-server-api/internal/model"
	desc "github.com/anton7191/note-server-api/pkg/note_v1"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

const address = "localhost:6151"

func main() {
	ctx := context.Background()
	con, err := grpc.Dial(address, grpc.WithInsecure()) //nolint
	if err != nil {
		log.Fatalf("failed to grpc connect: %s", err.Error())
	}
	defer con.Close()

	client := desc.NewNoteV1Client(con)
	res, err := client.CreateNote(ctx, &desc.CreateNoteRequest{
		Note: &desc.NoteInfo{
			Title:  "first",
			Text:   "mycun",
			Author: "Anton",
		},
	})
	if err != nil {
		log.Fatalf(err.Error())
	}
	log.Println("--Create Note--")
	log.Println("Id: ", res.GetId())

	resGetNote, err := client.GetNote(ctx, &desc.GetNoteRequest{
		Id: 10,
	})
	if err != nil {
		log.Fatalf(err.Error())
	}
	log.Println("--Get Note--")
	log.Println("ID: ", resGetNote.GetNote().GetId())
	log.Println("Title: ", resGetNote.GetNote().GetNote().GetTitle())
	log.Println("Text: ", resGetNote.GetNote().GetNote().GetText())
	log.Println("Author: ", resGetNote.GetNote().GetNote().GetAuthor())
	log.Println("Created: ", resGetNote.GetNote().GetCreatedAt().AsTime())
	if resGetNote.GetNote().GetUpdatedAt() != nil {
		log.Println("Updated: ", resGetNote.GetNote().GetUpdatedAt().AsTime())
	} else {
		log.Println("Updated: ", "no updated")
	}

	updateNote := new(model.UpdateNoteInfo)
	updateNote.Title.Scan("NEW")
	updateNote.Text.Scan("NEW NEW")
	updateNote.Author.Scan("New Anton")
	_, err = client.UpdateNote(ctx, &desc.UpdateNoteRequest{
		Id:   40,
		Note: converter.ToDescUpdateNoteInfo(updateNote),
	})
	if err != nil {
		log.Fatalf(err.Error())
	}
	log.Println("--Status update Note--")

	_, err = client.DeleteNote(ctx, &desc.DeleteNoteRequest{
		Id: 9,
	})
	if err != nil {
		log.Fatalf(err.Error())
	}
	log.Println("--Status delete Note--")

	resListnote, err := client.GetListNote(ctx, &emptypb.Empty{})
	if err != nil {
		log.Fatalf(err.Error())
	}
	log.Println("--List Note--")
	log.Println("List Note:", resListnote.GetNote())
}
