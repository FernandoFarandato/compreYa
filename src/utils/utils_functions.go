package utils

import (
	"compreYa/src/core/constants"
	"regexp"
)

func ValidateRegex(text string, regex constants.Regex) bool {
	match, _ := regexp.MatchString(regex.String(), text)
	return match
}
