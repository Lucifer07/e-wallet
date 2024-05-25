package payload

import (
	"github.com/Lucifer07/e-wallet/dto"
	"github.com/Lucifer07/e-wallet/entity"
)

func WalletTransactionToResponse(wallets []entity.Wallet, history entity.HistoryTransaction) dto.WalletTransactionResponse {

	return dto.WalletTransactionResponse{
		Id:                 history.Id,
		TransactionMethode: history.TransactionMethod,
		RecipientWalletId:  history.RecipientWalletId,
		Recipient:          dto.WalletNumberTransaction{WalletNumber: wallets[0].WalletNumber},
		SenderWalletId:     history.SenderWalletId,
		Sender:             dto.WalletNumberTransaction{WalletNumber: wallets[1].WalletNumber},
		Amount:             history.Amount,
		Description:        history.Description,
		CreatedAt:          history.CreatedAt,
		UpdatedAt:          history.UpdatedAt,
		DeletedAt:          history.DeletedAt.Time,
	}
}
