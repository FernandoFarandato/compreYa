package store

import (
	contracts "compreYa/src/core/contracts/store"
	"compreYa/src/core/errors"
	usecases "compreYa/src/core/usecases/store"
	"compreYa/src/infrastructure/entrypoints"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ChangeStatusStore struct {
	ChangeStatus usecases.ChangeStatusStore
}

func (handler *ChangeStatusStore) Handle(c *gin.Context) {
	entrypoints.ErrorWrapper(handler.handle, c)
}

func (handler *ChangeStatusStore) handle(c *gin.Context) *errors.ApiError {
	requestData, ownerID, err := handler.getRequest(c)
	if err != nil {
		return errors.NewBadRequest(nil, errors.BindingError)
	}

	err = handler.ChangeStatus.Execute(c, requestData, ownerID)
	if err != nil {
		return err
	}

	c.JSON(http.StatusOK, nil)

	return nil
}

func (handler *ChangeStatusStore) getRequest(c *gin.Context) (*contracts.HideStoreData, int64, *errors.ApiError) {
	var request *contracts.HideStoreData
	err := c.Bind(&request)
	if err != nil {
		return nil, 0, errors.NewBadRequest(nil, errors.BindingError)
	}

	ownerID, parseErr := strconv.ParseInt(c.GetHeader("X-CALLER-ID"), 10, 64)
	if parseErr != nil {
		return nil, 0, errors.NewBadRequest(nil, errors.BindingError)
	}

	return request, ownerID, nil
}
