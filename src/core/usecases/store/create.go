package usecases

import (
	"compreYa/src/core/entities"
	"compreYa/src/core/errors"
	"compreYa/src/core/providers"

	"github.com/gin-gonic/gin"
)

type CreateStore interface {
	Execute(c *gin.Context, store *entities.Store) *errors.ApiError
}

type CreateStoreImpl struct {
	StoreRepository providers.Store
}

func (uc *CreateStoreImpl) Execute(c *gin.Context, store *entities.Store) *errors.ApiError {
	// validate if user already has a store
	stores, err := uc.StoreRepository.GetUserStores(c, *store.OwnerID)
	if err != nil {
		return err
	}
	if stores != 0 {
		return errors.NewNotAuthorizeError(nil, "") // max amount of stores reached - another error code
	}

	err = uc.StoreRepository.CreateStore(c, store.Name, store.URLName, *store.OwnerID)
	if err != nil {
		return err
	}

	return nil
}
