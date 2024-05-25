package repository

import (
	"context"
	"database/sql"

	"github.com/Lucifer07/e-wallet/entity"
	"github.com/Lucifer07/e-wallet/util"
)

type PasswordTokenRepositoryDb struct {
	db *sql.DB
}

type PasswordTokenRepository interface {
	CreateToken(ctx context.Context, id int) (*entity.PasswordToken, error)
	CheckToken(ctx context.Context, userId int) error
	DeleteToken(ctx context.Context, userId int) error
	ValidateToken(ctx context.Context, token string) (*entity.PasswordToken, error)
}

func NewPasswordTokenRepository(db *sql.DB) *PasswordTokenRepositoryDb {
	return &PasswordTokenRepositoryDb{
		db: db,
	}
}
func (r *PasswordTokenRepositoryDb) CreateToken(ctx context.Context, id int) (*entity.PasswordToken, error) {
	db := util.GetQueryRunner(ctx, r.db)
	ranToken, _ := util.RandomString(16)
	var passwordToken entity.PasswordToken
	statment := `insert into password_tokens(password_token,user_id)values($1,$2) returning password_token,expired_at;`
	err := db.QueryRowContext(ctx, statment, ranToken, id).Scan(&passwordToken.Token, &passwordToken.ExpiredAt)
	if err != nil {
		return nil, err
	}
	return &passwordToken, nil
}

func (r *PasswordTokenRepositoryDb) CheckToken(ctx context.Context, userId int) error {
	db := util.GetQueryRunner(ctx, r.db)
	statment := `select id from password_tokens where user_id=$1 limit 1;`
	_, err := db.ExecContext(ctx, statment, userId)
	if err != nil {
		return err
	}
	return nil
}
func (r *PasswordTokenRepositoryDb) DeleteToken(ctx context.Context, userId int) error {
	db := util.GetQueryRunner(ctx, r.db)
	statment := `delete from password_tokens where user_id=$1;`
	_, err := db.ExecContext(ctx, statment, userId)
	if err != nil {
		return err
	}
	return nil
}
func (r *PasswordTokenRepositoryDb) ValidateToken(ctx context.Context, token string) (*entity.PasswordToken, error) {
	var passwordToken entity.PasswordToken
	db := util.GetQueryRunner(ctx, r.db)
	statment := `select user_id,expired_at from password_tokens where password_token=$1;`
	err := db.QueryRowContext(ctx, statment, token).Scan(&passwordToken.UserId, &passwordToken.ExpiredAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, util.ErrorInvalidToken
		}
		return nil, err
	}
	return &passwordToken, nil
}
