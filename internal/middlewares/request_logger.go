package middlewares

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jumayevgadam/evernote-go/pkg/utils"
)

// RequestLoggerMiddleware logs details about each HTTP request.
func (mw *MDWManager) RequestLoggerMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()

		ctx.Next()

		req := ctx.Request
		status := ctx.Writer.Status() // Get status from the response writer.
		size := ctx.Writer.Size()     // Get size from the response writer.
		latency := time.Since(start).String()
		requestID := utils.GetRequestID(ctx)

		mw.Logger.Infof(
			"RequestID: %s, Method: %s, URI: %s, Status: %v, Size: %v, Time: %s",
			requestID, req.Method, req.RequestURI, status, size, latency,
		)
	}
}
