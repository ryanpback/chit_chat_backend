package models

import (
	th "chitChat/testhelpers"
	"fmt"
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
		hash, _ := hashAndSaltPassword([]byte(u.password))
		th.UserPersistToDB(u.name, u.userName, u.email, hash)
	}
}

func TestUsersAll(t *testing.T) {
	assert := assert.New(t)
	DBConn = tc.DBConn
	createUsers()
	defer th.TruncateUsers()

	u, err := UsersAll()

	assert.Nil(err)
	assert.Equal(len(users), len(u), "Number of users should be equal to what was saved to the database")
}

func TestUserFindByIDUserNotExists(t *testing.T) {
	assert := assert.New(t)
	DBConn = tc.DBConn
	createUsers()
	defer th.TruncateUsers()

	_, err := UserFindByID(999999)

	assert.NotNil(err)
}

func TestUserFindByIDUserExists(t *testing.T) {
	assert := assert.New(t)
	DBConn = tc.DBConn
	createUsers()
	defer th.TruncateUsers()

	u, _ := UserFindByID(1)

	assert.Equal(int64(1), (*u).ID)
}

func TestUserCreate(t *testing.T) {
	assert := assert.New(t)
	DBConn = tc.DBConn
	userData := map[string]interface{}{
		"name":     "Test",
		"userName": "tee",
		"email":    "test@test.com",
		"password": "password",
	}
	defer th.TruncateUsers()

	u, _ := UserCreate(userData)

	assert.Equal(userData["name"], u.Name)
}

func TestUserEdit(t *testing.T) {
	assert := assert.New(t)
	DBConn = tc.DBConn
	createUsers()
	userData := map[string]interface{}{
		"name":     "Updated Name",
		"userName": "updateUserName",
		"email":    "updated@email.com",
		"password": "thispasswordshouldntupdatewhenediting",
	}
	defer th.TruncateUsers()

	u, _ := UserFindByID(1)
	updatedUser, _ := UserEdit(u, userData)

	assert.Equal(userData["name"], (*updatedUser).Name)
	assert.Equal(userData["userName"], (*updatedUser).UserName)
	assert.Equal(userData["email"], (*updatedUser).Email)
	assert.NotEqual(userData["password"], (*updatedUser).password)
}

func TestFindByEmailUserNotExists(t *testing.T) {
	assert := assert.New(t)
	DBConn = tc.DBConn
	createUsers()
	defer th.TruncateUsers()

	_, err := UserFindByEmail("thisisanon@existentemail.com")

	assert.NotNil(err)
}

func TestFindByEmailUserExists(t *testing.T) {
	assert := assert.New(t)
	DBConn = tc.DBConn
	createUsers()
	defer th.TruncateUsers()

	u, _ := UserFindByEmail(users[0].email)

	assert.Equal(users[0].email, (*u).Email)
}

func TestUserLoginUserNotExist(t *testing.T) {
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

func TestUserLoginPasswordNotMatch(t *testing.T) {
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

func TestUserLoginSuccessful(t *testing.T) {
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
