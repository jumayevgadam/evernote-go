package server

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jumayevgadam/evernote-go/internal/config"
	"github.com/jumayevgadam/evernote-go/internal/database"
	"github.com/jumayevgadam/evernote-go/pkg/logger"
)

// Server struct keeps needed confs for server.
type Server struct {
	HttpServer *http.Server
	Cfg        *config.Config
	DataStore  database.DataStore
	Logger     logger.Logger
}

// NewServer creates and returns a new instance of Server.
func NewServer(
	cfg *config.Config,
	dataStore database.DataStore,
	logger logger.Logger,
) *Server {
	ginEngine := gin.New()

	return &Server{
		HttpServer: &http.Server{
			Addr:         ":" + cfg.Server.Port,
			Handler:      ginEngine,
			ReadTimeout:  cfg.Server.ReadTimeout,
			WriteTimeout: cfg.Server.WriteTimeout,
		},
		Logger:    logger,
		Cfg:       cfg,
		DataStore: dataStore,
	}
}

func (s *Server) Run() error {
	return s.HttpServer.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	return s.HttpServer.Shutdown(ctx)
}
