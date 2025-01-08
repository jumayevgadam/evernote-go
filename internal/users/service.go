package users

import (
	"context"

	userModel "github.com/jumayevgadam/evernote-go/internal/models/user"
)

// Service is a service for managing users.
type Service interface {
	SignUp(ctx context.Context, req userModel.SignUpReq) (int, error)
}
