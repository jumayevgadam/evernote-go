package server

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	userHttp "github.com/jumayevgadam/evernote-go/internal/users/handler"
	userService "github.com/jumayevgadam/evernote-go/internal/users/service"
)

// MapHandlers method keeps all needed middlewares and endpoints.
func (s *Server) MapHandlers() {
	g := gin.Default()

	// // create metrics.
	// _, err := metrics.CreateMetrics(s.Cfg.Metrics.URL, s.Cfg.Metrics.ServiceName)
	// if err != nil {
	// 	s.Logger.Errorf("create metrics error: %v", err)
	// }

	// s.Logger.Infof(
	// 	"Metrics available URL: %s, ServiceName: %s",
	// 	s.Cfg.Metrics.URL,
	// 	s.Cfg.Metrics.ServiceName,
	// )

	// init services.
	userSrv := userService.NewUserService(s.DataStore)

	// init handlers.
	userHandler := userHttp.NewUserHandler(userSrv)

	// init middleware this place.

	// init other middlewares and cors.
	g.Use(
		// gin.Recovery(),
		cors.New(
			cors.Config{
				AllowMethods:     []string{"POST", "GET", "PUT", "PATCH", "DELETE", "OPTION", "HEAD"},
				AllowHeaders:     []string{"Origin", "Content-Type", "Content-Length", "Authorization"},
				AllowCredentials: false,
				AllowAllOrigins:  true,
				MaxAge:           12 * time.Hour,
			},
		),
		// gzip.Gzip(gzip.DefaultCompression),
		//limiter.RequestSizeLimiter(10),
	)

	// v1 group.
	v1 := g.Group("/api/v1")

	// auth group.
	authGroup := v1.Group("/auth")

	// routes.
	userHttp.MapUserRoutes(authGroup, userHandler)
}
