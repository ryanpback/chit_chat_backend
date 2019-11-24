package helpers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateUserMissingFields(t *testing.T) {
	assert := assert.New(t)
	payload := map[string]interface{}{
		"name":         "",
		"userName":     "",
		"email":        "",
		"emailConfirm": "",
		"password":     "",
	}

	ok, err := ValidateUser(payload, GetUserCreateFields())

	// I could assert an error message but that's
	// prone to breaking if the message changes.
	assert.False(ok)
	assert.Equal(len(payload), len(err), "All fields were left blank")
}

func TestValidateUserInvalidEmail(t *testing.T) {
	assert := assert.New(t)
	payload := map[string]interface{}{
		"name":         "Test",
		"userName":     "tess",
		"email":        "testemail",
		"emailConfirm": "t@t.com",
		"password":     "password",
	}

	failed, err := ValidateUser(payload, GetUserCreateFields())
	assert.False(failed)

	_, ok := err[0]["email"]
	assert.Equal(true, ok, "Email error message exists")
}

func TestValidateUserAllValidFields(t *testing.T) {
	assert := assert.New(t)
	payload := map[string]interface{}{
		"name":         "Test",
		"userName":     "tess",
		"email":        "t@t.com",
		"emailConfirm": "t@t.com",
		"password":     "password",
	}

	ok, err := ValidateUser(payload, GetUserCreateFields())

	assert.True(ok)
	assert.Nil(err)
}
