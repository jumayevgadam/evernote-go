package users

import "github.com/gin-gonic/gin"

// Handler interface for managing users.
type Handler interface {
	SignUp() gin.HandlerFunc
}
