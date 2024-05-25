package dto

type GettokenRequest struct {
	Email string `form:"email" binding:"required,email"`
}
