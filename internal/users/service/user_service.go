package service

import (
	"context"

	"github.com/jumayevgadam/evernote-go/internal/database"
	"github.com/jumayevgadam/evernote-go/internal/helpers"
	jwtModel "github.com/jumayevgadam/evernote-go/internal/models/jwt"
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

// Login func returns access and refresh token.
func (s *UserService) Login(ctx context.Context, loginReq userModel.LoginReq) (jwtModel.Tokens, error) {
	// get userdetails.
	user, err := s.repo.UsersRepo().GetUserByEmail(ctx, loginReq.Email)
	if err != nil {
		return jwtModel.Tokens{}, err
	}

	// compare hashed password with loginReq.Password.
	if err := utils.CheckAndComparePassword(loginReq.Password, user.Password); err != nil {
		return jwtModel.Tokens{}, err
	}

	// accessToken generation.
	accessToken, err := helpers.GenerateAccessToken(user.Username, user.Email, user.ID)
	if err != nil {
		return jwtModel.Tokens{}, err
	}

	// refreshToken generation.
	refreshToken, err := helpers.GenerateRefreshToken(user.ID)
	if err != nil {
		return jwtModel.Tokens{}, err
	}

	return jwtModel.Tokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
