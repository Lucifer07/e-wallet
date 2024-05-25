package repository

import (
	"context"
	"database/sql"

	"github.com/Lucifer07/e-wallet/entity"
	"github.com/Lucifer07/e-wallet/util"
)

type PayLaterTransactionRepositoryDb struct {
	db *sql.DB
}

type PayLaterTransactionRepository interface {
	CreatePayLaterTransaction(ctx context.Context, userId int) (*entity.PaylaterAccount, error)
}

func NewPayLaterTransactionRepository(db *sql.DB) *PayLaterTransactionRepositoryDb {
	return &PayLaterTransactionRepositoryDb{
		db: db,
	}
}

func (r *PayLaterTransactionRepositoryDb) CreatePayLaterTransaction(ctx context.Context, userId int) (*entity.PaylaterAccount, error) {
	var PayLaterTransaction entity.PaylaterAccount
	db := util.GetQueryRunner(ctx, r.db)
	statment := `insert into pay_later_account(user_id)values($1) returning id,user_id;`
	err := db.QueryRowContext(ctx, statment, userId).Scan(&PayLaterTransaction.Id, &PayLaterTransaction.UserId)
	if err != nil {
		return nil, err
	}
	return &PayLaterTransaction, nil
}
