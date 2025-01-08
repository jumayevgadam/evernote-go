package postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jumayevgadam/evernote-go/internal/connection"
	"github.com/jumayevgadam/evernote-go/internal/database"
	"go.uber.org/zap"
)

// WithTransaction starts a new transaction and executes the transactionFn.
func (d *DataStore) WithTransaction(ctx context.Context, transactionFn database.Transaction) error {
	db, ok := d.db.(connection.DBOps)
	if !ok {
		return errors.New("invalid database connection and err in type assertion")
	}

	// begin transaction.
	tx, err := db.Begin(ctx, pgx.TxOptions{})
	if err != nil {
		return fmt.Errorf("error beginning transaction: %w", err)
	}

	defer func() {
		if err != nil {
			if rbErr := tx.Rollback(ctx); rbErr != nil {
				zap.L().Error("error rolling back transaction", zap.Error(rbErr))
			} else {
				zap.L().Info("transaction rolled back")
			}
		}
	}()

	transactionalDB := &DataStore{db: tx}

	// execute transactionFn.
	err = transactionFn(transactionalDB)
	if err != nil {
		return fmt.Errorf("error executing transaction: %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("error committing transaction: %w", err)
	}

	return nil
}
