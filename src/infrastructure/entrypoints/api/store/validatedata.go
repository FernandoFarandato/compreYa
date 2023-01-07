package auth

import (
	contracts "compreYa/src/core/contracts/store"
	"compreYa/src/core/entities"
	"compreYa/src/core/errors"
	usecases "compreYa/src/core/usecases/store"
	"compreYa/src/infrastructure/entrypoints"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ValidateData struct {
	ValidateStoreData usecases.ValidateStoreData
}

func (handler *ValidateData) Handle(c *gin.Context) {
	entrypoints.ErrorWrapper(handler.handle, c)
}

func (handler *ValidateData) handle(c *gin.Context) *errors.ApiError {
	request, err := handler.getRequest(c)
	if err != nil {
		return err
	}

	store := entities.NewStore(request.Name, request.URLName, nil)
	err = handler.ValidateStoreData.Execute(c, store)
	if err != nil {
		return err
	}

	c.JSON(http.StatusOK, nil)

	return nil
}

func (handler *ValidateData) getRequest(c *gin.Context) (*contracts.ValidateStoreData, *errors.ApiError) {
	var request *contracts.ValidateStoreData
	err := c.Bind(&request)
	if err != nil {
		return nil, errors.NewBadRequest(nil, errors.BindingError)
	}

	return request, nil
}
