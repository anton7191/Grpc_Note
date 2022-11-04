package app

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

type App struct {
	noteImpl        *note_v1.Implementation
	serviceProvider *serviceProvider

	pathConfig string

	grpcServer *grpc.Server
	mux        *runtime.ServeMux
}

func NewApp(ctx context.Context, pathConfig string) (*App, error) {
	a := &App{
		pathConfig: pathConfig,
	}
	err := a.initDeps(ctx)
	if err != nil {
		return nil, err
	}

	return a, nil
}

func (a *App) Run() error {
	defer func() {
		a.serviceProvider.db.Close()
	}()

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		err := a.runGRPC()
		if err != nil {
			return
		}
	}()

	go func() {
		defer wg.Done()
		err := a.runHttp()
		if err != nil {
			return
		}
	}()

	wg.Wait()
	return nil
}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initServiceProvider,
		a.initServer,
		a.initGRPCServer,
		a.initPublicHTTPHandlers,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initServiceProvider(_ context.Context) error {
	a.serviceProvider = newServiceProvider(a.pathConfig)
	return nil
}

func (a *App) initServer(ctx context.Context) error {
	a.noteImpl = note_v1.NewNote(a.serviceProvider.GetNoteService(ctx))

	return nil
}

func (a *App) initGRPCServer(_ context.Context) error {
	a.grpcServer = grpc.NewServer(
		grpc.UnaryInterceptor(grpcValidator.UnaryServerInterceptor()),
	)
	desc.RegisterNoteV1Server(a.grpcServer, a.noteImpl)
	return nil
}

func (a *App) initPublicHTTPHandlers(ctx context.Context) error {
	a.mux = runtime.NewServeMux()

	opts := []grpc.DialOption{grpc.WithInsecure()} //nolint
	err := desc.RegisterNoteV1HandlerFromEndpoint(ctx, a.mux, a.serviceProvider.config.GRPC.GetAddress(), opts)
	if err != nil {
		return err
	}

	return nil
}

func (a *App) runGRPC() error {
	list, err := net.Listen("tcp", a.serviceProvider.config.GRPC.GetAddress())
	if err != nil {
		log.Fatalf("failed to mapping port: %s", err.Error())
	}

	fmt.Println("GRPC server is running on port: ", a.serviceProvider.config.GRPC.GetAddress())

	if err := a.grpcServer.Serve(list); err != nil {
		log.Fatalf("failed the serve: %s", err.Error())
		return err
	}

	return nil
}

func (a *App) runHttp() error {
	fmt.Println("Http server is running on port: ", a.serviceProvider.config.HTTP.GetAddress())

	if err := http.ListenAndServe(a.serviceProvider.config.HTTP.GetAddress(), a.mux); err != nil {
		log.Fatalf("failed to process muxer: %s", err.Error())
	}

	return nil
}
