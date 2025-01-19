package users

import (
	"context"

	jwtModel "github.com/jumayevgadam/evernote-go/internal/models/jwt"
	userModel "github.com/jumayevgadam/evernote-go/internal/models/user"
)

// Service is a service for managing users.
type Service interface {
	SignUp(ctx context.Context, req userModel.SignUpReq) (int, error)
	Login(ctx context.Context, req userModel.LoginReq) (jwtModel.Tokens, error)
}
