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
func (mw *MDWManager) AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		jwtToken, err := mw.extractBearerToken(ctx.Request.Header.Get("Authorization"))
		if err != nil {
			httpError.Response(ctx, err)
			return
		}

		claims, err := helpers.ParseAccessToken(jwtToken)
		if err != nil {
			httpError.Response(ctx, httpError.NewUnauthorizedError(err.Error()))
			return
		}

		ctx.Set(UserCtx, claims.UserID)

		ctx.Next()
	}
}

func (mw *MDWManager) extractBearerToken(header string) (string, error) {
	if header == "" {
		return "", errors.New("header value is empty")
	}

	jwtToken := strings.Split(header, " ")
	if len(jwtToken) != 2 || jwtToken[0] != "Bearer" {
		return "", errors.New("incorrectly formatted authorization header")
	}

	return jwtToken[1], nil
}
