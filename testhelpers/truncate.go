package testhelpers

// TruncateUsers will truncate the users table
func TruncateUsers() {
	_, err := testConfig.DBConn.Exec("TRUNCATE TABLE users RESTART IDENTITY")

	if err != nil {
		panic(err.Error())
	}
}

// TruncateMessages will truncate the messages table
func TruncateMessages() {
	var err error

	_, err = testConfig.DBConn.Exec("TRUNCATE TABLE messages RESTART IDENTITY")
	if err != nil {
		panic(err.Error())
	}

	_, err = testConfig.DBConn.Exec("TRUNCATE TABLE conversations RESTART IDENTITY")
	if err != nil {
		panic(err.Error())
	}

	_, err = testConfig.DBConn.Exec("TRUNCATE TABLE conversations_users RESTART IDENTITY")
	if err != nil {
		panic(err.Error())
	}

	_, err = testConfig.DBConn.Exec("TRUNCATE TABLE conversations_messages RESTART IDENTITY")
	if err != nil {
		panic(err.Error())
	}
}
