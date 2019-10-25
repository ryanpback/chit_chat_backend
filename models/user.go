package models

import (
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
	UserName  string `json:"user_name"`
	password  string
	CreatedAt time.Time `json:"created_at"`
}

type payload map[string]interface{}

// UserLogin will validate the user login and return the user. TODO: return a session, too.
func UserLogin(p payload) (*User, error) {
	failedLoginMessage := "The email or password you provided does not match any records"
	email := fmt.Sprintf("%v", p["email"])
	user, err := UserFindByEmail(email)
	if err != nil {
		return nil, errors.New(failedLoginMessage)
	}

	providedPassword := fmt.Sprintf("%v", p["password"])
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

	users := make([]*User, 0)

	for rows.Next() {
		var user User

		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.UserName, &user.password, &user.CreatedAt)
		if err != nil {
			return nil, err
		}

		users = append(users, &user)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

// UserFindByEmail will retreive a user by email
func UserFindByEmail(email string) (*User, error) {
	var user User
	const qry = `
		SELECT
			*
		FROM
			users
		WHERE
			email = $1;
	`

	row := DBConn.QueryRow(qry, email)
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.UserName, &user.password, &user.CreatedAt)

	if err == sql.ErrNoRows {
		er := fmt.Errorf(fmt.Sprintf("No user found with email: %v", email))

		return nil, er
	}

	if err != nil {
		return nil, err
	}

	return &user, nil
}

// UserFindByID will return a single user
func UserFindByID(id int) (*User, error) {
	var user User
	const qry = `
		SELECT
			*
		FROM
			users
		WHERE
			id = $1;
	`

	row := DBConn.QueryRow(qry, id)
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.UserName, &user.password, &user.CreatedAt)

	if err == sql.ErrNoRows {
		er := fmt.Errorf(fmt.Sprintf("No user found with ID: %d", id))

		return nil, er
	}

	if err != nil {
		return nil, err
	}

	return &user, nil
}

// UserCreate will insert a new record into the DB
func UserCreate(p payload) (*User, error) {
	var user User
	const qry = `
		INSERT INTO users(name, user_name, email, password)
		VALUES
			($1, $2, $3, $4)
		RETURNING *;
	`

	password := fmt.Sprintf("%v", p["password"])
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
		&user.ID,
		&user.Name,
		&user.Email,
		&user.UserName,
		&user.password,
		&user.CreatedAt)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

// UserEdit will update the user
func UserEdit(u *User, p payload) (*User, error) {
	const qry = `
		UPDATE users
		SET name = $1,
			user_name = $2,
			email = $3
		WHERE id = $4
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
		&u.CreatedAt)

	if err != nil {
		return nil, err
	}

	return u, nil
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
