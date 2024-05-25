package dto

import (
	"time"

	"github.com/shopspring/decimal"
)

type UserWalet struct {
	UserId          int
	Name            string
	Avatar          string
	Email           string
	UserCreatedAt   time.Time
	UserUpdatedAt   time.Time
	WalletId        int
	WalletNumber    int
	Balance         decimal.Decimal
	WalletCreatedAt time.Time
	WalletUpdatedAt time.Time
	Income          decimal.Decimal
	Expense         decimal.Decimal
}
