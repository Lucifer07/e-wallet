package util

import (
	"context"
	"database/sql"
)

type QueryRunner interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
	QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row
}
type txKey struct{}

func InjectTx(ctx context.Context, tx *sql.Tx) context.Context {
	return context.WithValue(ctx, txKey{}, tx)
}
func ExtractTx(ctx context.Context) *sql.Tx {
	if tx, ok := ctx.Value(txKey{}).(*sql.Tx); ok {
		return tx
	}
	return nil
}
func GetQueryRunner(ctx context.Context, db QueryRunner) QueryRunner {
	tx := ExtractTx(ctx)
	if tx != nil {
		return tx
	}
	return db
}
