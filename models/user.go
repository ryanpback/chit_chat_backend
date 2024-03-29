package models

import (
	"chitChat/helpers"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// User describes the data attributes of a user
type User struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	UserName  string `json:"userName"`
	password  string
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

/*
 * Exported Methods
 */

// UserLogin will validate the user login and return the user. TODO: return a session, too.
func UserLogin(p Payload) (*User, error) {
	failedLoginMessage := "The email or password you provided does not match any records"
	email := helpers.ConvertInterfaceToString(p["email"])
	user, err := UserFindByEmail(email)
	if err != nil {
		return nil, errors.New(failedLoginMessage)
	}

	providedPassword := helpers.ConvertInterfaceToString(p["password"])
	userPassword := user.password

	if match := comparePasswords(providedPassword, userPassword); !match {
		return nil, errors.New(failedLoginMessage)
	}

	return user, nil
}

// UsersAll will retrieve all users in the DB
func UsersAll() ([]*User, error) {
	const qry = `
		SELECT
			*
		FROM
			users;
		`

	rows, err := DBConn.Query(qry)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users, err := buildUsersSlice(rows)
	if err != nil {
		return nil, err
	}

	return users, nil
}

// UserFindByEmail will retreive a user by email
func UserFindByEmail(email string) (*User, error) {
	var u User
	const qry = `
		SELECT
			*
		FROM
			users
		WHERE
			email = $1;
	`

	row := DBConn.QueryRow(qry, email)
	err := row.Scan(
		&u.ID,
		&u.Name,
		&u.Email,
		&u.UserName,
		&u.password,
		&u.CreatedAt,
		&u.UpdatedAt)

	if err == sql.ErrNoRows {
		er := fmt.Errorf(fmt.Sprintf("No user found with email: %v", email))

		return nil, er
	}

	if err != nil {
		return nil, err
	}

	return &u, nil
}

// UserFindByID will return a single user
func UserFindByID(id int) (*User, error) {
	var u User
	const qry = `
		SELECT
			*
		FROM
			users
		WHERE
			id = $1;
	`

	row := DBConn.QueryRow(qry, id)
	err := row.Scan(
		&u.ID,
		&u.Name,
		&u.Email,
		&u.UserName,
		&u.password,
		&u.CreatedAt,
		&u.UpdatedAt)

	if err == sql.ErrNoRows {
		er := fmt.Errorf(fmt.Sprintf("No user found with ID: %d", id))

		return nil, er
	}

	if err != nil {
		return nil, err
	}

	return &u, nil
}

// UserCreate will insert a new record into the DB
func UserCreate(p Payload) (*User, error) {
	var u User
	const qry = `
		INSERT INTO
			users(name, user_name, email, password)
		VALUES
			($1, $2, $3, $4)
		RETURNING *;
	`

	password := helpers.ConvertInterfaceToString(p["password"])
	hash, er := hashAndSaltPassword([]byte(password))
	if er != nil {
		return nil, er
	}

	row := DBConn.QueryRow(
		qry,
		p["name"],
		p["userName"],
		p["email"],
		hash)

	err := row.Scan(
		&u.ID,
		&u.Name,
		&u.Email,
		&u.UserName,
		&u.password,
		&u.CreatedAt,
		&u.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return &u, nil
}

// UserEdit will update the user
func UserEdit(u *User, p Payload) (*User, error) {
	const qry = `
		UPDATE
			users
		SET
			name = $1,
			user_name = $2,
			email = $3,
			updated_at = NOW()
		WHERE
			id = $4
		RETURNING *;
	`

	row := DBConn.QueryRow(
		qry,
		p["name"],
		p["userName"],
		p["email"],
		(*u).ID)

	err := row.Scan(
		&u.ID,
		&u.Name,
		&u.Email,
		&u.UserName,
		&u.password,
		&u.CreatedAt,
		&u.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return u, nil
}

// UserInConversation checks if the user is already in a conversation so we don't create a new record
func UserInConversation(u, c int64) (bool, error) {
	const qry = `
		SELECT
			COUNT(*)
		FROM
			conversations_users
		WHERE
			user_id = $1 AND conversation_id = $2
	`
	var count int

	row := DBConn.QueryRow(qry, u, c)
	err := row.Scan(&count)
	if err != nil {
		return false, err
	}

	if count < 1 {
		return false, nil
	}

	return true, nil
}

// UsersTypeahead will find users where the email contains the text passed in
func UsersTypeahead(s string) ([]*User, error) {
	const qry = `
		SELECT
			*
		FROM
			users
		WHERE
			email ILIKE $1
		LIMIT
			10;
	`

	rows, err := DBConn.Query(qry, "%"+s+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users, err := buildUsersSlice(rows)
	if err != nil {
		return nil, err
	}

	if len(users) < 1 {
		err = fmt.Errorf(fmt.Sprintf("No users containing email: '%v'", s))
	}

	return users, err
}

/*
 * Un-exported Methods
 */

func buildUsersSlice(r *sql.Rows) ([]*User, error) {
	users := make([]*User, 0)

	for r.Next() {
		var u User

		err := r.Scan(&u.ID, &u.Name, &u.Email, &u.UserName, &u.password, &u.CreatedAt, &u.UpdatedAt)
		if err != nil {
			return nil, err
		}

		users = append(users, &u)
	}

	if err := r.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func hashAndSaltPassword(b []byte) (string, error) {
	hash, err := bcrypt.GenerateFromPassword(b, 5)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func comparePasswords(plainText string, hashedPassword string) bool {
	userPass := []byte(hashedPassword)
	provided := []byte(plainText)

	err := bcrypt.CompareHashAndPassword(userPass, provided)
	if err != nil {
		return false
	}

	return true
}
