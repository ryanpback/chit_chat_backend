package models

import (
	"database/sql"
	"fmt"
	"time"
)

// User describes the data attributes of a user
type User struct {
	ID          int64
	Name        string
	DisplayName string
	password    string
	CreatedAt   time.Time
}

// GetAllUsers will retrieve all users in the DB
func GetAllUsers() ([]*User, error) {
	const qry = `
		SELECT
			*
		FROM
			users
		`

	rows, err := db.Query(qry)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	users := make([]*User, 0)

	for rows.Next() {
		var user User

		err := rows.Scan(&user.ID, &user.Name, &user.DisplayName, &user.password, &user.CreatedAt)
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
			id = $1
	`

	row := db.QueryRow(qry, id)
	err := row.Scan(&user.ID, &user.Name, &user.DisplayName, &user.password, &user.CreatedAt)

	if err == sql.ErrNoRows {
		er := fmt.Errorf(fmt.Sprintf("No user found with ID: %d", id))

		return nil, er
	}

	if err != nil {
		return nil, err
	}

	return &user, nil
}
