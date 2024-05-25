package dto

import "time"

type RegisterResponse struct {
	Id        int            `json:"id"`
	Name      string         `json:"name"`
	Email     string         `json:"email"`
	Wallet    WalletResponse `json:"wallet"`
	CreatedAt time.Time      `json:"created_at"`
}
