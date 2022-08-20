package main

import (
	"context"
	desc "github.com/anton7191/Grpc_Note/pkg/note_v1"
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

	res_getnote, err := client.GetNote(context.Background(), &desc.GetNoteRequest{
		Id: 1,
	})
	if err != nil {
		log.Fatalf(err.Error())
	}
	log.Println("--Get Note--")
	log.Println("Title: ", res_getnote.Title)
	log.Println("Text: ", res_getnote.Text)
	log.Println("Autor: ", res_getnote.Author)

	res_updatenote, err := client.UpdateNote(context.Background(), &desc.UpdateNoteRequest{
		Id:   1,
		Text: "New text",
	})
	if err != nil {
		log.Fatalf(err.Error())
	}
	log.Println("--Status update Note--")
	log.Println("Status:", res_updatenote.Status)

	res_deletenote, err := client.DeleteNote(context.Background(), &desc.DeleteNoteRequest{
		Id: 1,
	})
	if err != nil {
		log.Fatalf(err.Error())
	}
	log.Println("--Status delete Note--")
	log.Println("Status:", res_deletenote.Status)

	res_listnote, err := client.GetListNote(context.Background(), &desc.GetListNoteRequest{
		Req: true,
	})
	if err != nil {
		log.Fatalf(err.Error())
	}
	log.Println("--List Note--")
	log.Println("List Note:", res_listnote.Note)
}
