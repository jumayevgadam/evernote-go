package middlewares

import (
	"errors"

	"github.com/gin-gonic/gin"
)

// GetUserIDFromCtx func retrieves userID from ctx.
func GetUserIDFromCtx(c *gin.Context) (int, error) {
	id, ok := c.Get(UserCtx)
	if !ok {
		return 0, errors.New("user id not found in the ctx")
	}

	userID, ok := id.(int)
	if !ok {
		return 0, errors.New("user id is not int type")
	}

	return userID, nil
}
