package service

import (
	"context"
	"time"

	"github.com/Lucifer07/e-wallet/dto"
	"github.com/Lucifer07/e-wallet/entity"
	"github.com/Lucifer07/e-wallet/payload"
	"github.com/Lucifer07/e-wallet/repository"
	"github.com/Lucifer07/e-wallet/util"
)

type PasswordTokenServiceImp struct {
	PasswordTokenRepo repository.PasswordTokenRepository
	transaction       util.Transactor
	userRepo          repository.UserRepository
	helperTool        util.HelperInf
}
type PasswordTokenService interface {
	CreateResetPassword(ctx context.Context, email dto.GettokenRequest) (*dto.GetTokenResponse, error)
	ResetPassword(ctx context.Context, tokenPassword dto.TokenPassword) error
}

func NewPasswordTokenService(PasswordTokenRepo repository.PasswordTokenRepository, transaction util.Transactor, userRepo repository.UserRepository, helperTool util.HelperInf) *PasswordTokenServiceImp {
	return &PasswordTokenServiceImp{
		PasswordTokenRepo: PasswordTokenRepo,
		transaction:       transaction,
		userRepo:          userRepo,
		helperTool:        helperTool,
	}
}
func (s *PasswordTokenServiceImp) CreateResetPassword(ctx context.Context, email dto.GettokenRequest) (*dto.GetTokenResponse, error) {
	var tokenPass entity.PasswordToken
	err := s.transaction.WithinTransaction(ctx, func(ctx context.Context) error {
		userId, err := s.userRepo.CheckEmail(ctx, email.Email)
		if err != nil {
			return err
		}
		if userId != nil {
			err := s.PasswordTokenRepo.CheckToken(ctx, *userId)
			if err != nil {
				return err
			}
			err = s.PasswordTokenRepo.DeleteToken(ctx, *userId)
			if err != nil {
				return err
			}
			tokenPassword, err := s.PasswordTokenRepo.CreateToken(ctx, *userId)
			if err != nil {
				return err
			}
			tokenPass = *tokenPassword
			return nil
		}
		return util.ErrorUserNotFound
	})
	if err != nil {
		return nil, err
	}
	responseData := payload.PasswordTokenToResponse(tokenPass)
	return &responseData, nil
}
func (s *PasswordTokenServiceImp) ResetPassword(ctx context.Context, tokenPassword dto.TokenPassword) error {
	err := s.transaction.WithinTransaction(ctx, func(ctx context.Context) error {
		passwordReset, err := s.PasswordTokenRepo.ValidateToken(ctx, tokenPassword.Token)
		if err != nil {
			return err
		}
		if passwordReset.ExpiredAt.Before(time.Now()) {
			return util.ErrorTokenExp
		}
		passwordHash, err := s.helperTool.HashPassword(tokenPassword.Password, 12)
		if err != nil {
			return err
		}
		password := string(passwordHash)
		err = s.userRepo.UpdatePassword(ctx, password, passwordReset.UserId)
		if err != nil {
			return err
		}
		err = s.PasswordTokenRepo.DeleteToken(ctx, passwordReset.UserId)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}
