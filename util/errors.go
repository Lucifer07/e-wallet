package util

import (
	"errors"
	"strings"
)

var (
	ErrorBadRequest      = errors.New("error bad request")
	ErrorInternal        = errors.New("internal server error")
	ErrorUserNotFound    = errors.New("user not found")
	ErrorWrongPassword   = errors.New("wrong password")
	ErrorUnauthorized    = errors.New("unauthorized")
	ErrorInvalidToken    = errors.New("invalid token")
	ErrorForbidden       = errors.New("forbidden access")
	ErrorEmailUnique     = errors.New("email already used")
	ErrorTokenExp        = errors.New("token expired")
	ErrorBalance         = errors.New("insufficient balance")
	ErroMinimumTranfer   = errors.New("minimum transfer is 1.000")
	ErroMinimumTopUp     = errors.New("minimum topup is 50.000")
	ErroMaximalTranfer   = errors.New("maxsimum transfer is 50.000.000")
	ErroMaximalTopUp     = errors.New("maxsimum topup is 10.000.000")
	ErrorInvalidTransfer = errors.New("unable to transfer to own account")
)

func CheckErrorUniqueEmail(err error) bool {
	unique := "unique"
	email := "email"
	uniqueCondition := strings.Contains(err.Error(), unique)
	emailCondition := strings.Contains(err.Error(), email)
	if uniqueCondition && emailCondition {
		return true
	}
	return false
}
