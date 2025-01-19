package helpers

import (
	"errors"
	"os"

	"github.com/golang-jwt/jwt/v5"
	jwtModel "github.com/jumayevgadam/evernote-go/internal/models/jwt"
	"github.com/jumayevgadam/evernote-go/pkg/httpError"
)

// ParseToken func parses accessToken.
func ParseAccessToken(accessToken string) (int, error) {
	tokenStr, err := jwt.ParseWithClaims(accessToken, &jwtModel.AccessTokenClaims{}, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("invalid jwt signing method")
		}

		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})
	if err != nil {
		return 0, httpError.NewUnauthorizedError("invalid access token, can not parse it")
	}

	claims, ok := tokenStr.Claims.(*jwtModel.AccessTokenClaims)
	if !ok {
		return 0, errors.New("invalid jwt claims")
	}

	return claims.UserID, nil
}
