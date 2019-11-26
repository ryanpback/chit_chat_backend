package models

import (
	"time"
)

// Conversation is a struct that holds a Conversation (conversations table) ID.
// And a slice of messages
type Conversation struct {
	ConversationID int64      `json:"conversationId"`
	Messages       []*Message `json:"messages"`
}

/*
 * Exported Methods
 */

// CreateConversation adds a new row to the conversations table
func CreateConversation() (int64, time.Time, error) {
	var convID int64
	var convTime time.Time
	const qry = `
		INSERT INTO
			conversations
		DEFAULT VALUES
		RETURNING *;
	`

	row := DBConn.QueryRow(qry)

	err := row.Scan(&convID, &convTime)
	if err != nil {
		return 0, time.Now(), err
	}

	return convID, convTime, nil
}

// HandleConvJoins handles the creation and error handling for the messaging join tables
func HandleConvJoins(convID, messID, userID int64, receiverIds []int64) error {
	convUserExists, err := UserInConversation(userID, convID)
	if err != nil {
		return err
	}

	if convUserExists {
		return nil
	}

	err = createConversationsUsers(convID, userID, receiverIds)
	if err != nil {
		er := DeleteConversationsUsers(convID, userID, receiverIds)

		if er != nil {
			return er
		}

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

	return err
}

// DeleteConversationsUsers deletes a record from the conversations_users table
func DeleteConversationsUsers(convID, userID int64, receiverIds []int64) error {
	const qry = `
		DELETE FROM
			conversations_users
		WHERE
			conversation_id = $1 AND user_id IN $2
	`
	ids := append(receiverIds, userID)

	_, err := DBConn.Exec(qry, convID, ids)

	// skip err != nil because it will either be and error or nil
	return err
}

// ConversationsFindByID finds a conversation between the user and receiver
func ConversationsFindByID(cID int) (*Conversation, error) {
	const qry = `
		SELECT
			*
		FROM
			messages
		WHERE
			conversation_id = $1
		LIMIT
			50
	`

	rows, err := DBConn.Query(qry, cID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	messages := make([]*Message, 0)

	for rows.Next() {
		var m Message

		err := rows.Scan(&m.ID, &m.SenderID, &m.ConversationID, &m.Message, &m.CreatedAt)
		if err != nil {
			return nil, err
		}

		messages = append(messages, &m)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	conversation := Conversation{
		int64(cID),
		messages,
	}

	return &conversation, nil
}

// ConversationsFindByUserID will retreive all messages in all conversations of a particular user
func ConversationsFindByUserID(uID int) ([]*Conversation, error) {
	const qry = `
		SELECT
			*
		FROM (
			SELECT
				m.*,
				ROW_NUMBER() OVER(
					PARTITION BY
						cu.user_id, m.conversation_id
					ORDER BY
						m.created_at DESC
				) rn
			FROM
				messages m
			INNER JOIN
				conversations_users cu
			ON
				cu.conversation_id = m.conversation_id
					AND cu.user_id = $1
		) t
		WHERE
			rn <= 50
		ORDER BY
			conversation_id, created_at DESC;
	`

	rows, err := DBConn.Query(qry, uID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	messages := make(map[int64][]*Message)
	position := 0

	for rows.Next() {
		var m Message

		err := rows.Scan(&m.ID, &m.SenderID, &m.ConversationID, &m.Message, &m.CreatedAt, &position)
		if err != nil {
			return nil, err
		}

		messages[m.ConversationID] = append(messages[m.ConversationID], &m)
	}

	conversations := make([]*Conversation, 0)
	for k, v := range messages {
		conv := Conversation{k, v}
		conversations = append(conversations, &conv)
	}

	return conversations, nil
}

/*
 * Un-exported Methods
 */

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

	DBConn.QueryRow(
		qry,
		convID,
		userID)

	// TODO: THIS IS BAD. DO A BATCH INSERT
	for _, id := range receiverIds {
		DBConn.QueryRow(
			qry,
			convID,
			id)
	}

	return nil
}
