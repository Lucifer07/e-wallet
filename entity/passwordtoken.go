package entity

import (
	"time"
)

type PasswordToken struct {
	Id        int
	UserId    int
	Token     string
	ExpiredAt time.Time
}
