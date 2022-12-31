package auth

type RecoverPasswordChange struct {
	UserID        int64  `json:"user_id" binding:"required"`
	NewPassword   string `json:"new_password" binding:"required"`
	RecoveryToken string `json:"recovery_token" binding:"required"`
}
