package repositories

import (
	"compreYa/src/core/entities"
	"compreYa/src/core/errors"
	"database/sql"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

type Auth struct {
	DB *sql.DB
}

// Add wraper for methods and manege errors

func (a *Auth) InsertUser(c *gin.Context, email, password string) *errors.ApiError {
	const query = "INSERT INTO compreYa.users (email, password) VALUES (?, ?)"

	var args []interface{}
	args = append(args, email)
	args = append(args, password)

	_, err := a.DB.Query(query, args...)
	if err != nil {
		messageError := errors.DataBaseError
		if strings.Contains(err.Error(), "Duplicate") {
			messageError = errors.EmailAlreadyRegisterd
		}
		return errors.NewInternalServerError(nil, messageError)
	}

	return nil
}

func (a *Auth) GetUserByEmail(c *gin.Context, email string) (*entities.User, *errors.ApiError) {
	const query = "SELECT id, email, password FROM compreYa.users WHERE email = ?"

	user := &entities.User{}
	var args []interface{}
	args = append(args, email)

	row, err := a.DB.Query(query, args...)
	if err != nil {
		return nil, errors.NewInternalServerError(nil, errors.DataBaseError)
	}

	row.Next()
	err = row.Scan(&user.ID, &user.Email, &user.Password)
	if err != nil {
		return nil, errors.NewInternalServerError(nil, errors.DataBaseError)
	}

	return user, nil
}

func (a *Auth) GetUserById(c *gin.Context, id int64) (*entities.User, *errors.ApiError) {
	const query = "SELECT id, email, password FROM compreYa.users WHERE id = ?"

	user := &entities.User{}
	var args []interface{}
	args = append(args, id)

	row, err := a.DB.Query(query, args...)
	if err != nil {
		return nil, errors.NewInternalServerError(nil, errors.DataBaseError)
	}

	row.Next()
	err = row.Scan(&user.ID, &user.Email, &user.Password)
	if err != nil {
		return nil, errors.NewInternalServerError(nil, errors.DataBaseError)
	}

	return user, nil
}

func (a *Auth) CheckEmailExistence(c *gin.Context, email string) (bool, *errors.ApiError) {
	const query = "SELECT COUNT(*)	FROM users WHERE email = ?"

	var exist int
	var args []interface{}
	args = append(args, email)

	rows, err := a.DB.Query(query, args...)
	if err != nil {
		return false, errors.NewInternalServerError(nil, errors.DataBaseError)
	}

	rows.Next()
	err = rows.Scan(&exist)
	if err != nil {
		return false, errors.NewInternalServerError(nil, errors.DataBaseError)
	}

	return exist == 1, nil
}

func (a *Auth) InsertPasswordRecoveryToken(c *gin.Context, user_id, expireDateToken int64, token string) *errors.ApiError {
	const query = "INSERT INTO compreYa.password_reset_tokens (user_id, token, token_expiry) VALUES (?, ?, ?)"

	var args []interface{}
	args = append(args, user_id)
	args = append(args, token)
	args = append(args, expireDateToken)

	_, err := a.DB.Query(query, args...)
	if err != nil {
		messageError := errors.DataBaseError
		return errors.NewInternalServerError(nil, messageError)
	}

	return nil
}

func (a *Auth) GetPasswordTokenRecovery(c *gin.Context, userID int64) (*string, *errors.ApiError) {
	const query = "SELECT token FROM password_reset_tokens WHERE user_id = ? AND token_expiry > ? ORDER BY token_expiry LIMIT 1"

	var token string
	var args []interface{}
	args = append(args, userID)
	args = append(args, time.Now().Unix())

	rows, err := a.DB.Query(query, args...)
	if err != nil {
		return nil, errors.NewInternalServerError(nil, errors.DataBaseError)
	}

	rows.Next()
	err = rows.Scan(&token)
	if err != nil {
		return nil, errors.NewInternalServerError(nil, errors.DataBaseError)
	}

	return &token, nil
}

func (a *Auth) CleanRecoveryTokens(c *gin.Context, userId int64) *errors.ApiError {
	const query = "DELETE FROM password_reset_tokens  WHERE user_id = ?"

	var args []interface{}
	args = append(args, userId)

	_, err := a.DB.Query(query, args...)
	if err != nil {
		messageError := errors.DataBaseError
		return errors.NewInternalServerError(nil, messageError)
	}

	return nil
}
func (a *Auth) UpdatePassword(c *gin.Context, userId int64, password string) *errors.ApiError {
	const query = "UPDATE users SET password = ? WHERE id = ?"

	var args []interface{}
	args = append(args, password)
	args = append(args, userId)

	_, err := a.DB.Query(query, args...)
	if err != nil {
		messageError := errors.DataBaseError
		return errors.NewInternalServerError(nil, messageError)
	}

	return nil
}
