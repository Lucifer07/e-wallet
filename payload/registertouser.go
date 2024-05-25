package payload

import (
	"github.com/Lucifer07/e-wallet/dto"
	"github.com/Lucifer07/e-wallet/entity"
)

func RegisterToUser(register dto.RegisterRequest) entity.User {
	var user entity.User
	user.Email = register.Email
	user.Name = register.Name
	user.Password = register.Password
	return user
}
