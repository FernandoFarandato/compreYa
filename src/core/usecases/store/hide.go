package usecases

import (
	contracts "compreYa/src/core/contracts/store"
	"compreYa/src/core/errors"
	"compreYa/src/core/providers"

	"github.com/gin-gonic/gin"
)

type ChangeStatusStore interface {
	Execute(c *gin.Context, request *contracts.HideStoreData, ownerID int64) *errors.ApiError
}

type ChangeStatusStoreImpl struct {
	StoreRepository providers.Store
}

func (uc *ChangeStatusStoreImpl) Execute(c *gin.Context, request *contracts.HideStoreData, ownerID int64) *errors.ApiError {
	store, err := uc.StoreRepository.GetStoreData(c, request.URLName)
	if err != nil {
		return err
	}

	if store.OwnerID != ownerID {
		return errors.NewNotAuthorizeError(nil, "User is not the owner of this store")
	}

	var isHidden bool
	isHidden = false
	if request.Status == "show" {
		isHidden = true
	}

	err = uc.StoreRepository.HideStore(c, isHidden, store.ID)
	if err != nil {
		return err
	}

	return nil
}
