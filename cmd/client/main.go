package main

import (
	"context"
	"log"

	desc "github.com/anton7191/note-server-api/pkg/note_v1"
	"google.golang.org/grpc"
)

const address = "localhost:5161"

func main() {
	ctx := context.Background()
	con, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to grpc connect: %s", err.Error())
	}
	defer con.Close()

	client := desc.NewNoteV1Client(con)
	res, err := client.CreateNote(ctx, &desc.CreateNoteRequest{
		Title:  "First",
		Text:   "Help me!",
		Author: "Anton",
	})
	if err != nil {
		log.Fatalf(err.Error())
	}
	log.Println("--Create Note--")
	log.Println("Id: ", res.GetId())

	resGetnote, err := client.GetNote(ctx, &desc.GetNoteRequest{
		Id: 2,
	})
	if err != nil {
		log.Fatalf(err.Error())
	}
	log.Println("--Get Note--")
	log.Println("ID: ", resGetnote.Note.GetId())
	log.Println("Title: ", resGetnote.Note.GetTitle())
	log.Println("Text: ", resGetnote.Note.GetText())
	log.Println("Autor: ", resGetnote.Note.GetAuthor())

	_, err = client.UpdateNote(ctx, &desc.UpdateNoteRequest{
		Note: &desc.Note{
			Id:     10,
			Title:  "new Title",
			Text:   "new Text",
			Author: "new Author"},
	})
	if err != nil {
		log.Fatalf(err.Error())
	}
	log.Println("--Status update Note--")

	_, err = client.DeleteNote(ctx, &desc.DeleteNoteRequest{
		Id: 6,
	})
	if err != nil {
		log.Fatalf(err.Error())
	}
	log.Println("--Status delete Note--")

	resListnote, err := client.GetListNote(ctx, &desc.Empty{})
	if err != nil {
		log.Fatalf(err.Error())
	}
	log.Println("--List Note--")
	log.Println("List Note:", resListnote.GetNote())
}
