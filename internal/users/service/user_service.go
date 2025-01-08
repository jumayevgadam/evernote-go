package service

import (
	"context"

	"github.com/jumayevgadam/evernote-go/internal/database"
	userModel "github.com/jumayevgadam/evernote-go/internal/models/user"
	"github.com/jumayevgadam/evernote-go/internal/users"
	"github.com/jumayevgadam/evernote-go/pkg/utils"
)

var _ users.Service = (*UserService)(nil)

// UserService is a service for managing users.
type UserService struct {
	repo database.DataStore
}

// NewUserService returns a new UserService.
func NewUserService(repo database.DataStore) *UserService {
	return &UserService{
		repo: repo,
	}
}

// SignUp method creates a new user.
func (s *UserService) SignUp(ctx context.Context, req userModel.SignUpReq) (int, error) {
	hashedPass, err := utils.HashPassword(req.Password)
	if err != nil {
		return 0, err
	}
	req.Password = hashedPass
	
	userID, err := s.repo.UsersRepo().SignUp(ctx, req.ToPsqlDBStorage())
	if err != nil {
		return 0, err
	}

	return userID, nil
}
