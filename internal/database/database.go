package database

import (
	"context"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Database interface {
	// Exec queries the database and returns the affected rows
	Exec(sql string, args ...interface{}) (int64, error)
	// Query queries the database and return the rows
	Query(sql string, args ...interface{}) (Rows, error)
	// QueryRow queries the database and return a single row
	QueryRow(sql string, args ...interface{}) Row
	// BeginTx Starts a database Transaction
	BeginTx() (Transaction, error)
}

type database struct {
	ctx  context.Context
	conn *pgxpool.Pool
}

// Row is the result returned from a query
type Row interface {
	// Scan reads the values from the current row into dest values positionally
	Scan(dest ...interface{}) error
}

// Rows is the result set returned from a query
type Rows interface {
	// Scan reads the values from the current row into dest values positionally
	Scan(dest ...interface{}) error
	// Next prepares the next row for reading. It returns true if there is another row and false if no more rows are available. It automatically closes rows when all rows are read.
	Next() bool
	// Close closes the rows, making the connection ready for use again. It is safe to call Close after rows is already closed.
	Close()
	// Err returns any error that occurred while reading.
	Err() error
}

// Transaction represents an SQL database transaction
type Transaction interface {
	// Commit commits the database transaction
	Commit() error
	// Rollback rollbacks the database transaction
	Rollback() error
	// Exec queries the database and returns the affected rows
	Exec(sql string, args ...interface{}) (int64, error)
	// Query queries the database and return the rows
	Query(sql string, args ...interface{}) (Rows, error)
	// QueryRow queries the database and return a single row
	QueryRow(sql string, args ...interface{}) Row
}

// New creates a new postgresql database instance
func New(ctx context.Context, connString string) (Database, error) {
	cfg, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return nil, err
	}

	conn, err := pgxpool.ConnectConfig(ctx, cfg)
	if err != nil {
		return nil, err
	}

	return &database{
		ctx:  ctx,
		conn: conn,
	}, nil
}

func (p *database) Exec(sql string, args ...interface{}) (int64, error) {
	result, err := p.conn.Exec(p.ctx, sql, args...)
	return result.RowsAffected(), err
}

func (p *database) Query(sql string, args ...interface{}) (Rows, error) {
	rows, err := p.conn.Query(p.ctx, sql, args...)
	return newDatabaseRows(rows), err
}

func (p *database) QueryRow(sql string, args ...interface{}) Row {
	row := p.conn.QueryRow(p.ctx, sql, args...)
	return newDatabaseRow(row)
}

func (p *database) BeginTx() (Transaction, error) {
	tx, err := p.conn.BeginTx(p.ctx, pgx.TxOptions{})
	if err != nil {
		return nil, err
	}

	return newDatabaseTransaction(p.ctx, tx), nil
}

type databaseRow struct {
	row pgx.Row
}

func newDatabaseRow(row pgx.Row) Row {
	return &databaseRow{
		row: row,
	}
}

func (p *databaseRow) Scan(dest ...interface{}) error {
	return p.row.Scan(dest...)
}

type databaseRows struct {
	rows pgx.Rows
}

func newDatabaseRows(rows pgx.Rows) Rows {
	return &databaseRows{
		rows: rows,
	}
}

func (p *databaseRows) Scan(dest ...interface{}) error {
	return p.rows.Scan(dest...)
}

func (p *databaseRows) Next() bool {
	return p.rows.Next()
}

func (p *databaseRows) Close() {
	p.rows.Close()
}

func (p *databaseRows) Err() error {
	return p.rows.Err()
}

type transaction struct {
	ctx context.Context
	tx  pgx.Tx
}

func newDatabaseTransaction(ctx context.Context, tx pgx.Tx) *transaction {
	return &transaction{
		ctx: ctx,
		tx:  tx,
	}
}

func (p *transaction) Commit() error {
	return p.tx.Commit(p.ctx)
}

func (p *transaction) Rollback() error {
	return p.tx.Rollback(p.ctx)
}

func (p *transaction) Exec(sql string, args ...interface{}) (int64, error) {
	result, err := p.tx.Exec(p.ctx, sql, args...)
	return result.RowsAffected(), err
}

func (p *transaction) Query(sql string, args ...interface{}) (Rows, error) {
	rows, err := p.tx.Query(p.ctx, sql, args...)
	return newDatabaseRows(rows), err
}

func (p *transaction) QueryRow(sql string, args ...interface{}) Row {
	row := p.tx.QueryRow(p.ctx, sql, args...)
	return newDatabaseRow(row)
}
