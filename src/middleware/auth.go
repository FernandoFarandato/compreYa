package middleware

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"compreYa/src/core/errors"
	"compreYa/src/infrastructure/entrypoints"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type Validation struct {
}

func (middleware *Validation) Handle(c *gin.Context) {
	entrypoints.ErrorWrapper(middleware.handle, c)
}

func (middleware *Validation) handle(c *gin.Context) *errors.ApiError {
	tokenString, err := middleware.getToken(c)
	if err != nil {
		return errors.NewInternalServerError(nil, "Error getting token")
	}

	token, err := middleware.parseToken(c, tokenString)
	if err != nil {
		return errors.NewInternalServerError(nil, "Error parsing token")
	}

	isAuthorize, err := middleware.validateToken(c, token)
	if !isAuthorize {
		errors.NewNotAuthorizeError(nil, "User is not authenticated")
	}

	return nil
}

func (middleware *Validation) getToken(c *gin.Context) (*string, *errors.ApiError) {
	token, err := c.Cookie("Authorize")
	if token == "" || err != nil {
		return nil, errors.NewNotAuthorizeError(nil)
	}

	return &token, nil
}

func (middleware *Validation) parseToken(c *gin.Context, tokenString *string) (*jwt.Token, *errors.ApiError) {
	token, err := jwt.Parse(*tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("HASHING_KEY")), nil
	})
	if err != nil {
		return nil, errors.NewInternalServerError(nil, "Error parsing token")
	}

	return token, nil
}

func (middleware *Validation) validateToken(c *gin.Context, token *jwt.Token) (bool, *errors.ApiError) {
	var tokenUserID float64
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			return false, errors.NewNotAuthorizeError(nil, "")
		}
		tokenUserID = claims["sub"].(float64)
	} else {
		return false, errors.NewNotAuthorizeError(nil, "")
	}
	requestUserID, _ := strconv.ParseInt(c.GetHeader("X-CALLER-ID"), 10, 64)

	return int64(tokenUserID) == requestUserID, nil
}
