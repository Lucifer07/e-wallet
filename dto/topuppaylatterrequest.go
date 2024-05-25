package dto

import "github.com/shopspring/decimal"

type TopupPaylaterRequest struct {
	Amount      decimal.Decimal `json:"amount" binding:"required,min=50000,max=10000000"`
	Description string          `json:"description"`
}
