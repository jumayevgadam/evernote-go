package server

import (
	_ "github.com/jumayevgadam/evernote-go/docs"
	"github.com/jumayevgadam/evernote-go/internal/config"
	"github.com/jumayevgadam/evernote-go/internal/database"
	"github.com/jumayevgadam/evernote-go/pkg/logger"
)

// Server struct keeps needed confs for server.
type Server struct {
	Cfg       *config.Config
	DataStore database.DataStore
	Logger    logger.Logger
}

// NewServer creates and returns a new instance of Server.
func NewServer(
	cfg *config.Config,
	dataStore database.DataStore,
	logger logger.Logger,
) *Server {
	return &Server{
		Cfg:       cfg,
		DataStore: dataStore,
		Logger:    logger,
	}
}
