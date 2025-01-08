package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/jumayevgadam/evernote-go/internal/users"
)

// MapUserRoutes func keeps endpoints for users.
func MapUserRoutes(userGroup *gin.RouterGroup, h users.Handler) {
	userGroup.POST("/sign-up", h.SignUp())
}
