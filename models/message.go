package models

import (
	"chitChat/helpers"
	"chitChat/services"
	"strconv"
	"time"
)

// Message describes the data attributes of a message
type Message struct {
	ID        int64     `json:"id"`
	SenderID  int64     `json:"senderId"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"createdAt"`
}

// MessageCreate will persist a new message record into the DB
func MessageCreate(p Payload) (Payload, error) {
	var message Message
	const qry = `
		INSERT INTO
			messages(sender_id, message)
		VALUES
			($1, $2)
		RETURNING *;
	`

	row := DBConn.QueryRow(
		qry,
		p["senderId"],
		p["message"])

	err := row.Scan(
		&message.ID,
		&message.SenderID,
		&message.Message,
		&message.CreatedAt)

	if err != nil {
		return nil, err
	}

	convID, _ := strconv.Atoi(
		helpers.ConvertInterfaceToString(p["conversationId"]))
	cID := int64(convID)
	receiverIds := helpers.ConvertReceiverIDs(p["receiverIds"])

	cID, err = services.HandleConversation(
		cID,
		message.ID,
		message.SenderID,
		message.CreatedAt,
		receiverIds)

	if err != nil {
		return nil, err
	}

	response := Payload{
		"conversationId": cID,
		"message":        &message,
	}

	return response, nil
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
