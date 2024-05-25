package dto

import (
	"time"

	"github.com/shopspring/decimal"
)

type TopupPaylaterResponse struct {
	Id                 int                     `json:"id"`
	TransactionMethode string                  `json:"transaction_methode"`
	RecipientWalletId  int                     `json:"recipient_wallet_id"`
	Recipient          WalletNumberTransaction `json:"recipient"`
	SenderWalletId     int                     `json:"sender_wallet_id"`
	Sender             Email                   `json:"sender"`
	Amount             decimal.Decimal         `json:"amount"`
	Description        string                  `json:"description"`
	CreatedAt          time.Time               `json:"created_at"`
	UpdatedAt          time.Time               `json:"updated_at"`
	DeletedAt          time.Time               `json:"deleted_at,omitempty"`
}
