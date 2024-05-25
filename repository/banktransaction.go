package repository

import (
	"context"
	"database/sql"

	"github.com/Lucifer07/e-wallet/entity"
	"github.com/Lucifer07/e-wallet/util"
)

type BankTransactionRepositoryDb struct {
	db *sql.DB
}

type BankTransactionRepository interface {
	CreateBankAccount(ctx context.Context, accountNumber int) (*entity.BankAccount, error)
}

func NewBankTransactionRepository(db *sql.DB) *BankTransactionRepositoryDb {
	return &BankTransactionRepositoryDb{
		db: db,
	}
}

func (r *BankTransactionRepositoryDb) CreateBankAccount(ctx context.Context, accountNumber int) (*entity.BankAccount, error) {
	var BankTransaction entity.BankAccount
	db := util.GetQueryRunner(ctx, r.db)
	statment := `insert into bank_account(account_number)values($1) returning id,account_number;`
	err := db.QueryRowContext(ctx, statment, accountNumber).Scan(&BankTransaction.Id, &BankTransaction.AccountNumber)
	if err != nil {
		return nil, err
	}
	return &BankTransaction, nil
}
