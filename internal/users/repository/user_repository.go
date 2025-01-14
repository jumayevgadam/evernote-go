package repository

import (
	"context"

	"github.com/jumayevgadam/evernote-go/internal/connection"
	userModel "github.com/jumayevgadam/evernote-go/internal/models/user"
	"github.com/jumayevgadam/evernote-go/internal/users"
)

var _ users.Repository = (*UserRepository)(nil)

// UserRepository struct.
type UserRepository struct {
	psqlDB connection.DB
}

// NewUserRepository returns a new UserRepository.
func NewUserRepository(db connection.DB) *UserRepository {
	return &UserRepository{
		psqlDB: db,
	}
}

// SignUp method for user repository.
func (r *UserRepository) SignUp(ctx context.Context, data userModel.SignUpReqData) (int, error) {
	var userID int

	err := r.psqlDB.QueryRow(
		ctx,
		signUpQuery,
		data.Username,
		data.Password,
		data.Email,
	).Scan(&userID)

	if err != nil {
		return 0, err
	}

	return userID, nil
}

func (r *UserRepository) GetUserByEmail(ctx context.Context, email string) (userModel.AllUserData, error) {
	var userData userModel.AllUserData

	err := r.psqlDB.Get(
		ctx,
		&userData,
		getUserByEmailQuery,
		email,
	)

	if err != nil {
		return userModel.AllUserData{}, err
	}

	return userData, nil
}
