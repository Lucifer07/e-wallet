package dto

import (
	"time"

	"github.com/shopspring/decimal"
)

type WalletResponse struct {
	WalletNumber string          `json:"wallet_number"`
	Balance      decimal.Decimal `json:"balance"`
	CreatedAt    time.Time       `json:"created_at"`
}
