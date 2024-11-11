package safesql

import (
	"context"
	"database/sql"
)

type DB struct {
	db *sql.DB
}

type Rows = sql.Rows
type Result = sql.Result

func (db *DB) QueryContext(ctx context.Context, query TrustedSQL, args ...any) (*Rows, error) {
	r, err := db.db.QueryContext(ctx, query.s, args...)

	return r, err
}

func (db *DB) ExecContext(ctx context.Context, query TrustedSQL, args ...any) (Result, error) {
	r, err := db.db.ExecContext(ctx, query.s, args...)

	return r, err
}

func Open(driverName, dataSourceName TrustedSQL) (*sql.DB, error) {
	return sql.Open(driverName.s, dataSourceName.s)
}

func NewDB(db *sql.DB) *DB {
	return &DB{db}
}
