package payload

import (
	"github.com/Lucifer07/e-wallet/dto"
	"github.com/Lucifer07/e-wallet/entity"
)

func HistoryToPaylaterHistory(history entity.HistoryTransaction, walletNumber string, email string) dto.TopupPaylaterResponse {
	return dto.TopupPaylaterResponse{
		Id:                 history.Id,
		TransactionMethode: history.TransactionMethod,
		RecipientWalletId:  history.RecipientWalletId,
		Recipient:          dto.WalletNumberTransaction{WalletNumber: walletNumber},
		SenderWalletId:     history.SenderWalletId,
		Sender:             dto.Email{Email: email},
		Amount:             history.Amount,
		Description:        history.Description,
		CreatedAt:          history.CreatedAt,
		UpdatedAt:          history.UpdatedAt,
		DeletedAt:          history.DeletedAt.Time,
	}
}
