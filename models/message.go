package models

import (
	"chitChat/helpers"
	"strconv"
	"time"
)

// Message describes the data attributes of a message
type Message struct {
	ID             int64     `json:"id"`
	SenderID       int64     `json:"senderId"`
	ConversationID int64     `json:"conversationId"`
	Message        string    `json:"message"`
	CreatedAt      time.Time `json:"createdAt"`
}

// MessageCreate will persist a new message record into the DB
func MessageCreate(p Payload) (*Payload, error) {
	var message Message
	var err error
	const qry = `
		INSERT INTO
			messages(sender_id, conversation_id, message, created_at)
		VALUES
			($1, $2, $3, $4)
		RETURNING *;
	`
	convID, _ := strconv.Atoi(
		helpers.ConvertInterfaceToString(p["conversationId"]))
	cID := int64(convID)
	// Use conversation created_at timestamp for first message in conversation. The rest will use now.
	cTime := time.Now()

	if cID == 0 {
		cID, cTime, err = CreateConversation()

		if err != nil {
			return nil, err
		}
	}

	row := DBConn.QueryRow(
		qry,
		p["senderId"],
		cID,
		p["message"],
		cTime)

	err = row.Scan(
		&message.ID,
		&message.SenderID,
		&message.ConversationID,
		&message.Message,
		&message.CreatedAt)

	if err != nil {
		return nil, err
	}

	receiverIds := helpers.ConvertReceiverIDs(p["receiverIds"])

	err = HandleConvJoins(
		cID,
		message.ID,
		message.SenderID,
		receiverIds)

	if err != nil {
		_ = DeleteConversation(cID)
		// _ = DeleteMessage(message.ID)
		return nil, err
	}

	response := Payload{
		"conversationId": cID,
		"message":        &message,
	}

	return &response, nil
}

// MessagesUser will return the messages belonging
// to the user passed to it
func MessagesUser(userID int) ([]*Message, error) {
	const qry = ``

	rows, err := DBConn.Query(qry, userID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	messages := make([]*Message, 0)

	for rows.Next() {
		var m Message

		err := rows.Scan(&m.ID, &m.SenderID, &m.Message, &m.CreatedAt)
		if err != nil {
			return nil, err
		}

		messages = append(messages, &m)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return messages, nil
}
