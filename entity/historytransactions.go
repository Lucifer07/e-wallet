package entity

import (
	"database/sql"
	"time"

	"github.com/shopspring/decimal"
)

type HistoryTransaction struct {
	Id                int
	TransactionMethod string
	Amount            decimal.Decimal
	SenderWalletId    int
	RecipientWalletId int
	Description       string
	CreatedAt         time.Time
	UpdatedAt         time.Time
	DeletedAt         sql.NullTime `json:"-"`
}
