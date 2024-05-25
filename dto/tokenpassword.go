package dto

type TokenPassword struct {
	Token    string `form:"token" binding:"required,min=16,max=16"`
	Password string `form:"password"`
}
