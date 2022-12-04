package usecases

import (
	"compreYa/src/core/errors"
	"compreYa/src/core/providers"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type SignUp interface {
	Execute(c *gin.Context, email, password string) *errors.ApiError
}

type SignUpImpl struct {
	AuthRepository providers.Auth
}

func (uc *SignUpImpl) Execute(c *gin.Context, email, password string) *errors.ApiError {
	exists, err := uc.AuthRepository.CheckEmailExistence(c, email)
	if exists || err != nil {
		if exists {
			return errors.NewBadRequest(nil, errors.EmailAlreadyRegisterd)
		}
		return err
	}

	hashPassword, err := uc.encryptPassword(password)
	if err != nil {
		return err
	}

	err = uc.AuthRepository.InsertUser(c, email, string(hashPassword))
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func (uc *SignUpImpl) encryptPassword(password string) ([]byte, *errors.ApiError) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return nil, errors.NewInternalServerError(nil, "Error hashing password")
	}

	return hashPassword, nil
}
