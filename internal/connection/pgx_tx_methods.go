package connection

import (
	"context"
	"fmt"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

// TxOps interface with Commit and Rollback methods.
type TxOps interface {
	DB
	Commit(ctx context.Context) error
	Rollback(ctx context.Context) error
}

// Transaction struct keeps pgx.Tx.
type Transaction struct {
	Tx pgx.Tx
}

// Get returns a row from the database.
func (tx *Transaction) Get(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return pgxscan.Get(ctx, tx.Tx, dest, query, args...)
}

// Select returns rows from the database.
func (tx *Transaction) Select(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return pgxscan.Select(ctx, tx.Tx, dest, query, args...)
}

// QueryRow returns a row from the database.
func (tx *Transaction) QueryRow(ctx context.Context, query string, args ...interface{}) pgx.Row {
	return tx.Tx.QueryRow(ctx, query, args...)
}

// Query returns rows from the database.
func (tx *Transaction) Query(ctx context.Context, query string, args ...interface{}) (pgx.Rows, error) {
	return tx.Tx.Query(ctx, query, args...)
}

// Exec executes a query in the database.
func (tx *Transaction) Exec(ctx context.Context, query string, args ...interface{}) (pgconn.CommandTag, error) {
	result, err := tx.Tx.Exec(ctx, query, args...)
	if err != nil {
		return pgconn.CommandTag{}, fmt.Errorf("executing query: %w", err)
	}

	if result.RowsAffected() == 0 {
		return pgconn.CommandTag{}, pgx.ErrNoRows
	}

	return result, nil
}

// Commit commits the transaction.
func (tx *Transaction) Commit(ctx context.Context) error {
	return tx.Tx.Commit(ctx)
}

// Rollback rolls back the transaction.
func (tx *Transaction) Rollback(ctx context.Context) error {
	return tx.Tx.Rollback(ctx)
}
