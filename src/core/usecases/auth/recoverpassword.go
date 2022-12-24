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

	print("start")
	print(recoveryToken)
	print("end")

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

func (uc *PasswordRecoveryImpl) ChangePassword(c *gin.Context, userID int64, newPassword, recoveryToken string) *errors.ApiError {
	hashRecoveryToken, err := uc.encryptToken(recoveryToken)
	if err != nil {
		return err
	}

	isValid, err := uc.AuthRepository.CheckPasswordTokenRecovery(c, userID, string(hashRecoveryToken))
	if err != nil {
		return err
	}
	if !isValid {
		// return not auth
	}

	hashPassword, err := uc.encryptPassword(newPassword)
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

func (uc *PasswordRecoveryImpl) encryptToken(token string) ([]byte, *errors.ApiError) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(token), 10)
	if err != nil {
		return nil, errors.NewInternalServerError(nil, "Error hashing token")
	}

	return hashPassword, nil
}

func (uc *PasswordRecoveryImpl) encryptPassword(password string) ([]byte, *errors.ApiError) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return nil, errors.NewInternalServerError(nil, "Error hashing password")
	}

	return hashPassword, nil
}
