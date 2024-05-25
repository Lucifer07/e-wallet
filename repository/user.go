package repository

import (
	"context"
	"database/sql"

	"github.com/Lucifer07/e-wallet/dto"
	"github.com/Lucifer07/e-wallet/entity"
	"github.com/Lucifer07/e-wallet/util"
)

type UserRepositoryDb struct {
	db *sql.DB
}

type UserRepository interface {
	Register(ctx context.Context, user entity.User) (*entity.User, error)
	Login(ctx context.Context, user entity.User) (*entity.User, error)
	CheckEmail(ctx context.Context, emailRequest string) (*int, error)
	UpdatePassword(ctx context.Context, newPassword string, userId int) error
	GetSelf(ctx context.Context, userId int) (*dto.UserWalet, error)
	UpdateProfile(ctx context.Context, email string,fullname string, userId int) error 
	UpdateAvatar(ctx context.Context, avatar string,userId int) error 
}

func NewUserRepository(db *sql.DB) *UserRepositoryDb {
	return &UserRepositoryDb{
		db: db,
	}
}

func (r *UserRepositoryDb) Register(ctx context.Context, user entity.User) (*entity.User, error) {
	var userData entity.User
	db := util.GetQueryRunner(ctx, r.db)
	statment := `insert into users(name,email,password)values($1,$2,$3) returning id,name,email,created_at;`
	err := db.QueryRowContext(ctx, statment, user.Name, user.Email, user.Password).Scan(&userData.Id, &user.Name, &userData.Email, &userData.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &userData, nil
}
func (r *UserRepositoryDb) Login(ctx context.Context, user entity.User) (*entity.User, error) {
	var userData entity.User
	statment := `select id,email,password from users where email=$1;`
	err := r.db.QueryRowContext(ctx, statment, user.Email).Scan(&userData.Id, &userData.Email, &userData.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &userData, nil
}
func (r *UserRepositoryDb) CheckEmail(ctx context.Context, emailRequest string) (*int, error) {
	var id int
	db := util.GetQueryRunner(ctx, r.db)
	statment := `select id from users where email = $1 and deleted_at is null;`
	err := db.QueryRowContext(ctx, statment, emailRequest).Scan(&id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &id, nil
}
func (r *UserRepositoryDb) UpdatePassword(ctx context.Context, newPassword string, userId int) error {
	db := util.GetQueryRunner(ctx, r.db)
	statment := `update users set password=$1,updated_at=now() where id=$2;`
	_, err := db.ExecContext(ctx, statment, newPassword, userId)
	if err != nil {
		return err
	}
	return nil
}
func (r *UserRepositoryDb) GetSelf(ctx context.Context, userId int) (*dto.UserWalet, error) {
	var data dto.UserWalet
	db := util.GetQueryRunner(ctx, r.db)
	statment := `SELECT 
    u.id, u.name, u.avatar, u.email, u.created_at, u.updated_at, 
    w.id, w.wallet_number, w.balance, w.created_at AS wallet_created_at, w.updated_at AS wallet_updated_at,
    COALESCE((SELECT SUM(amount) FROM history_transactions WHERE (sender_wallet_id = w.id or sender_wallet_id = CAST(w.wallet_number AS bigint)) AND source_of_fund = 'wallet'), 0) AS expense,
    COALESCE((SELECT SUM(amount) FROM history_transactions WHERE (recipient_wallet_id = w.id or recipient_wallet_id = CAST(w.wallet_number AS bigint))), 0) AS income
FROM 
    users u 
JOIN 
    wallets w ON w.user_id = u.id 
WHERE 
    u.id = $1;

`
	err := db.QueryRowContext(ctx, statment, userId).Scan(&data.UserId, &data.Name,&data.Avatar ,&data.Email, &data.UserCreatedAt, &data.UserUpdatedAt, &data.WalletId, &data.WalletNumber, &data.Balance, &data.WalletCreatedAt, &data.WalletUpdatedAt, &data.Expense, &data.Income)
	if err != nil {
		if err==sql.ErrNoRows {
			return nil,nil
		}
		return nil, err
	}
	return &data, nil
}
func (r *UserRepositoryDb) UpdateProfile(ctx context.Context, email string,fullname string, userId int) error {
	db := util.GetQueryRunner(ctx, r.db)
	statment := `update users set email=$1, name=$2,updated_at=now() where id=$3;`
	_, err := db.ExecContext(ctx, statment, email, fullname,userId)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepositoryDb) UpdateAvatar(ctx context.Context, avatar string,userId int) error {
	db := util.GetQueryRunner(ctx, r.db)
	statment := `update users set avatar=$1, updated_at=now() where id=$2;`
	_, err := db.ExecContext(ctx, statment, avatar,userId)
	if err != nil {
		return err
	}
	return nil
}