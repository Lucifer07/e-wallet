package dto
type UpdateProfile struct {
	Email string `json:"email" binding:"required,email"`
	FullName string `json:"fullname" binding:"required"`
}