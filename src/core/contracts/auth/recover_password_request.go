package auth

type RecoverPasswordRequest struct {
	Email string `json:"email" binding:"required"`
}
