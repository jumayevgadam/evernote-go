package middlewares

import (
	"net/http"
	"net/http/httputil"

	"github.com/gin-gonic/gin"
)

// Debug middleware.
func (mw *MiddlewareManager) DebugMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if mw.Cfg.Server.Mode == "Development" {
			dump, err := httputil.DumpRequest(ctx.Request, true)
			if err != nil {
				ctx.AbortWithStatus(http.StatusInternalServerError)
				return
			}

			mw.Logger.Infof("\nRequest dump begin :--------------\n\n%s\n\nRequest dump end :--------------", string(dump))
		}

		ctx.Next()
	}
}
