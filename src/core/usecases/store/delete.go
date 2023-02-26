package usecases

import (
	"compreYa/src/core/errors"
	"compreYa/src/core/providers"

	"github.com/gin-gonic/gin"
)

type DeleteStore interface {
	Execute(c *gin.Context, urlName string, ownerID int64) *errors.ApiError
}

type DeleteStoreImpl struct {
	StoreRepository providers.Store
}

func (uc *DeleteStoreImpl) Execute(c *gin.Context, urlName string, ownerID int64) *errors.ApiError {
	store, err := uc.StoreRepository.GetStoreData(c, urlName)
	if err != nil {
		return err
	}

	if store.OwnerID != ownerID {
		return errors.NewNotAuthorizeError(nil, "User is not the owner of this store")
	}

	err = uc.StoreRepository.DeleteStore(c, store.ID, ownerID)
	if err != nil {
		return err
	}

	return nil
}
