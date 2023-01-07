package providers

import (
	"compreYa/src/core/errors"
	"github.com/gin-gonic/gin"
)

type Store interface {
	CreateStore(c *gin.Context, name, urlName string, ownerID int64) *errors.ApiError
	GetStoreData(c *gin.Context, name, urlName string, ownerID int64) *errors.ApiError
	GetUserStores(c *gin.Context, ownerID int64) (int64, *errors.ApiError)
	GetStoresCountByName(c *gin.Context, name string) (int64, *errors.ApiError)
	ValidateStoreURLName(c *gin.Context, name, urlName string, ownerID int64) *errors.ApiError
	HideStore(c *gin.Context, name, urlName string, ownerID int64) *errors.ApiError
	ShowStore(c *gin.Context, name, urlName string, ownerID int64) *errors.ApiError
}
