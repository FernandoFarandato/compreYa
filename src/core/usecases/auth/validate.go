package usecases

import (
	"compreYa/src/core/entities"
	"fmt"
	"os"
	"time"

	"compreYa/src/core/errors"
	"compreYa/src/core/providers"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type Validate interface {
	ValidateToken(c *gin.Context, token *string) (*entities.User, *errors.ApiError)
}

type ValidateImpl struct {
	AuthRepository providers.Auth
}

func (uc *ValidateImpl) ValidateToken(c *gin.Context, tokenString *string) (*int64, *errors.ApiError) {
	token, err := jwt.Parse(*tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("HASHING_KEY")), nil
	})
	if err != nil {
		return nil, errors.NewInternalServerError(nil, "Error parsing token")
	}

	var userID int64
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			return nil, errors.NewNotAuthorizeError(nil, "")
		}

		userID = claims["sub"].(int64)
	} else {
		return nil, errors.NewNotAuthorizeError(nil, "")
	}

	return &userID, nil
}
