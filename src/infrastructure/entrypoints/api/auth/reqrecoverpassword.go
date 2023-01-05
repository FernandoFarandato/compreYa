package auth

import (
	contracts "compreYa/src/core/contracts/auth"
	"compreYa/src/core/errors"
	usecases "compreYa/src/core/usecases/auth"
	"compreYa/src/infrastructure/entrypoints"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type RecoverPasswordChange struct {
	RecoverPassword usecases.PasswordRecovery
}

func (handler *RecoverPasswordChange) Handle(c *gin.Context) {
	entrypoints.ErrorWrapper(handler.handle, c)
}

func (handler *RecoverPasswordChange) handle(c *gin.Context) *errors.ApiError {
	request, err := handler.getRequest(c)
	if err != nil {
		return err
	}

	err = handler.RecoverPassword.ChangePassword(c, request.UserID, request.NewPassword, request.RecoveryToken)
	if err != nil {
		return err
	}

	// show error only for system errors not if email not_found
	c.JSON(http.StatusOK, gin.H{
		"access_token": fmt.Sprintf("%s", "ok"),
	})

	return nil
}

func (handler *RecoverPasswordChange) getRequest(c *gin.Context) (*contracts.RecoverPasswordChange, *errors.ApiError) {
	var request *contracts.RecoverPasswordChange
	err := c.Bind(&request)
	if err != nil {
		return nil, errors.NewBadRequest(nil, errors.BindingError)
	}

	return request, nil
}
