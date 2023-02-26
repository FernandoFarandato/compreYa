package store

import (
	"compreYa/src/core/errors"
	usecases "compreYa/src/core/usecases/store"
	"compreYa/src/infrastructure/entrypoints"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type DeleteStore struct {
	DeleteStore usecases.DeleteStore
}

func (handler *DeleteStore) Handle(c *gin.Context) {
	entrypoints.ErrorWrapper(handler.handle, c)
}

func (handler *DeleteStore) handle(c *gin.Context) *errors.ApiError {
	urlName, ownerID, err := handler.getParams(c)
	if err != nil {
		return errors.NewBadRequest(nil, errors.BindingError)
	}

	err = handler.DeleteStore.Execute(c, urlName, ownerID)
	if err != nil {
		return err
	}

	c.JSON(http.StatusOK, nil)

	return nil
}

func (handler *DeleteStore) getParams(c *gin.Context) (string, int64, *errors.ApiError) {
	urlName := c.Param("url_name")
	if urlName == "" {
		return "", 0, errors.NewBadRequest(nil, errors.BindingError)
	}

	ownerID, parseErr := strconv.ParseInt(c.GetHeader("X-CALLER-ID"), 10, 64)
	if parseErr != nil {
		return "", 0, errors.NewBadRequest(nil, errors.BindingError)
	}

	return urlName, ownerID, nil
}
