package dto

type LoginDTO struct {
	Email    string `json:"email" binding:"required" validate:"email"`
	Password string `json:"password" form:"password" binding:"required" validate:"min:6"`
}
