package dto

import "time"

type UserWalletResponse struct {
	Id        int        `json:"id"`
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	Avatar    string     `json:"avatar"`
	WalletId  int        `json:"wallet_id"`
	Wallet    WalletData `json:"wallet"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}
