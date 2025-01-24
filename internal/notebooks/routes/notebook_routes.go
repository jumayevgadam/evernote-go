package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jumayevgadam/evernote-go/internal/notebooks"
)

func MapNotebookRoutes(r *gin.RouterGroup, h notebooks.Handler) {
	r.POST("", h.AddNotebook())
	r.GET("", h.ListNotebooks())
}
