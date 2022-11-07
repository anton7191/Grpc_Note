package app

import (
	"context"
	"log"

	"github.com/anton7191/note-server-api/internal/config"
	"github.com/anton7191/note-server-api/internal/pkg/db"
	note2 "github.com/anton7191/note-server-api/internal/repository/note"
	"github.com/anton7191/note-server-api/internal/service/note"
)

type serviceProvider struct {
	db             db.Client
	configPath     string
	config         *config.Config
	noteRepository note2.Repository
	noteService    *note.Service
}

func newServiceProvider(configPath string) *serviceProvider {
	return &serviceProvider{
		configPath: configPath,
	}
}

func (s *serviceProvider) GetConfig() *config.Config {
	if s.config == nil {
		cfg, err := config.NewConfig(s.configPath)
		if err != nil {
			log.Fatalf("failed to get config: %s", err.Error())
		}

		s.config = cfg
	}

	return s.config
}

func (s *serviceProvider) GetDB(ctx context.Context) db.Client {
	if s.db == nil {
		cfg, err := s.GetConfig().GetDBConfig()
		if err != nil {
			log.Fatalf("failed to get db config: %s", err)
		}

		dbc, err := db.NewClient(ctx, cfg)
		if err != nil {
			log.Fatalf("can't connect to db err: %s", err)
		}
		s.db = dbc
	}

	return s.db
}

func (s *serviceProvider) GetNoteRepository(ctx context.Context) note2.Repository {
	if s.noteRepository == nil {
		s.noteRepository = note2.NewNoteRepository(s.GetDB(ctx))
	}

	return s.noteRepository
}

func (s *serviceProvider) GetNoteService(ctx context.Context) *note.Service {
	if s.noteService == nil {
		s.noteService = note.NewService(s.GetNoteRepository(ctx))
	}

	return s.noteService
}
