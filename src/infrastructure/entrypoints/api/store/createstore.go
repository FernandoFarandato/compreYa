package auth

import (
	contracts "compreYa/src/core/contracts/store"
	"compreYa/src/core/entities"
	"compreYa/src/core/errors"
	usecases "compreYa/src/core/usecases/store"
	"compreYa/src/infrastructure/entrypoints"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

const tokenExpirationTime = 3600 * 24 * 15

type CreateStore struct {
	CreateStore usecases.CreateStore
}

func (handler *CreateStore) Handle(c *gin.Context) {
	entrypoints.ErrorWrapper(handler.handle, c)
}

func (handler *CreateStore) handle(c *gin.Context) *errors.ApiError {
	request, err := handler.getRequest(c)
	if err != nil {
		return err
	}
	/*
		err = handler.validateFields(c, request)
		if err != nil {
			return err
		}
	*/
	requestUserID, parseErr := strconv.ParseInt(c.GetHeader("X-CALLER-ID"), 10, 64)
	if parseErr != nil {
		return errors.NewBadRequest(nil, errors.BindingError)
	}

	store := entities.NewStore(request.Name, request.URLName, &requestUserID)
	err = handler.CreateStore.Execute(c, store)
	if err != nil {
		return err
	}

	c.JSON(http.StatusOK, nil)

	return nil
}

func (handler *CreateStore) getRequest(c *gin.Context) (*contracts.CreateStoreData, *errors.ApiError) {
	var request *contracts.CreateStoreData
	err := c.Bind(&request)
	if err != nil {
		return nil, errors.NewBadRequest(nil, errors.BindingError)
	}

	return request, nil
}

func (handler *CreateStore) validateFields(c *gin.Context, request *contracts.CreateStoreData) *errors.ApiError {

	// validate that are strings

	return nil
}
