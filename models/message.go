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

// MessagesUser will return the messages belonging
// to the user passed to it
func MessagesUser(userID int) ([]*Message, error) {
	// For each "conversation" (same sender_id and receiver_id in either column),
	// group them and order by created_at DESC - limit to 50 messages per convo
	const qry = `
		SELECT
			id,
			sender_id,
			receiver_id,
			message,
			created_at
		FROM (
			SELECT
				t.*,
				rank() OVER(
					PARTITION BY
						LEAST(sender_id, receiver_id), GREATEST(sender_id, receiver_id)
					ORDER BY
						created_at DESC
				) rnk
			FROM
				messages t
			WHERE
				sender_id = $1 OR receiver_id = $1
		) t
		WHERE
			rnk <= 50
		ORDER BY
			LEAST(sender_id, receiver_id), GREATEST(sender_id, receiver_id), rnk;
	`

	rows, err := DBConn.Query(qry, userID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	messages := make([]*Message, 0)

	for rows.Next() {
		var m Message

		err := rows.Scan(&m.ID, &m.SenderID, &m.ReceiverID, &m.Message, &m.CreatedAt)
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
