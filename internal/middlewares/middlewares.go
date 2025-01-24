package middlewares

import (
	"github.com/jumayevgadam/evernote-go/internal/config"
	"github.com/jumayevgadam/evernote-go/pkg/logger"
)

// MiddlewareManager.
type MDWManager struct {
	Cfg    *config.Config
	Logger logger.Logger
}

// NewMiddlewareManager creates a new instance of MiddlewareManager.
func NewMiddlewareManager(cfg *config.Config, logger logger.Logger) *MDWManager {
	return &MDWManager{Cfg: cfg, Logger: logger}
}
