package repository

import (
	"context"
	"database/sql"

	"github.com/Lucifer07/e-wallet/entity"
	"github.com/Lucifer07/e-wallet/util"
)

type WalletRepositoryDb struct {
	db *sql.DB
}

type WalletRepository interface {
	CreateWallet(ctx context.Context, userId int) (*entity.Wallet, error)
	GetWalletId(ctx context.Context, userId int) (*int, error)
	UpdateBalance(ctx context.Context, wallet entity.Wallet) error
	GetWalletData(ctx context.Context, userId int) (*entity.Wallet, error)
	GetPatnerWallet(ctx context.Context, walletNumber string) (*entity.Wallet, error)
}

func NewWalletRepository(db *sql.DB) *WalletRepositoryDb {
	return &WalletRepositoryDb{
		db: db,
	}
}
func (r *WalletRepositoryDb) CreateWallet(ctx context.Context, userId int) (*entity.Wallet, error) {
	var wallet entity.Wallet
	db := util.GetQueryRunner(ctx, r.db)

	statment := `insert into wallets(user_id)values($1) returning wallet_number,balance,created_at;`
	err := db.QueryRowContext(ctx, statment, userId).Scan(&wallet.WalletNumber, &wallet.Balance, &wallet.CreatedAt)
	if err != nil {
		return nil, err
	}
	wallet.UserId = userId
	return &wallet, nil
}
func (r *WalletRepositoryDb) GetWalletId(ctx context.Context, userId int) (*int, error) {
	var walletId int
	db := util.GetQueryRunner(ctx, r.db)
	statment := `select id from wallets where user_id=$1 and deleted_at is null;`
	err := db.QueryRowContext(ctx, statment, userId).Scan(&walletId)
	if err != nil {
		return nil, err
	}
	return &walletId, nil
}
func (r *WalletRepositoryDb) UpdateBalance(ctx context.Context, wallet entity.Wallet) error {
	db := util.GetQueryRunner(ctx, r.db)
	statment := `update wallets set balance=$1,updated_at=now() where user_id=$2;`
	_, err := db.ExecContext(ctx, statment, wallet.Balance, wallet.UserId)
	if err != nil {
		return err
	}
	return nil
}
func (r *WalletRepositoryDb) GetWalletData(ctx context.Context, userId int) (*entity.Wallet, error) {
	var wallet entity.Wallet
	db := util.GetQueryRunner(ctx, r.db)
	statment := `select id,balance,user_id,wallet_number from wallets where user_id=$1 and deleted_at is null for update;`
	err := db.QueryRowContext(ctx, statment, userId).Scan(&wallet.Id, &wallet.Balance, &wallet.UserId, &wallet.WalletNumber)
	if err != nil {
		return nil, err
	}
	return &wallet, nil
}
func (r *WalletRepositoryDb) GetPatnerWallet(ctx context.Context, walletNumber string) (*entity.Wallet, error) {
	var wallet entity.Wallet
	db := util.GetQueryRunner(ctx, r.db)
	statment := `select id,balance,user_id,wallet_number from wallets where wallet_number=$1 and deleted_at is null for update;`
	err := db.QueryRowContext(ctx, statment, walletNumber).Scan(&wallet.Id, &wallet.Balance, &wallet.UserId, &wallet.WalletNumber)
	if err != nil {
		return nil, err
	}
	return &wallet, nil
}
