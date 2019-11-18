package models

import (
	"time"
)

// Message describes the data attributes of a message
type Message struct {
	ID         int64     `json:"id"`
	SenderID   int64     `json:"senderId"`
	ReceiverID int64     `json:"receiverId"`
	Message    string    `json:"message"`
	CreatedAt  time.Time `json:"createdAt"`
}

// MessageCreate will persist a new message record into the DB
func MessageCreate(p payload) (*Message, error) {
	var message Message
	const qry = `
		INSERT INTO messages(sender_id, receiver_id, message)
		VALUES
			($1, $2, $3)
		RETURNING *;
	`

	row := DBConn.QueryRow(
		qry,
		p["senderId"],
		p["receiverId"],
		p["message"])

	err := row.Scan(
		&message.ID,
		&message.SenderID,
		&message.ReceiverID,
		&message.Message,
		&message.CreatedAt)

	if err != nil {
		return nil, err
	}

	return &message, nil
}
