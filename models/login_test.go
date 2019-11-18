package models

import (
	th "chitChat/testhelpers"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var loginTC = th.BootstrapTestConfig()

func TestLoginUserNotExist(t *testing.T) {
	assert := assert.New(t)
	DBConn = loginTC.DBConn
	defer th.TruncateUsers()

	createUsers()
	loginData := map[string]interface{}{
		"email":    "thisuser@doesnotexist.com",
		"password": "password",
	}

	_, err := UserLogin(loginData)

	assert.NotNil(err)
}

func TestLoginPasswordNotMatch(t *testing.T) {
	assert := assert.New(t)
	DBConn = loginTC.DBConn
	defer th.TruncateUsers()

	createUsers()
	loginData := map[string]interface{}{
		"email":    users[0].email,
		"password": fmt.Sprintf("makethisfake%v", users[0].password),
	}

	_, err := UserLogin(loginData)

	assert.NotNil(err)
}

func TestLoginSuccessful(t *testing.T) {
	assert := assert.New(t)
	DBConn = loginTC.DBConn
	defer th.TruncateUsers()

	createUsers()
	loginData := map[string]interface{}{
		"email":    users[0].email,
		"password": "password",
	}

	u, _ := UserLogin(loginData)

	assert.Equal(int64(1), (*u).ID)
}
