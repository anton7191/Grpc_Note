package main

import (
	"context"
	desc "github.com/anton7191/Note-server-api/pkg/note_v1"
	"google.golang.org/grpc"
	"log"
)

const address = "localhost:2406"

func main() {

	con, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("didn't connect: %s", err.Error())
	}
	defer con.Close()
	client := desc.NewNoteV1Client(con)
	res, err := client.CreateNote(context.Background(), &desc.CreateNoteRequest{
		Title:  "First",
		Text:   "Help me!",
		Author: "Anton",
	})
	if err != nil {
		log.Fatalf(err.Error())
	}
	log.Println("--Create Note--")
	log.Println("Id: ", res.Id)

	resGetnote, err := client.GetNote(context.Background(), &desc.GetNoteRequest{
		Id: 1,
	})
	if err != nil {
		log.Fatalf(err.Error())
	}
	log.Println("--Get Note--")
	log.Println("Title: ", resGetnote.Title)
	log.Println("Text: ", resGetnote.Text)
	log.Println("Autor: ", resGetnote.Author)

	resUpdatenote, err := client.UpdateNote(context.Background(), &desc.UpdateNoteRequest{
		Id:   1,
		Text: "New text",
	})
	if err != nil {
		log.Fatalf(err.Error())
	}
	log.Println("--Status update Note--")
	log.Println("Status:", resUpdatenote.Status)

	resDeletenote, err := client.DeleteNote(context.Background(), &desc.DeleteNoteRequest{
		Id: 1,
	})
	if err != nil {
		log.Fatalf(err.Error())
	}
	log.Println("--Status delete Note--")
	log.Println("Status:", resDeletenote.Status)

	resListnote, err := client.GetListNote(context.Background(), &desc.GetListNoteRequest{
		Req: true,
	})
	if err != nil {
		log.Fatalf(err.Error())
	}
	log.Println("--List Note--")
	log.Println("List Note:", resListnote.Note)
}
