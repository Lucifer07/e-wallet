package dto

import "time"

type GetTokenResponse struct {
	Token   string    `json:"token"`
	ExpDate time.Time `json:"expired_at"`
}
