package app

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/jumayevgadam/evernote-go/internal/config"
	"github.com/jumayevgadam/evernote-go/internal/connection"
	"github.com/jumayevgadam/evernote-go/internal/database/postgres"
	"github.com/jumayevgadam/evernote-go/internal/server"
	"github.com/jumayevgadam/evernote-go/pkg/logger"
	"go.uber.org/zap"
)

// Run func starts application.
func Run(configPath string) {
	cfgFile, err := config.LoadConfig(configPath)
	if err != nil {
		zap.L().Error("app.config.LoadConfig: error", zap.Error(err))
		// return
	}

	cfg, err := config.ParseConfig(cfgFile)
	if err != nil {
		zap.L().Error("app.config.ParseConfig: error", zap.Error(err))
		// return
	}

	appLogger := logger.NewAPILogger(cfg)
	appLogger.InitLogger()

	psqlDB, err := connection.GetDBConnection(context.Background(), cfg.Postgres)
	if err != nil {
		appLogger.Errorf("app.connection.GetDBConnection: %v", err.Error())
		return
	}

	defer func() {
		psqlDB.Close()
		appLogger.Info("database connection closed successfully")
	}()

	dataStore := postgres.NewDataStore(psqlDB)

	// HTTP server.
	srv := server.NewServer(cfg, dataStore, appLogger)

	go func() {
		appLogger.Infof("server started on http port: %s", srv.HttpServer.Addr)

		if err := srv.Run(); err != nil && err != http.ErrServerClosed {
			appLogger.Errorf("error occured while running http server: %v", err.Error())
		}
	}()

	srv.MapHandlers()

	appLogger.Info("server started")

	// graceful shutdown.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	appLogger.Info("shutdown server")

	const timeout = 5 * time.Second

	ctx, shutdown := context.WithTimeout(context.Background(), timeout)
	defer shutdown()

	if err := srv.Stop(ctx); err != nil {
		appLogger.Errorf("failed to stop server: %v", err.Error())
	}
}
