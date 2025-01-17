package utils

import "github.com/gin-gonic/gin"

// GetRequestID func retrieve request ID.
func GetRequestID(c *gin.Context) string {
	return c.Writer.Header().Get("X-Request-ID")
}
