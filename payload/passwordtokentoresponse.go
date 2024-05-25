package payload

import (
	"github.com/Lucifer07/e-wallet/dto"
	"github.com/Lucifer07/e-wallet/entity"
)

func PasswordTokenToResponse(passwordToken entity.PasswordToken) dto.GetTokenResponse {
	return dto.GetTokenResponse{
		Token:   passwordToken.Token,
		ExpDate: passwordToken.ExpiredAt,
	}
}
