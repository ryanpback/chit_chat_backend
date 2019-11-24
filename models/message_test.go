package models

import (
	"chitChat/services"
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

		th.MessagePersistToDB(senderID, message)
	}
}

func TestMessageCreate(t *testing.T) {
	assert := assert.New(t)
	DBConn = userTC.DBConn
	services.DBConn = userTC.DBConn
	defer th.TruncateUsers()
	defer th.TruncateMessages()

	createUsers()
	users, _ := UsersAll()
	receiverID := int(users[1].ID)
	messageData := map[string]interface{}{
		"senderId":       users[0].ID,
		"receiverIds":    []int{receiverID},
		"message":        "Hello World",
		"conversationId": "",
	}

	_, err := MessageCreate(messageData)

	assert.Nil(err)
}
