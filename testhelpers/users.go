package testhelpers

import "log"

// PersistUserToDB and persist to DB
func PersistUserToDB(name string, userName string, email string, password string) {
	const qry = `
		INSERT INTO users(name, user_name, email, password)
		VALUES
			($1, $2, $3, $4)
		`
	_, err := testConfig.DBConn.Exec(qry, name, userName, email, password)

	if err != nil {
		TruncateUsers()
		log.Print(err.Error())
		panic("Something went wrong!")
	}
}

// TruncateUsers will truncate the users table
func TruncateUsers() {
	_, err := testConfig.DBConn.Exec("TRUNCATE TABLE users RESTART IDENTITY")

	if err != nil {
		panic(err.Error())
	}
}
