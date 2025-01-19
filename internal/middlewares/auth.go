package middlewares

import (
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jumayevgadam/evernote-go/internal/helpers"
	"github.com/jumayevgadam/evernote-go/pkg/httpError"
)

const UserCtx = "user_id"

// AuthMiddleware for checking user.
func (mw *MiddlewareManager) AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		jwtToken, err := mw.extractBearerToken(ctx.GetHeader("Authorization"))
		if err != nil {
			httpError.Response(ctx, err)
			return
		}

		userID, err := helpers.ParseAccessToken(jwtToken)
		if err != nil {
			httpError.Response(ctx, err)
			return
		}

		ctx.Set(UserCtx, userID)

		ctx.Next()
	}
}

func (mw *MiddlewareManager) extractBearerToken(header string) (string, error) {
	if header == "" {
		return "", errors.New("header value is empty")
	}

	jwtToken := strings.Split(header, " ")
	if len(jwtToken) != 2 {
		return "", errors.New("incorrectly formatted authorization header")
	}

	return jwtToken[1], nil
}
