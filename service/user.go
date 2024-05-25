package service

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/Lucifer07/e-wallet/dto"
	"github.com/Lucifer07/e-wallet/payload"
	"github.com/Lucifer07/e-wallet/repository"
	"github.com/Lucifer07/e-wallet/util"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserServiceImp struct {
	userRepo    repository.UserRepository
	walletRepo  repository.WalletRepository
	transaction util.Transactor
	helperTool  util.HelperInf
}
type UserService interface {
	Register(ctx context.Context, registerData dto.RegisterRequest) (*dto.RegisterResponse, error)
	Login(ctx context.Context, loginData dto.Login) (*string, error)
	GetSelf(ctx context.Context) (*dto.UserWalletResponse, error)
	UpdateProfile(ctx context.Context, dataUser dto.UpdateProfile) error
	UpdateAvatar(ctx *gin.Context) error
}

func NewUserService(userRepo repository.UserRepository, helperRepo util.HelperInf, walletRepo repository.WalletRepository, transaction util.Transactor) *UserServiceImp {
	return &UserServiceImp{
		userRepo:    userRepo,
		helperTool:  helperRepo,
		walletRepo:  walletRepo,
		transaction: transaction,
	}
}
func (s *UserServiceImp) Register(ctx context.Context, registerData dto.RegisterRequest) (*dto.RegisterResponse, error) {
	var dataResponse dto.RegisterResponse
	err := s.transaction.WithinTransaction(ctx, func(ctx context.Context) error {
		user := payload.RegisterToUser(registerData)
		passwordHash, err := s.helperTool.HashPassword(user.Password, 12)
		if err != nil {
			return err
		}
		user.Password = string(passwordHash)
		userData, err := s.userRepo.Register(ctx, user)
		if err != nil {
			if util.CheckErrorUniqueEmail(err) {
				return util.ErrorEmailUnique
			}
			return err
		}
		userData.Name = user.Name
		wallet, err := s.walletRepo.CreateWallet(ctx, userData.Id)
		if err != nil {
			return err
		}
		dataResponse = payload.RegisterToResponse(*userData, *wallet)
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &dataResponse, nil
}
func (s *UserServiceImp) Login(ctx context.Context, loginData dto.Login) (*string, error) {
	user := payload.LoginToUser(loginData)
	userData, err := s.userRepo.Login(ctx, user)
	if err != nil {
		return nil, err
	}
	if userData != nil {
		passwordHash := []byte(userData.Password)
		result, err := s.helperTool.CheckPassword(user.Password, passwordHash)
		if err != nil {
			return nil, util.ErrorWrongPassword
		}
		if !result {
			return nil, util.ErrorWrongPassword
		}
		jwt, err := s.helperTool.CreateAndSign(*userData)
		if err != nil {
			return nil, err
		}
		return &jwt, nil
	}
	return nil, util.ErrorUserNotFound
}
func (s *UserServiceImp) GetSelf(ctx context.Context) (*dto.UserWalletResponse, error) {
	user, err := util.CheckClaim(ctx)
	if err != nil {
		return nil, err
	}
	data, err := s.userRepo.GetSelf(ctx, user.Id)
	if err != nil {
		return nil, err
	}
	dataResponse := payload.UserWalletToResponse(*data)
	return &dataResponse, nil

}
func (s *UserServiceImp) UpdateProfile(ctx context.Context, dataUser dto.UpdateProfile) error {
	user, err := util.CheckClaim(ctx)
	if err != nil {
		return err
	}
	err = s.userRepo.UpdateProfile(ctx, dataUser.Email, dataUser.FullName, user.Id)
	if err != nil {
		return err
	}
	return nil
}
func (s *UserServiceImp) UpdateAvatar(ctx *gin.Context) error {
	user, err := util.CheckClaim(ctx)
	if err != nil {
		return err
	}
	dataUser, err := s.userRepo.GetSelf(ctx, user.Id)
	if err != nil {
		return err
	}
	file, header, err := ctx.Request.FormFile("file")
	if err != nil {
		return err
	}
	defer file.Close()
	ext := filepath.Ext(header.Filename)
	if dataUser.Avatar != "avatar.png" {
		avatarPath := filepath.Join("images/", dataUser.Avatar)
		os.Remove(avatarPath)
	}
	newAvatarName := fmt.Sprintf("%s%s", uuid.New().String(), ext)
	newPath := filepath.Join("images/", newAvatarName)
	if err := ctx.SaveUploadedFile(header, newPath); err != nil {
		return err
	}
	err = s.userRepo.UpdateAvatar(ctx, newAvatarName, user.Id)
	if err != nil {
		return err
	}
	return nil
}
