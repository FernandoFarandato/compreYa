package usecases

import (
	"compreYa/src/core/errors"
	"compreYa/src/core/providers"
	"compreYa/src/utils"
	"fmt"
	"github.com/gin-gonic/gin"
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

	hashPassword, err := utils.EncryptString(password)
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
