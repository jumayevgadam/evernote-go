package jwt

import "github.com/golang-jwt/jwt/v5"

type AccessTokenClaims struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	UserID   int    `json:"userID"`
	jwt.RegisteredClaims
}

type RefreshTokenClaims struct {
	UserID int `json:"userID"`
	jwt.RegisteredClaims
}
