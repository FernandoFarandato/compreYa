package entrypoints

import (
	"compreYa/src/core/errors"
	"github.com/gin-gonic/gin"
)

type Handler interface {
	Handle(c *gin.Context)
}

type HandlerFunc func(c *gin.Context) *errors.ApiError

func ErrorWrapper(handlerFunc HandlerFunc, c *gin.Context) {
	err := handlerFunc(c)
	if err != nil {
		c.JSON(err.Status, err)
		c.Abort()
	}
}
