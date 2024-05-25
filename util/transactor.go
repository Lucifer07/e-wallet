package util

import (
	"context"
	"database/sql"
)

type Transactor interface {
	WithinTransaction(ctx context.Context, fn func(ctx context.Context) error) error
}
type transactorImpl struct {
	db *sql.DB
}

func NewTransactor(db *sql.DB) *transactorImpl {
	return &transactorImpl{
		db: db,
	}
}
func (t *transactorImpl) WithinTransaction(ctx context.Context, fn func(ctx context.Context) error) error {
	tx, err := t.db.Begin()
	if err != nil {
		return err
	}
	err = fn(InjectTx(ctx, tx))
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}
