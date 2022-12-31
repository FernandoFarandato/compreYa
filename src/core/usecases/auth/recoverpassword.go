package usecases

import (
	"compreYa/src/core/errors"
	"compreYa/src/core/providers"
	"compreYa/src/utils"
	"crypto/rand"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"time"
)

type PasswordRecovery interface {
	ValidateRequest(c *gin.Context, email string) *errors.ApiError
	ChangePassword(c *gin.Context, userID int64, newPassword, recoveryToken string) *errors.ApiError
}

type PasswordRecoveryImpl struct {
	AuthRepository providers.Auth
	Email          providers.SendGrid
}

func (uc *PasswordRecoveryImpl) ValidateRequest(c *gin.Context, email string) *errors.ApiError {
	user, err := uc.AuthRepository.GetUserByEmail(c, email)
	if user == nil || err != nil {
		return err
	}

	recoveryToken := uc.generateRandomToken(255)
	if err != nil {
		return err
	}

	hashRecoveryToken, err := utils.EncryptString(recoveryToken)
	expireDateToken := time.Now().Add(time.Minute * 30).Unix()

	err = uc.AuthRepository.CleanRecoveryTokens(c, user.ID)
	if err != nil {
		return err
	}

	err = uc.AuthRepository.InsertPasswordRecoveryToken(c, user.ID, expireDateToken, string(hashRecoveryToken))
	if err != nil {
		return err
	}

	err = uc.Email.SendPasswordRecovery(user, recoveryToken)
	if err != nil {
		return err
	}

	return nil
}

func (uc *PasswordRecoveryImpl) ChangePassword(c *gin.Context, userID int64, newPassword, token string) *errors.ApiError {
	hashedToken, err := uc.AuthRepository.GetPasswordTokenRecovery(c, userID)
	if err != nil {
		return err
	}

	err = utils.VerifyEncryptedString(*hashedToken, token)
	if err != nil {
		// return not auth
		return err
	}

	hashPassword, err := utils.EncryptString(newPassword)
	if err != nil {
		return err
	}

	err = uc.AuthRepository.UpdatePassword(c, userID, string(hashPassword))
	if err != nil {
		return err
	}

	// Send password changed notification
	//err = uc.Email.SendPasswordRecovery(user, recoveryToken)

	return nil
}

func (uc *PasswordRecoveryImpl) generateRandomToken(length int) string {
	b := make([]byte, length)
	if _, err := rand.Read(b); err != nil {
		return ""
	}
	return hex.EncodeToString(b)
}
