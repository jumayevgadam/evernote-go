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
	notebookHandler "github.com/jumayevgadam/evernote-go/internal/notebooks/handler"
	notebookRoutes "github.com/jumayevgadam/evernote-go/internal/notebooks/routes"
	notebookService "github.com/jumayevgadam/evernote-go/internal/notebooks/service"
	userHandler "github.com/jumayevgadam/evernote-go/internal/users/handler"
	userRoutes "github.com/jumayevgadam/evernote-go/internal/users/routes"
	userService "github.com/jumayevgadam/evernote-go/internal/users/service"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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
	notebookService := notebookService.NewNotebookService(s.DataStore)

	// init handlers.
	userHandler := userHandler.NewUserHandler(userService)
	notebookHandler := notebookHandler.NewNotebookHandler(notebookService)

	// init middleware manager.
	mw := middlewares.NewMiddlewareManager(s.Cfg, s.Logger)

	// create a new gin instance.
	r := gin.New()

	r.Use(
		gin.Logger(),
		gin.Recovery(),
	)

	// add swagger url.
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.RedirectTrailingSlash = false
	r.RedirectFixedPath = true

	// metrics and request logger middleware.
	r.Use(
		mw.MetricsMiddleware(metrics),
		mw.RequestLoggerMiddleware(),
	)

	// other middlewares.
	r.Use(
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

	// v1 subgroups.
	authGroup := v1.Group("/auth")
	notebookGroup := v1.Group("/notebooks")

	// init routes.
	userRoutes.MapUserRoutes(authGroup, userHandler)

	notebookGroup.Use(mw.AuthMiddleware())
	notebookRoutes.MapNotebookRoutes(notebookGroup, notebookHandler)

	return r
}
