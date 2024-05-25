package payload

import (
	"github.com/Lucifer07/e-wallet/dto"
	"github.com/Lucifer07/e-wallet/entity"
)

func LoginToUser(login dto.Login) entity.User {
	var user entity.User
	user.Email = login.Email
	user.Password = login.Password
	return user
}
