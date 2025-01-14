package dto

type RegisterRequest struct {
	Username string `json:"username" binding:"required,min=3,max=20"`
	Password string `json:"password" binding:"required,min=3"`
	Role     string `json:"role" binding:"required"`
}
