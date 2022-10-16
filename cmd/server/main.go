package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"sync"

	"github.com/anton7191/note-server-api/internal/app/api/note_v1"
	desc "github.com/anton7191/note-server-api/pkg/note_v1"
	grpcValidator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
)

const (
	hostGrpc = "localhost:5161"
	hostHttp = "localhost:8090"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		defer wg.Done()
		startGRPC()
	}()

	//go func() {
	//	defer wg.Done()
	//	startHttp()
	//}()

	wg.Wait()
}

func startGRPC() error {
	list, err := net.Listen("tcp", hostGrpc)
	if err != nil {
		log.Fatalf("failed to mapping port: %s", err.Error())
	}

	s := grpc.NewServer(
		grpc.UnaryInterceptor(grpcValidator.UnaryServerInterceptor()),
	)
	desc.RegisterNoteV1Server(s, note_v1.NewNote())

	fmt.Println("GRPC server is running on port: ", hostGrpc)
	if err = s.Serve(list); err != nil {
		log.Fatalf("failed the serve: %s", err.Error())
		return err
	}

	return nil
}

func startHttp() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := desc.RegisterNoteV1HandlerFromEndpoint(ctx, mux, hostGrpc, opts)
	fmt.Println("Http server is running on port: ", hostHttp)
	if err != nil {
		return err
	}

	return http.ListenAndServe(hostHttp, mux)
}
