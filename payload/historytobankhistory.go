package payload

import (
	"github.com/Lucifer07/e-wallet/dto"
	"github.com/Lucifer07/e-wallet/entity"
)

func HistoryToBankHistory(history entity.HistoryTransaction, walletNumber string, accountNumber int) dto.TopupBankResponse {
	return dto.TopupBankResponse{
		Id:                 history.Id,
		TransactionMethode: history.TransactionMethod,
		RecipientWalletId:  history.RecipientWalletId,
		Recipient:          dto.WalletNumberTransaction{WalletNumber: walletNumber},
		SenderWalletId:     history.SenderWalletId,
		Sender:             dto.BankAccountNumber{AccountNumber: accountNumber},
		Amount:             history.Amount,
		Description:        history.Description,
		CreatedAt:          history.CreatedAt,
		UpdatedAt:          history.UpdatedAt,
		DeletedAt:          history.DeletedAt.Time,
	}
}
