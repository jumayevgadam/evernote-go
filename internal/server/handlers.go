package server

import (
	"github.com/gin-gonic/gin"
	userHandler "github.com/jumayevgadam/evernote-go/internal/users/handler"
	userRoutes "github.com/jumayevgadam/evernote-go/internal/users/routes"
	userService "github.com/jumayevgadam/evernote-go/internal/users/service"
)

func (s *Server) MapHandlers() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/api/v1")

	// init services.
	userService := userService.NewUserService(s.DataStore)

	// init handlers.
	userHandler := userHandler.NewUserHandler(userService)

	// init routes.
	userRoutes.MapUserRoutes(v1, userHandler)

	return r
}
