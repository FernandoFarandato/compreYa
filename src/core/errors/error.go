package errors

import (
	"net/http"
	"strings"
)

const (

	// BadRequestMessage is the default message when the input parameters on a request are wrong or it is malformed.
	BadRequestMessage = "Invalid request parameters."
	// ResourceNotFoundMessage is the default message when a requested resource is not available.
	ResourceNotFoundMessage = "Resource not found."
	// ResourceNotOwnedMessage is the default message when a user is requesting for a resource that he doesn't own.
	ResourceNotOwnedMessage = "You are not allowed to access this resource."
	// MethodNotAllowedMessage is the default message when a HTTP verb is forbidden on a resource.
	MethodNotAllowedMessage = "Method not allowed on the current resource."
	// StatusUnavailableForLegalReasonsMessage is the default message when a request over a LockDown user occurs.
	StatusUnavailableForLegalReasonsMessage = "The requested user is not available due to legal reasons"
	// InternalServerErrorMessage is the default message when an unexpected condition occurs.
	InternalServerErrorMessage = "Internal Server Error."

	AuthorizeErrorMeMessage = "Not authorize error."
)

type ApiError struct {
	Status  int      `json:"status"`
	Message string   `json:"message"`
	Err     string   `json:"error"`
	Cause   []string `json:"cause"`
}

func newApiError(code int, message string, err string, causes []string) *ApiError {
	apiError := &ApiError{
		Status:  code,
		Message: message,
		Err:     err,
		Cause:   causes,
	}
	return apiError
}

func NewBadRequest(causes []string, messages ...string) *ApiError {
	message := BadRequestMessage
	if len(messages) > 0 {
		message = strings.Join(messages, " - ")
	}
	return newApiError(http.StatusBadRequest, message, "bad_request", causes)
}

func NewInternalServerError(causes []string, messages ...string) *ApiError {
	message := InternalServerErrorMessage
	if len(messages) > 0 {
		message = strings.Join(messages, " - ")
	}
	return newApiError(http.StatusBadRequest, message, "bad_request", causes)
}

func NewNotAuthorizeError(causes []string, messages ...string) *ApiError {
	message := AuthorizeErrorMeMessage
	if len(messages) > 0 {
		message = strings.Join(messages, " - ")
	}
	return newApiError(http.StatusUnauthorized, message, "bad_request", causes)
}
