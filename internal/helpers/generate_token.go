package helpers

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	jwtModel "github.com/jumayevgadam/evernote-go/internal/models/jwt"
	"github.com/jumayevgadam/evernote-go/pkg/constants"
)

// GenerateAccessToken.
func GenerateAccessToken(username, email string, userID int) (string, error) {
	claims := jwtModel.AccessTokenClaims{
		Username: username,
		Email:    email,
		UserID:   userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(constants.AccessTokenExpiryTime)),
		},
	}

	tokenStr, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	if err != nil {
		return "", fmt.Errorf("can not generate access token: %w", err)
	}

	return tokenStr, nil
}

// GenerateRefreshToken.
func GenerateRefreshToken(userID int) (string, error) {
	claims := jwtModel.RefreshTokenClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(constants.RefreshTokenExpiryTime)),
		},
	}

	tokenStr, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	if err != nil {
		return "", fmt.Errorf("can not generate refresh token: %w", err)
	}

	return tokenStr, nil
}
