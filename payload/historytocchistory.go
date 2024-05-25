package payload

import (
	"github.com/Lucifer07/e-wallet/dto"
	"github.com/Lucifer07/e-wallet/entity"
)

func HistoryToCCHistory(history entity.HistoryTransaction, walletNumber string, ccNumber int) dto.TopupCreditCardResponse {
	return dto.TopupCreditCardResponse{
		Id:                 history.Id,
		TransactionMethode: history.TransactionMethod,
		RecipientWalletId:  history.RecipientWalletId,
		Recipient:          dto.WalletNumberTransaction{WalletNumber: walletNumber},
		SenderWalletId:     history.SenderWalletId,
		Sender:             dto.CreditCardNumber{CreditCardNumber: ccNumber},
		Amount:             history.Amount,
		Description:        history.Description,
		CreatedAt:          history.CreatedAt,
		UpdatedAt:          history.UpdatedAt,
		DeletedAt:          history.DeletedAt.Time,
	}
}
