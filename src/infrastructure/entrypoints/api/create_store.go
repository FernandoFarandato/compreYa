package api

import (
	"compreYa/src/core/errors"
	"compreYa/src/infrastructure/entrypoints"
	"fmt"
	"github.com/gin-gonic/gin"
)

type CreateStore struct {
}

func (handler *CreateStore) Handle(c *gin.Context) {
	entrypoints.ErrorWrapper(handler.handle, c)
}

func (handler *CreateStore) handle(c *gin.Context) *errors.ApiError {

	fmt.Println("Validation passed")
	c.JSON(200, gin.H{
		"response": "andom",
	})
	return nil
}
