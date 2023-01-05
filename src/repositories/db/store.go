package repositories

import (
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

func (s *Store) GetStoreData(c *gin.Context, name, urlName string, ownerID int64) *errors.ApiError {
	const query = "INSERT INTO store (name, urlName, ownerID) VALUES (?, ?, ?)"

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
	const query = "INSERT INTO store (name, urlName, ownerID) VALUES (?, ?, ?)"

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
	const query = "INSERT INTO store (name, urlName, ownerID) VALUES (?, ?, ?)"

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

func (s *Store) HideStore(c *gin.Context, name, urlName string, ownerID int64) *errors.ApiError {
	const query = "INSERT INTO store (name, urlName, ownerID) VALUES (?, ?, ?)"

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

func (s *Store) ShowStore(c *gin.Context, name, urlName string, ownerID int64) *errors.ApiError {
	const query = "INSERT INTO store (name, urlName, ownerID) VALUES (?, ?, ?)"

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
