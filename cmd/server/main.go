package main

import (
	"fmt"
	"log"
	"net"
	desc "github.com/anton7191/testGrpc/pkg/note_v1"
	"github.com/anton7191/testGrpc/internal/app/api/note_v1"
	"google.golang.org/grpc"
)

const port = ":2406"

func main() {
	list, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to mapping port: %s", err.Error())
	}
	
	s := grpc.NewServer()
	desc.RegisterNoteServer(s, note_v1.NewNote())
	
	if err = s.Server(list); err != nil {
		log.Fatalf("failed te server: %s", err.Error())
	}
	fmt.Println() 
}