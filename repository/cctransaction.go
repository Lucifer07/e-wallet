package repository

import (
	"context"
	"database/sql"

	"github.com/Lucifer07/e-wallet/entity"
	"github.com/Lucifer07/e-wallet/util"
)

type CCTransactionRepositoryDb struct {
	db *sql.DB
}

type CCTransactionRepository interface {
	CreateCCAccount(ctx context.Context, ccNumber int) (*entity.CreditCardAccount, error)
}

func NewCCTransactionRepository(db *sql.DB) *CCTransactionRepositoryDb {
	return &CCTransactionRepositoryDb{
		db: db,
	}
}

func (r *CCTransactionRepositoryDb) CreateCCAccount(ctx context.Context, ccNumber int) (*entity.CreditCardAccount, error) {
	var CCTransaction entity.CreditCardAccount
	db := util.GetQueryRunner(ctx, r.db)
	statment := `insert into credit_card_account(cc_number)values($1) returning id,cc_number;`
	err := db.QueryRowContext(ctx, statment, ccNumber).Scan(&CCTransaction.Id, &CCTransaction.CCNumber)
	if err != nil {
		return nil, err
	}
	return &CCTransaction, nil
}
