package notebooks

import "github.com/gin-gonic/gin"

// Handler interface for notebooks.
type Handler interface {
	AddNotebook() gin.HandlerFunc
	ListNotebooks() gin.HandlerFunc
}
