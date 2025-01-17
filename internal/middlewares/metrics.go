package middlewares

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jumayevgadam/evernote-go/internal/metrics"
)

// Prometheus metrics middleware.
func (mw *MiddlewareManager) MetricsMiddleware(metrics metrics.Metrics) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		// Calculate status and response time.
		status := c.Writer.Status()
		latency := time.Since(start).Seconds()

		// Record metrics.
		metrics.ObserveResponseTime(status, c.Request.Method, c.FullPath(), latency)
		metrics.IncHits(status, c.Request.Method, c.FullPath())
	}
}
