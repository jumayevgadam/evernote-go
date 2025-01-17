package connection

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jumayevgadam/evernote-go/internal/config"
)

var _ DB = (*Database)(nil)

// Querier interface.
type Querier interface {
	QueryRow(ctx context.Context, query string, args ...interface{}) pgx.Row
	Query(ctx context.Context, query string, args ...interface{}) (pgx.Rows, error)
	Exec(ctx context.Context, query string, args ...interface{}) (pgconn.CommandTag, error)
}

// DB interface keeps needed methods for psqlDB.
type DB interface {
	Get(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	Select(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	Querier
}

// DBOps interface with Transaction method.
type DBOps interface {
	DB
	Begin(ctx context.Context, txOpts pgx.TxOptions) (TxOps, error)
	Close()
}

// Database struct keeps pgxpool.
type Database struct {
	Db *pgxpool.Pool
}

func GetDBConnection(ctx context.Context, cfg config.PostgresDB) (*Database, error) {
	const (
		retryAttempts = 3
		retryDelay    = 2 * time.Second
	)

	var (
		db  *pgxpool.Pool
		err error
	)

	for i := 0; i < retryAttempts; i++ {
		db, err = connectToDB(ctx, cfg)
		if err == nil {
			return &Database{Db: db}, nil
		}

		time.Sleep(retryDelay)
	}

	// After three times, if db connection failed, then throw fatal.
	log.Fatalf("failed to connect to db after %d attempts: %v", retryAttempts, err)

	return nil, fmt.Errorf("failed to connect to db after %d attempts: %w", retryAttempts, err)
}

// connectToDB func connects to db.
func connectToDB(ctx context.Context, cfg config.PostgresDB) (*pgxpool.Pool, error) {
	const (
		maxConnections = 200
		minConnections = 10
		// maxConnLifetime   = 1 * time.Hour
		// maxConnIdleTime   = 30 * time.Minute
		// healthCheckPeriod = 1 * time.Minute
	)

	hostPort := net.JoinHostPort(cfg.Host, cfg.Port)
	connStr := fmt.Sprintf(
		"postgres://%s:%s@%s/%s?sslmode=%s",
		cfg.User,
		cfg.Password,
		hostPort,
		cfg.DBName,
		cfg.SslMode,
	)

	config, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		return nil, fmt.Errorf("parsing connection config: %w", err)
	}

	// Configure the connection pool settings.
	config.MaxConns = maxConnections
	config.MinConns = minConnections
	// config.MaxConnLifetime = maxConnLifetime
	// config.MaxConnIdleTime = maxConnIdleTime
	// config.HealthCheckPeriod = healthCheckPeriod

	// create a new connection pool.
	db, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return nil, fmt.Errorf("creating connection pool: %w", err)
	}

	// ping the database to verify the connection.
	if err := db.Ping(ctx); err != nil {
		return nil, fmt.Errorf("pinging database error: %w", err)
	}

	return db, nil
}

// Get func is used for getting single row.
func (d *Database) Get(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return pgxscan.Get(ctx, d.Db, dest, query, args...)
}

// Select func is used for getting multiple rows.
func (d *Database) Select(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return pgxscan.Select(ctx, d.Db, dest, query, args...)
}

// QueryRow func is used for querying single row.
func (d *Database) QueryRow(ctx context.Context, query string, args ...interface{}) pgx.Row {
	return d.Db.QueryRow(ctx, query, args...)
}

// Query func is used for querying multiple rows.
func (d *Database) Query(ctx context.Context, query string, args ...interface{}) (pgx.Rows, error) {
	return d.Db.Query(ctx, query, args...)
}

// Exec func is used for executing query.
func (d *Database) Exec(ctx context.Context, query string, args ...interface{}) (pgconn.CommandTag, error) {
	result, err := d.Db.Exec(ctx, query, args...)
	if err != nil {
		return pgconn.CommandTag{}, fmt.Errorf("executing query: %w", err)
	}

	if result.RowsAffected() == 0 {
		return pgconn.CommandTag{}, pgx.ErrNoRows
	}

	return result, nil
}

// Begin func is used for starting transaction.
func (d *Database) Begin(ctx context.Context, txOpts pgx.TxOptions) (TxOps, error) {
	if d == nil || d.Db == nil {
		return nil, fmt.Errorf("database connection is nil")
	}

	tx, err := d.Db.BeginTx(ctx, txOpts)
	if err != nil {
		return nil, fmt.Errorf("connection begin transaction: %w", err)
	}

	return &Transaction{Tx: tx}, nil
}

// Close func is used for closing db connection.
func (d *Database) Close() {
	d.Db.Close()
}
