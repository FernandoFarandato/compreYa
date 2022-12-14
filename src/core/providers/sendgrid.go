package providers

import (
	"compreYa/src/core/entities"
	"compreYa/src/core/errors"
)

type SendGrid interface {
	SendLoginEmail(userData *entities.User) *errors.ApiError
	SendPasswordRecovery(userData *entities.User, token string) *errors.ApiError
}
