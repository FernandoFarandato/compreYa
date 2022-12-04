package usecases

import (
	"fmt"
	"os"
	"time"

	"compreYa/src/core/entities"
	"compreYa/src/core/errors"
	"compreYa/src/core/providers"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type LogIn interface {
	Execute(c *gin.Context, email, password string) (*string, *errors.ApiError)
}

type LogInImpl struct {
	AuthRepository providers.Auth
}

func (uc *LogInImpl) Execute(c *gin.Context, email, password string) (*string, *errors.ApiError) {
	exists, err := uc.AuthRepository.CheckEmailExistence(c, email)
	if !exists {
		return nil, errors.NewBadRequest(nil, errors.NoAccountRegistered)
	}

	user, err := uc.AuthRepository.GetUserByEmail(c, email)

	err = uc.verifyPassword(user.Password, password)
	if err != nil {
		return nil, err
	}

	token, err := uc.createToken(user)
	if err != nil {
		return nil, err
	}

	return token, nil
}

func (uc *LogInImpl) verifyPassword(userPassword, password string) *errors.ApiError {
	err := bcrypt.CompareHashAndPassword([]byte(userPassword), []byte(password))
	if err != nil {
		return errors.NewInternalServerError(nil, errors.CredentialsNotValid)
	}

	return nil
}

func (uc *LogInImpl) createToken(user *entities.User) (*string, *errors.ApiError) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 15).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("HASHING_KEY")))
	if err != nil {
		return nil, errors.NewInternalServerError(nil, "Error Creating Token")
	}

	fmt.Println(tokenString, err)

	return &tokenString, nil
}
