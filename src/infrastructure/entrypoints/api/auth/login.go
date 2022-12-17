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

type LogIn struct {
	usecases.LogIn
}

func (handler *LogIn) Handle(c *gin.Context) {
	entrypoints.ErrorWrapper(handler.handle, c)
}

func (handler *LogIn) handle(c *gin.Context) *errors.ApiError {
	request, err := handler.getRequest(c)
	if err != nil {
		return err
	}

	err = handler.validateFields(c, request)
	if err != nil {
		return err
	}

	authToken, err := handler.LogIn.Execute(c, request.Email, request.Password)
	if err != nil {
		return err
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorize", *authToken, 3600*24*15, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"access_token": fmt.Sprintf("%s", *authToken),
	})

	return nil
}

func (handler *LogIn) getRequest(c *gin.Context) (*contracts.RegistrationData, *errors.ApiError) {
	var request *contracts.RegistrationData
	err := c.Bind(&request)
	if err != nil {
		return nil, errors.NewBadRequest(nil, errors.BindingError)
	}

	return request, nil
}

func (handler *LogIn) validateFields(c *gin.Context, request *contracts.RegistrationData) *errors.ApiError {
	emailValid := utils.ValidateRegex(request.Email, constants.Email)
	//passwordValid := utils.ValidateRegex(request.Password, constants.Password)

	if !(emailValid) {
		return errors.NewBadRequest(nil, errors.CredentialsNotValid)
	}

	return nil
}
