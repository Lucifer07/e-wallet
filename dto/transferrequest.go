package dto

import "github.com/shopspring/decimal"

type TransferRequest struct {
	Amount       decimal.Decimal `json:"amount" binding:"required,numeric,min=50000,max=10000000"`
	WalletNumber string          `json:"wallet_number" binding:"required"`
	Description  string          `json:"description" binding:"max=35"`
}
