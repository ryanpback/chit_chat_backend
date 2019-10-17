package models

import (
	"database/sql"
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// User describes the data attributes of a user
type User struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	UserName  string `json:"user_name"`
	Email     string `json:"email"`
	password  string
	CreatedAt time.Time `json:"created_at"`
}

// GetAllUsers will retrieve all users in the DB
func GetAllUsers() ([]*User, error) {
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

// GetUserByID will return a single user
func GetUserByID(id int) (*User, error) {
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
	err := row.Scan(&user.ID, &user.Name, &user.UserName, &user.Email, &user.password, &user.CreatedAt)

	if err == sql.ErrNoRows {
		er := fmt.Errorf(fmt.Sprintf("No user found with ID: %d", id))

		return nil, er
	}

	if err != nil {
		return nil, err
	}

	return &user, nil
}

// CreateUser will insert a new record into the DB
func CreateUser(payload map[string]interface{}) (*User, error) {
	var user User
	const qry = `
		INSERT INTO users(name, user_name, email, password)
		VALUES
			($1, $2, $3, $4)
		RETURNING *;
	`

	password := fmt.Sprintf("%v", payload["password"])
	hash, er := hashAndSaltPassword([]byte(password))
	if er != nil {
		return nil, er
	}

	row := DBConn.QueryRow(
		qry,
		payload["name"],
		payload["userName"],
		payload["email"],
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

func hashAndSaltPassword(b []byte) (string, error) {
	hash, err := bcrypt.GenerateFromPassword(b, 5)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}
