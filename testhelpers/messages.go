package testhelpers

// MessagePersistToDB and persist to DB
func MessagePersistToDB(sender int, message string) {
	const qry = `
		INSERT INTO
			messages(sender_id, message)
		VALUES
			($1, $2)
		`
	_, err := testConfig.DBConn.Exec(qry, sender, message)

	if err != nil {
		TruncateMessages()
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
