package helpers

import (
	"errors"
	"os"

	"github.com/golang-jwt/jwt/v5"
	jwtModel "github.com/jumayevgadam/evernote-go/internal/models/jwt"
)

// ParseToken func parses accessToken.
func ParseAccessToken(accessToken string) (*jwtModel.AccessTokenClaims, error) {
	tokenStr, err := jwt.ParseWithClaims(accessToken, &jwtModel.AccessTokenClaims{}, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("invalid jwt signing method")
		}

		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := tokenStr.Claims.(*jwtModel.AccessTokenClaims)
	if !ok {
		return nil, errors.New("invalid jwt claims")
	}

	return claims, nil
}
