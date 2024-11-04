package sqlite

import (
	"context"
	"database/sql"
	"github.com/failuretoload/datamonster/store"
	_ "modernc.org/sqlite"
)

// rowAdapter wraps sql.Row to implement store.Row
type rowAdapter struct {
	row *sql.Row
}

func (r *rowAdapter) Scan(dest ...any) error {
	return r.row.Scan(dest...)
}

// rowsAdapter wraps sql.Rows to implement store.Rows
type rowsAdapter struct {
	rows *sql.Rows
}

func (r *rowsAdapter) Close() {
	_ = r.rows.Close()
}

func (r *rowsAdapter) Err() error {
	return r.rows.Err()
}

func (r *rowsAdapter) Next() bool {
	return r.rows.Next()
}

func (r *rowsAdapter) Scan(dest ...any) error {
	return r.rows.Scan(dest...)
}

// DAO implements store.ConnectionV2 for DAO
type DAO struct {
	db *sql.DB
}

// NewDAO creates a new in-memory DAO database
func NewDAO() (*DAO, error) {
	db, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		return nil, err
	}

	// Create the settlement table
	_, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS settlement (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            owner TEXT NOT NULL,
            name TEXT NOT NULL,
            survival_limit INTEGER NOT NULL,
            departing_survival INTEGER NOT NULL,
            collective_cognition INTEGER NOT NULL,
            year INTEGER NOT NULL
        )
    `)
	if err != nil {
		_ = db.Close()
		return nil, err
	}

	return &DAO{db: db}, nil
}

func (s *DAO) Query(ctx context.Context, query string, args ...interface{}) (store.Rows, error) {
	rows, err := s.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	return &rowsAdapter{rows: rows}, nil
}

func (s *DAO) QueryRow(ctx context.Context, query string, args ...interface{}) store.Row {
	return &rowAdapter{row: s.db.QueryRowContext(ctx, query, args...)}
}

func (s *DAO) Exec(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return s.db.ExecContext(ctx, query, args...)
}

// Close closes the database connection
func (s *DAO) Close() error {
	return s.db.Close()
}

// Ensure DAO implements store.ConnectionV2
var _ store.DAO = (*DAO)(nil)
