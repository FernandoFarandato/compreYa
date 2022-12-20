package providers

import (
	"compreYa/src/core/entities"
	"compreYa/src/core/errors"
	"github.com/gin-gonic/gin"
)

type Auth interface {
	InsertUser(c *gin.Context, email, password string) *errors.ApiError
	CheckEmailExistence(c *gin.Context, email string) (bool, *errors.ApiError)
	GetUserByEmail(c *gin.Context, email string) (*entities.User, *errors.ApiError)
	GetUserById(c *gin.Context, id int64) (*entities.User, *errors.ApiError)
	InsertPasswordRecoveryToken(c *gin.Context, userId, expireDateToken int64, token string) *errors.ApiError
}
