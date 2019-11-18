package testhelpers

import "log"

// MessagePersistToDB and persist to DB
func MessagePersistToDB(sender, receiver int, message string) {
	const qry = `
		INSERT INTO messages(sender_id, receiver_id, message)
		VALUES
			($1, $2, $3)
		`
	_, err := testConfig.DBConn.Exec(qry, sender, receiver, message)

	if err != nil {
		TruncateMessages()
		log.Print(err.Error())
		panic("Something went wrong!")
	}
}

// TruncateMessages will truncate the messages table
func TruncateMessages() {
	_, err := testConfig.DBConn.Exec("TRUNCATE TABLE messages RESTART IDENTITY")

	if err != nil {
		panic(err.Error())
	}
}
