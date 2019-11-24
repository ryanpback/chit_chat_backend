package services

import (
	"time"
)

/*
 * Exported Methods
 */

// HandleConversation determines the whether a new conversation
// record needs to be created or just the relationship
func HandleConversation(convID, messID, userID int64, messTime time.Time, receiverIds []int64) (int64, error) {
	var err error
	if convID == 0 {
		convID, err = CreateConversation(messID, userID, messTime)

		if err != nil {
			return 0, err
		}
	}

	err = HandleConvJoins(convID, messID, userID, receiverIds)
	if err != nil {
		er := DeleteConversation(convID)
		if er != nil {
			return 0, er
		}

		return 0, err
	}

	return convID, nil
}

// CreateConversation adds a new row to the conversations table
func CreateConversation(messID, userID int64, messTime time.Time) (int64, error) {
	var convID int64
	const qry = `
		INSERT INTO
			conversations(message_id, created_at)
		VALUES
			($1, $2)
		RETURNING id;
	`

	row := DBConn.QueryRow(
		qry,
		messID,
		messTime)

	err := row.Scan(&convID)
	if err != nil {
		return 0, err
	}

	return convID, nil
}

// HandleConvJoins handles the creation and error handling for the messaging join tables
func HandleConvJoins(convID, messID, userID int64, receiverIds []int64) error {
	cmID, err := createConversationsMessages(convID, messID)
	if err != nil {
		er := DeleteConversationsMessage(cmID)
		if er != nil {
			return er
		}

		return err
	}

	err = createConversationsUsers(convID, userID, receiverIds)
	if err != nil {
		// Due to the order of how joins are created, if
		// this fails there's no record to delete

		// er := DeleteConversationsUser(cuID)
		// if er != nil {
		// 	return 0, er
		// }

		return err
	}

	return nil
}

// DeleteConversation deletes a record from the conversations table
func DeleteConversation(convID int64) error {
	const qry = `
		DELETE FROM
			conversations
		WHERE
			id = $1
	`

	_, err := DBConn.Exec(qry, convID)

	// skip err != nil because it will either be and error or nil
	return err
}

// DeleteConversationsMessage deletes a record from the conversations_messages table
func DeleteConversationsMessage(convID int64) error {
	const qry = `
		DELETE FROM
			conversations_messages
		WHERE
			id = $1
	`

	_, err := DBConn.Exec(qry, convID)

	// skip err != nil because it will either be and error or nil
	return err
}

// DeleteConversationsUser deletes a record from the conversations table
func DeleteConversationsUser(convID int64) error {
	const qry = `
		DELETE FROM
			conversations_users
		WHERE
			id = $1
	`

	_, err := DBConn.Exec(qry, convID)

	// skip err != nil because it will either be and error or nil
	return err
}

/*
 * Un-exported Methods
 */

// CreateConversationsMessages adds a new row to the conversations_messages table
func createConversationsMessages(convID, messID int64) (int64, error) {
	var cmID int64
	const qry = `
		INSERT INTO
			conversations_messages(conversation_id, message_id)
		VALUES
			($1, $2)
		RETURNING id;
	`

	row := DBConn.QueryRow(
		qry,
		convID,
		messID)

	err := row.Scan(&cmID)
	if err != nil {
		return 0, err
	}

	return cmID, nil
}

// CreateConversationsUsers adds a new row to the conversations_users table
func createConversationsUsers(convID, userID int64, receiverIds []int64) error {
	// var cuID int64
	// const qry = `
	// 	INSERT INTO
	// 		conversations_users(conversation_id, user_id)
	// 	SELECT
	// 		$1, user_id
	// 	FROM
	// 		unnest($2);
	// `
	const qry = `
		INSERT INTO
			conversations_users(conversation_id, user_id)
		VALUES
			($1, $2)
	`

	// TODO: THIS IS BAD. DO A BATCH INSERT
	for _, id := range receiverIds {
		DBConn.QueryRow(
			qry,
			convID,
			id)
	}

	return nil
}
