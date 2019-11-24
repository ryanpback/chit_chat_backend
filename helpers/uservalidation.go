package helpers

import (
	"fmt"
	"regexp"
	"strings"
)

// GetUserLoginFields returns a slice of string containing the fields to be validated for user login
func GetUserLoginFields() []string {
	fields := []string{"email", "password"}

	return fields
}

// GetUserCreateFields returns a slice of string containing the fields to be validated for user creation
func GetUserCreateFields() []string {
	fields := []string{"name", "userName", "email", "emailConfirm", "password"}

	return fields
}

// GetUserEditFields returns a slice of string containing the fields to be validated for user editing
func GetUserEditFields() []string {
	fields := []string{"name", "userName", "email"}

	return fields
}

// ValidateUser validates the payload for the User model
func ValidateUser(p map[string]interface{}, f []string) (bool, []map[string]string) {
	var errors []map[string]string

	for _, keyString := range f {
		if p[keyString] == "" {
			errors = append(errors, map[string]string{keyString: fmt.Sprintf("The '%s' field is required.", strings.ToTitle(keyString))})

			continue
		}

		// Email is present - is it in valid email format?
		if keyString == "email" {
			re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

			if !re.MatchString(p["email"].(string)) {
				errors = append(errors, map[string]string{"email": "The 'Email' field must have a valid email format."})
			}

			continue
		}

		if keyString == "emailConfirm" && p[keyString] != p["email"] {
			errors = append(errors, map[string]string{"emailConfirm": "The 'Confirm Email' value does not match the 'Email' field value."})
		}
	}

	if len(errors) > 0 {
		return false, errors
	}

	return true, nil
}
