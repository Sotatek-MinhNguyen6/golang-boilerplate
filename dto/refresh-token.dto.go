package dto

type RefreshTokenRequest struct {
	RefreshToken string `json:"refreshToken" binding:"required"`
}
