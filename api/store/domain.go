package store

import (
	"context"
	"database/sql"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type Rows interface {
	Close()
	Err() error
	Next() bool
	Scan(dest ...any) error
}

// Row defines a single-row result interface
type Row interface {
	Scan(dest ...any) error
}

type Connection interface {
	Close()
	Begin(ctx context.Context) (pgx.Tx, error)
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, optionsAndArgs ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, optionsAndArgs ...interface{}) pgx.Row
}

// ConnectionV2 defines the common database operations
type DAO interface {
	Query(ctx context.Context, query string, args ...interface{}) (Rows, error)
	QueryRow(ctx context.Context, query string, args ...interface{}) Row
	Exec(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	Close() error
}
