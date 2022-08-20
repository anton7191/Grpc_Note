package main

import (
	"fmt"
	"github.com/anton7191/Grpc_Note/internal/app/api/note_v1"
	desc "github.com/anton7191/Grpc_Note/pkg/note_v1"
	"google.golang.org/grpc"
	"log"
	"net"
)

const port = ":2406"

func main() {
	list, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to mapping port: %s", err.Error())
	}

	s := grpc.NewServer()
	desc.RegisterNoteV1Server(s, note_v1.NewNote())

	if err = s.Serve(list); err != nil {
		log.Fatalf("failed the serve: %s", err.Error())
	}
	fmt.Println("Finish")
}
