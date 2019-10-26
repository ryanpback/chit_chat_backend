package models

import (
	th "chitChat/testhelpers"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoginUserNotExist(t *testing.T) {
	assert := assert.New(t)
	DBConn = tc.DBConn
	createUsers()
	loginData := map[string]interface{}{
		"email":    "thisuser@doesnotexist.com",
		"password": "password",
	}
	defer th.TruncateUsers()

	_, err := UserLogin(loginData)

	assert.NotNil(err)
}

func TestLoginPasswordNotMatch(t *testing.T) {
	assert := assert.New(t)
	DBConn = tc.DBConn
	createUsers()
	loginData := map[string]interface{}{
		"email":    users[0].email,
		"password": fmt.Sprintf("makethisfake%v", users[0].password),
	}
	defer th.TruncateUsers()

	_, err := UserLogin(loginData)

	assert.NotNil(err)
}

func TestLoginSuccessful(t *testing.T) {
	assert := assert.New(t)
	DBConn = tc.DBConn
	createUsers()
	loginData := map[string]interface{}{
		"email":    users[0].email,
		"password": "password",
	}
	defer th.TruncateUsers()

	u, _ := UserLogin(loginData)

	assert.Equal(int64(1), (*u).ID)
}
