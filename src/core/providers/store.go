package providers

import (
	"compreYa/src/core/entities"
	"compreYa/src/core/errors"
	"github.com/gin-gonic/gin"
)

type Store interface {
	CreateStore(c *gin.Context, name, urlName string, ownerID int64) *errors.ApiError
	DeleteStore(c *gin.Context, storeID, ownerID int64) *errors.ApiError
	GetStoreData(c *gin.Context, urlName string) (*entities.Store, *errors.ApiError)
	GetUserStores(c *gin.Context, ownerID int64) (int64, *errors.ApiError)
	ValidateStoreName(c *gin.Context, name string) *errors.ApiError
	ValidateStoreURLName(c *gin.Context, name, urlName string, ownerID int64) *errors.ApiError
	HideStore(c *gin.Context, name, urlName string, ownerID int64) *errors.ApiError
	ShowStore(c *gin.Context, name, urlName string, ownerID int64) *errors.ApiError
}
