package dto

type RegisterRequest struct {
	Name     string `form:"name" binding:"required,alpha"`
	Email    string `form:"email" binding:"required,email"`
	Password string `form:"password" binding:"required"`
}
