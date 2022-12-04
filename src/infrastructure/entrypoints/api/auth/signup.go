package auth

import (
	"compreYa/src/core/constants"
	contracts "compreYa/src/core/contracts/auth"
	"compreYa/src/core/errors"
	usecases "compreYa/src/core/usecases/auth"
	"compreYa/src/infrastructure/entrypoints"
	"compreYa/src/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type SignUp struct {
	SignUp usecases.SignUp
}

func (handler *SignUp) Handle(c *gin.Context) {
	entrypoints.ErrorWrapper(handler.handle, c)
}

func (handler *SignUp) handle(c *gin.Context) *errors.ApiError {
	request, err := handler.getRequest(c)
	if err != nil {
		return err
	}

	err = handler.validateFields(c, request)
	if err != nil {
		return err
	}

	err = handler.SignUp.Execute(c, request.Email, request.Password)
	if err != nil {
		return err
	}

	c.JSON(http.StatusOK, nil)
	return nil
}

func (handler *SignUp) getRequest(c *gin.Context) (*contracts.RegistrationData, *errors.ApiError) {
	var request *contracts.RegistrationData
	err := c.Bind(&request)
	if err != nil {
		return nil, errors.NewBadRequest(nil, errors.BindingError)
	}

	return request, nil
}

func (handler *SignUp) validateFields(c *gin.Context, request *contracts.RegistrationData) *errors.ApiError {
	emailValid := utils.ValidateRegex(request.Email, constants.Email)
	//passwordValid := utils.ValidateRegex(request.Password, constants.Password)

	if !(emailValid) {
		return errors.NewBadRequest(nil, errors.CredentialsNotValid)
	}

	return nil
}
