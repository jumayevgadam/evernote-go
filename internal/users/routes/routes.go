package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jumayevgadam/evernote-go/internal/users"
)

func MapUserRoutes(r *gin.RouterGroup, h users.Handler) {
	r.POST("/register", h.SignUp())
}
