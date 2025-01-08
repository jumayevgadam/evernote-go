package database

import (
	"context"

	"github.com/jumayevgadam/evernote-go/internal/users"
)

type Transaction func(db DataStore) error

// DataStore interface keeps all needed methods for database operations.
type DataStore interface {
	WithTransaction(ctx context.Context, transactionFn Transaction) error
	UsersRepo() users.Repository
}
