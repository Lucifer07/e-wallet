package dto

import (
	"time"

	"github.com/shopspring/decimal"
)

type WalletData struct {
	WalletId        int             `json:"id"`
	WalletNumber    int             `json:"wallet_number"`
	Balance         decimal.Decimal `json:"balance"`
	Income          decimal.Decimal `json:"income"`
	Expense         decimal.Decimal `json:"expense"`
	WalletCreatedAt time.Time       `json:"created_at"`
	WalletUpdatedAt time.Time       `json:"updated_at"`
}
