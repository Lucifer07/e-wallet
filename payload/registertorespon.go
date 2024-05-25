package payload

import (
	"github.com/Lucifer07/e-wallet/dto"
	"github.com/Lucifer07/e-wallet/entity"
)

func RegisterToResponse(user entity.User, wallet entity.Wallet) dto.RegisterResponse {
	walletResponse := dto.WalletResponse{
		WalletNumber: wallet.WalletNumber,
		Balance:      wallet.Balance,
		CreatedAt:    wallet.CreatedAt,
	}
	response := dto.RegisterResponse{
		Id:        user.Id,
		Name:      user.Name,
		Email:     user.Email,
		Wallet:    walletResponse,
		CreatedAt: user.CreatedAt,
	}
	return response
}
