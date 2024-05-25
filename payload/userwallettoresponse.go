package payload

import "github.com/Lucifer07/e-wallet/dto"

func UserWalletToResponse(userWallet dto.UserWalet) dto.UserWalletResponse {
	walletData := dto.WalletData{
		WalletId:        userWallet.WalletId,
		WalletNumber:    userWallet.WalletNumber,
		Balance:         userWallet.Balance,
		Income:          userWallet.Income,
		Expense:         userWallet.Expense,
		WalletCreatedAt: userWallet.WalletCreatedAt,
		WalletUpdatedAt: userWallet.WalletUpdatedAt,
	}
	return dto.UserWalletResponse{
		Id:        userWallet.UserId,
		Name:      userWallet.Name,
		Email:     userWallet.Email,
		Avatar:    userWallet.Avatar,
		WalletId:  userWallet.WalletId,
		Wallet:    walletData,
		CreatedAt: userWallet.UserCreatedAt,
		UpdatedAt: userWallet.UserUpdatedAt,
	}
}
