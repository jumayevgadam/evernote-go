package middlewares

import (
	"github.com/gin-gonic/gin"
)

// GetUserIDFromCtx func retrieves userID from ctx.
func GetUserIDFromCtx(c *gin.Context) (int, error) {
	id := c.GetInt(UserCtx)

	return id, nil
}
