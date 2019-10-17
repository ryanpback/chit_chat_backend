package models

import (
	th "chitChat/testhelpers"
	"testing"

	"github.com/stretchr/testify/assert"
)

var tc = th.BootstrapTestConfig()

var users = []struct {
	name     string
	userName string
	email    string
	password string
}{
	{
		name:     "test",
		userName: "tee",
		email:    "test1@test.com",
		password: "password",
	},
	{
		name:     "test",
		userName: "tee",
		email:    "test2@test.com",
		password: "password",
	},
	{
		name:     "test",
		userName: "tee",
		email:    "test3@test.com",
		password: "password",
	},
}

func createUsers() {
	for _, u := range users {
		th.CreateUser(u.name, u.userName, u.email, u.password)
	}
}

func TestGetAllUsers(t *testing.T) {
	assert := assert.New(t)
	DBConn = tc.DBConn
	createUsers()
	defer th.TruncateUsers()

	u, err := GetAllUsers()

	assert.Nil(err)
	assert.Equal(len(users), len(u), "Number of users should be equal to what was saved to the database")
}

func TestGetUserByIDUserNotExist(t *testing.T) {
	assert := assert.New(t)
	DBConn = tc.DBConn
	createUsers()
	defer th.TruncateUsers()

	_, err := GetUserByID(999999)

	assert.NotNil(err)
}

func TestGetUserByIDUserExists(t *testing.T) {
	assert := assert.New(t)
	DBConn = tc.DBConn
	createUsers()
	defer th.TruncateUsers()

	u, _ := GetUserByID(1)

	assert.Equal(int64(1), (*u).ID)
}

func TestCreateUser(t *testing.T) {
	assert := assert.New(t)
	DBConn = tc.DBConn
	userData := map[string]interface{}{
		"name":     "Test",
		"userName": "tee",
		"email":    "test@test.com",
		"password": "password",
	}
	defer th.TruncateUsers()

	u, _ := CreateUser(userData)

	assert.Equal(userData["name"], u.Name)
}
