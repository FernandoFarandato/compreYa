package usecases

import (
	"compreYa/src/core/errors"
	"compreYa/src/core/providers"
	"crypto/rand"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type PasswordRecovery interface {
	ValidateRequest(c *gin.Context, email string) *errors.ApiError
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

	hashRecoveryToken, err := uc.encryptToken(recoveryToken)
	expireDateToken := time.Now().Add(time.Minute * 30).Unix()

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

func (uc *PasswordRecoveryImpl) generateRandomToken(length int) string {
	b := make([]byte, length)
	if _, err := rand.Read(b); err != nil {
		return ""
	}
	return hex.EncodeToString(b)
}

func (uc *PasswordRecoveryImpl) encryptToken(token string) ([]byte, *errors.ApiError) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(token), 10)
	if err != nil {
		return nil, errors.NewInternalServerError(nil, "Error hashing token")
	}

	return hashPassword, nil
}
