package models

import (
	th "chitChat/testhelpers"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func createMessages(user1, user2, count int) {
	senderID := user1
	receiverID := user2

	for i := 0; i < count; i++ {
		if i%2 != 0 {
			senderID, receiverID = receiverID, senderID
		}

		message := fmt.Sprintf("User %d, sent user %d a message with an index value of %d", senderID, receiverID, i)

		th.MessagePersistToDB(senderID, receiverID, message)
	}
}

func TestMessageCreate(t *testing.T) {
	assert := assert.New(t)
	DBConn = userTC.DBConn
	defer th.TruncateUsers()
	defer th.TruncateMessages()

	createUsers()
	users, _ := UsersAll()
	messageData := map[string]interface{}{
		"senderId":   users[0].ID,
		"receiverId": users[1].ID,
		"message":    "Hello World",
	}

	_, err := MessageCreate(messageData)

	assert.Nil(err)
}

func TestMessagesUser(t *testing.T) {
	assert := assert.New(t)
	DBConn = userTC.DBConn
	defer th.TruncateUsers()
	defer th.TruncateMessages()
	createUsers()
	users, _ := UsersAll()

	createMessages(int(users[0].ID), int(users[1].ID), 2)
	createMessages(int(users[0].ID), int(users[2].ID), 2)
	createMessages(int(users[2].ID), int(users[1].ID), 2)
	createMessages(int(users[0].ID), int(users[2].ID), 2)
	createMessages(int(users[0].ID), int(users[1].ID), 2)
	createMessages(int(users[1].ID), int(users[2].ID), 2)

	userMessages, err := MessagesUser(int(users[0].ID))

	assert.Nil(err)
	fmt.Println(userMessages)

	// TODO after this commit - test that user messages are properly
	// grouped and grouped into multi-dimensional array
}
