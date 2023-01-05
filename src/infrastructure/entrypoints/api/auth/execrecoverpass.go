package auth

import (
	"compreYa/src/core/constants"
	contracts "compreYa/src/core/contracts/auth"
	"compreYa/src/core/errors"
	usecases "compreYa/src/core/usecases/auth"
	"compreYa/src/infrastructure/entrypoints"
	"compreYa/src/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type RecoverPasswordRequest struct {
	RecoverPassword usecases.PasswordRecovery
}

func (handler *RecoverPasswordRequest) Handle(c *gin.Context) {
	entrypoints.ErrorWrapper(handler.handle, c)
}

func (handler *RecoverPasswordRequest) handle(c *gin.Context) *errors.ApiError {
	request, err := handler.getRequest(c)
	if err != nil {
		return err
	}

	err = handler.validateFields(c, request)
	if err != nil {
		return err
	}

	err = handler.RecoverPassword.ValidateRequest(c, request.Email)
	if err != nil {
		return err
	}

	// show error only for system errors not if email not_found
	c.JSON(http.StatusOK, gin.H{
		"status": fmt.Sprintf("%s", "ok"),
	})

	return nil
}

func (handler *RecoverPasswordRequest) getRequest(c *gin.Context) (*contracts.RecoverPasswordRequest, *errors.ApiError) {
	var request *contracts.RecoverPasswordRequest
	err := c.Bind(&request)
	if err != nil {
		return nil, errors.NewBadRequest(nil, errors.BindingError)
	}

	return request, nil
}

func (handler *RecoverPasswordRequest) validateFields(c *gin.Context, request *contracts.RecoverPasswordRequest) *errors.ApiError {
	emailValid := utils.ValidateRegex(request.Email, constants.Email)

	if !(emailValid) {
		return errors.NewBadRequest(nil, errors.CredentialsNotValid)
	}

	return nil
}
