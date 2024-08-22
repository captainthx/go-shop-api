package utils

import (
	"errors"
	"regexp"
	"unicode/utf8"
)

var usernameRegex = regexp.MustCompile(`^[a-zA-Z0-9_]+$`)
var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

func validate(value, fieldname string, regex *regexp.Regexp, min, max int) (bool, error) {
	if value == "" {
		return true, errors.New(fieldname + " must not be empty")
	}
	if !regex.MatchString(value) {
		return true, errors.New(fieldname + " must contain only letters, numbers, or underscores")
	}
	if crossRange(value, min, max) {
		return true, errors.New(fieldname + " must be between " + string(rune(min)) + " and " + string(rune(max)) + " characters")
	}
	return false, nil
}

func InvalidUsername(username string) (bool, error) {
	return validate(username, "Username", usernameRegex, 4, 20)
}

func InvalidEmail(email string) (bool, error) {
	return validate(email, "Email", emailRegex, 4, 50)
}

func InvalidName(name string) (bool, error) {
	return validate(name, "Name", usernameRegex, 4, 30)
}

func InvalidPassword(password string) (bool, error) {
	return validate(password, "Password", usernameRegex, 4, 255)
}

func crossRange(str string, min, max int) bool {
	length := utf8.RuneCountInString(str)
	return length < min || length > max
}
