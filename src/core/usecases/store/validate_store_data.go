package usecases

import (
	"compreYa/src/core/entities"
	"compreYa/src/core/errors"
	"compreYa/src/core/providers"
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
)

type ValidateStoreData interface {
	Execute(c *gin.Context, store *entities.Store) *errors.ApiError
}

type ValidateStoreDataImpl struct {
	StoreRepository providers.Store
}

func (uc *ValidateStoreDataImpl) Execute(c *gin.Context, storeData *entities.Store) *errors.ApiError {
	storeDataMap, _ := uc.mapData(storeData)

	_ = uc.validateData(storeDataMap)

	return nil
}

func (uc *ValidateStoreDataImpl) mapData(storeData *entities.Store) (map[string]interface{}, *errors.ApiError) {
	var dataMap map[string]interface{}

	marshedData, err := json.Marshal(storeData)
	if err != nil {
		return nil, errors.NewInternalServerError(nil, "")
	}

	err = json.Unmarshal(marshedData, &dataMap)
	if err != nil {
		return nil, errors.NewInternalServerError(nil, "")
	}

	for k, v := range dataMap {
		fmt.Printf("key:%v value:%v\n", k, v)
	}

	return dataMap, nil
}

func (uc *ValidateStoreDataImpl) validateData(data map[string]interface{}) *errors.ApiError {
	for param, value := range data {
		if value != nil {
			switch param {
			case "name":
				fmt.Println("validating name", value)
			case "url_name":
				fmt.Println("validating url_name", value)
			}
		}
	}

	return nil
}
