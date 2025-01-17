package server

import (
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	limits "github.com/gin-contrib/size"
	"github.com/gin-gonic/gin"
	"github.com/jumayevgadam/evernote-go/internal/metrics"
	"github.com/jumayevgadam/evernote-go/internal/middlewares"
	userHandler "github.com/jumayevgadam/evernote-go/internal/users/handler"
	userRoutes "github.com/jumayevgadam/evernote-go/internal/users/routes"
	userService "github.com/jumayevgadam/evernote-go/internal/users/service"
)

func (s *Server) MapHandlers() *gin.Engine {
	metrics, err := metrics.CreateMetrics(s.Cfg.Metrics.URL, s.Cfg.Metrics.ServiceName)
	if err != nil {
		s.Logger.Errorf("create metrics error: %v", err.Error())
	}
	s.Logger.Infof(
		"Metrics available url: %s, ServiceName: %s",
		s.Cfg.Metrics.URL,
		s.Cfg.Metrics.ServiceName,
	)

	// init services.
	userService := userService.NewUserService(s.DataStore)

	// init handlers.
	userHandler := userHandler.NewUserHandler(userService)

	// init middleware manager.
	mw := middlewares.NewMiddlewareManager(s.Cfg, s.Logger)

	// create a new gin instance.
	r := gin.New()

	r.RedirectTrailingSlash = false
	r.RedirectFixedPath = true

	// metrics and request logger middleware.
	r.Use(
		mw.MetricsMiddleware(metrics),
		mw.RequestLoggerMiddleware(),
	)

	// other middlewares.
	r.Use(
		gin.Logger(),
		gin.Recovery(),
		cors.New(cors.Config{
			AllowAllOrigins: true,
			AllowMethods:    []string{"POST", "GET", "PUT", "DELETE", "PATCH", "HEAD", "OPTIONS"},
			AllowHeaders:    []string{"Origin", "Content-Type", "Authorization", "Content-Length"},
			MaxAge:          12 * time.Hour,
		}),
		gzip.Gzip(gzip.DefaultCompression),
		limits.RequestSizeLimiter(100),
	)

	// health check.
	r.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "hello")
	})

	// v1 group.
	v1 := r.Group("/api/v1")

	// auth group.
	authGroup := v1.Group("/auth")

	// init routes.
	userRoutes.MapUserRoutes(authGroup, userHandler)

	return r
}
