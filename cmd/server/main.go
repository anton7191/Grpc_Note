package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"sync"

	"github.com/anton7191/note-server-api/internal/app/api/note_v1"
	"github.com/anton7191/note-server-api/internal/repository"
	"github.com/anton7191/note-server-api/internal/service/note"
	desc "github.com/anton7191/note-server-api/pkg/note_v1"
	grpcValidator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc"
)

const (
	hostGrpc = "localhost:6151"
	hostHttp = "localhost:8090"
)

const (
	host       = "localhost"
	port       = "54321"
	dbUser     = "note-service-user"
	dbPassword = "note-service-password"
	dbName     = "note-service"
	sslMode    = "disable"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		err := startGRPC()
		if err != nil {
			return
		}
	}()

	go func() {
		defer wg.Done()
		err := startHttp()
		if err != nil {
			return
		}
	}()

	wg.Wait()
}

func startGRPC() error {
	list, err := net.Listen("tcp", hostGrpc)
	if err != nil {
		log.Fatalf("failed to mapping port: %s", err.Error())
	}

	dbDsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, port, dbUser, dbPassword, dbName, sslMode,
	)

	db, err := sqlx.Open("pgx", dbDsn)
	if err != nil {
		return err
	}
	defer db.Close()

	noteRepository := repository.NewNoteRepository(db)
	noteService := note.NewService(noteRepository)

	s := grpc.NewServer(
		grpc.UnaryInterceptor(grpcValidator.UnaryServerInterceptor()),
	)
	desc.RegisterNoteV1Server(s, note_v1.NewNote(noteService))

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
	if err != nil {
		return err
	}

	fmt.Println("Http server is running on port: ", hostHttp)

	return http.ListenAndServe(hostHttp, mux)
}
