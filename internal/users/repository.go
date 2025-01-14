package users

import (
	"context"

	userModel "github.com/jumayevgadam/evernote-go/internal/models/user"
)

// Repository interface for users.
type Repository interface {
	SignUp(ctx context.Context, data userModel.SignUpReqData) (int, error)
	GetUserByEmail(ctx context.Context, email string) (userModel.AllUserData, error)
}
