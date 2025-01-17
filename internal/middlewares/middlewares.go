package middlewares

import (
	"github.com/jumayevgadam/evernote-go/internal/config"
	"github.com/jumayevgadam/evernote-go/pkg/logger"
)

// MiddlewareManager.
type MiddlewareManager struct {
	Cfg    *config.Config
	Logger logger.Logger
}

// NewMiddlewareManager creates a new instance of MiddlewareManager.
func NewMiddlewareManager(cfg *config.Config, logger logger.Logger) *MiddlewareManager {
	return &MiddlewareManager{Cfg: cfg, Logger: logger}
}
