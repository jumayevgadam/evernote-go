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

	srv := server.NewServer(cfg, dataStore, appLogger)
	r := srv.MapHandlers()

	httpServer := &http.Server{
		Addr:         ":" + cfg.Server.Port,
		Handler:      r,
		ReadTimeout:  cfg.Server.ReadTimeout,
		WriteTimeout: cfg.Server.WriteTimeout,
	}

	go func() {
		appLogger.Infof("server started on http port: %s", cfg.Server.Port)

		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			appLogger.Errorf("error in running server: %v", err.Error())
		}
	}()

	// graceful shutdown.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := httpServer.Shutdown(ctx); err != nil {
		appLogger.Errorf("failed to stop server: %v", err.Error())
	}
}
