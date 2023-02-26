package repositories

import (
	"compreYa/src/core/entities"
	"compreYa/src/core/errors"
	"database/sql"
	"github.com/gin-gonic/gin"
	"strings"
)

type Store struct {
	DB *sql.DB
}

// Add wraper for methods and manege errors

func (s *Store) CreateStore(c *gin.Context, name, urlName string, ownerID int64) *errors.ApiError {
	const query = "INSERT INTO store (name, url_name, owner_id) VALUES (?, ?, ?)"

	var args []interface{}
	args = append(args, name)
	args = append(args, urlName)
	args = append(args, ownerID)

	_, err := s.DB.Query(query, args...)
	if err != nil {
		messageError := errors.DataBaseError
		if strings.Contains(err.Error(), "Duplicate") {
			messageError = errors.EmailAlreadyRegisterd
		}
		return errors.NewInternalServerError(nil, messageError)
	}

	return nil
}

func (s *Store) DeleteStore(c *gin.Context, storeID, ownerID int64) *errors.ApiError {
	const query = "DELETE FROM store WHERE id = ? AND owner_id = ?"

	var args []interface{}
	args = append(args, storeID)
	args = append(args, ownerID)

	_, err := s.DB.Query(query, args...)
	if err != nil {
		messageError := errors.DataBaseError
		if strings.Contains(err.Error(), "Duplicate") {
			messageError = errors.EmailAlreadyRegisterd
		}
		return errors.NewInternalServerError(nil, messageError)
	}

	return nil
}

func (s *Store) GetStoreData(c *gin.Context, urlName string) (*entities.Store, *errors.ApiError) {
	const query = "SELECT id, name, url_name, owner_id FROM store WHERE url_name = ?"

	var store = &entities.Store{}

	var args []interface{}
	args = append(args, urlName)

	rows, err := s.DB.Query(query, args...)
	if err != nil {
		return nil, errors.NewInternalServerError(nil, errors.DataBaseError)
	}

	rows.Next()
	err = rows.Scan(&store.ID, &store.Name, &store.URLName, &store.OwnerID)
	if err != nil {
		if strings.Contains(err.Error(), "Rows are closed") {
			return nil, errors.NewResourceNotFoundError(nil, errors.StoreNotFound)
		}
		return nil, errors.NewInternalServerError(nil, errors.DataBaseError)
	}

	return store, nil
}

func (s *Store) GetUserStores(c *gin.Context, ownerID int64) (int64, *errors.ApiError) {
	const query = "SELECT COUNT(*) FROM store WHERE owner_id = ?"

	var stores int64
	var args []interface{}
	args = append(args, ownerID)

	rows, err := s.DB.Query(query, args...)
	if err != nil {
		return 0, errors.NewInternalServerError(nil, errors.DataBaseError)
	}

	rows.Next()
	err = rows.Scan(&stores)
	if err != nil {
		return 0, errors.NewInternalServerError(nil, errors.DataBaseError)
	}

	return stores, nil
}

func (s *Store) ValidateStoreName(c *gin.Context, name string) *errors.ApiError {
	const query = "INSERT INTO store (name, url_name, owner_id) VALUES (?, ?, ?)"

	var args []interface{}
	args = append(args, name)

	_, err := s.DB.Query(query, args...)
	if err != nil {
		messageError := errors.DataBaseError
		if strings.Contains(err.Error(), "Duplicate") {
			messageError = errors.EmailAlreadyRegisterd
		}
		return errors.NewInternalServerError(nil, messageError)
	}

	return nil
}

func (s *Store) ValidateStoreURLName(c *gin.Context, name, urlName string, ownerID int64) *errors.ApiError {
	const query = "INSERT INTO store (name, url_name, owner_id) VALUES (?, ?, ?)"

	var args []interface{}
	args = append(args, name)
	args = append(args, urlName)
	args = append(args, ownerID)

	_, err := s.DB.Query(query, args...)
	if err != nil {
		messageError := errors.DataBaseError
		if strings.Contains(err.Error(), "Duplicate") {
			messageError = errors.EmailAlreadyRegisterd
		}
		return errors.NewInternalServerError(nil, messageError)
	}

	return nil
}

func (s *Store) HideStore(c *gin.Context, hidden bool, storeID int64) *errors.ApiError {
	const query = "UPDATE store SET hidden = ? WHERE id = ?"

	var args []interface{}
	args = append(args, hidden)
	args = append(args, storeID)

	_, err := s.DB.Query(query, args...)
	if err != nil {
		messageError := errors.DataBaseError
		if strings.Contains(err.Error(), "Duplicate") {
			messageError = errors.EmailAlreadyRegisterd
		}
		return errors.NewInternalServerError(nil, messageError)
	}

	return nil
}

func (s *Store) ShowStore(c *gin.Context, name, urlName string, ownerID int64) *errors.ApiError {
	const query = "INSERT INTO store (name, url_name, owner_id) VALUES (?, ?, ?)"

	var args []interface{}
	args = append(args, name)
	args = append(args, urlName)
	args = append(args, ownerID)

	_, err := s.DB.Query(query, args...)
	if err != nil {
		messageError := errors.DataBaseError
		if strings.Contains(err.Error(), "Duplicate") {
			messageError = errors.EmailAlreadyRegisterd
		}
		return errors.NewInternalServerError(nil, messageError)
	}

	return nil
}
