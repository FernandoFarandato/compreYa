package utils

import (
	"compreYa/src/core/constants"
	"compreYa/src/core/errors"
	"golang.org/x/crypto/bcrypt"
	"regexp"
)

func ValidateRegex(text string, regex constants.Regex) bool {
	match, _ := regexp.MatchString(regex.String(), text)
	return match
}

func EncryptString(token string) ([]byte, *errors.ApiError) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(token), 10)
	if err != nil {
		return nil, errors.NewInternalServerError(nil, "Error encrypting")
	}

	return hashPassword, nil
}

func VerifyEncryptedString(hashedString, string string) *errors.ApiError {
	err := bcrypt.CompareHashAndPassword([]byte(hashedString), []byte(string))
	if err != nil {
		return errors.NewInternalServerError(nil, errors.CredentialsNotValid)
	}

	return nil
}
