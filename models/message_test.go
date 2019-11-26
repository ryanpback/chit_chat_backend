package models

import (
	"chitChat/helpers"
	th "chitChat/testhelpers"
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

// This seeder works with a 1:1 conversation
func createConversations(user1, user2, count int) int {
	senderID := user1
	receiverID := user2
	convID := 0

	for i := 0; i < count; i++ {
		if i%2 != 0 {
			senderID, receiverID = receiverID, senderID
		}

		message := fmt.Sprintf("User %d, sent user %d a message with an index value of %d", senderID, receiverID, i)

		messageData := Payload{
			"senderId":       senderID,
			"receiverIds":    []int{receiverID},
			"message":        message,
			"conversationId": convID,
		}

		m, err := MessageCreate(messageData)

		if err != nil {
			th.TruncateMessages()
			panic(err.Error())
		}

		convID, _ = strconv.Atoi(
			helpers.ConvertInterfaceToString((*m)["conversationId"]))
	}

	return convID
}

func TestMessageCreate(t *testing.T) {
	assert := assert.New(t)
	DBConn = userTC.DBConn
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

func TestMessageCreateSameConversation(t *testing.T) {
	assert := assert.New(t)
	DBConn = userTC.DBConn
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

	res, _ := MessageCreate(messageData)
	convID := (*res)["conversationId"]

	messageData = map[string]interface{}{
		"senderId":       users[0].ID,
		"receiverIds":    []int{receiverID},
		"message":        "Hello World 2",
		"conversationId": convID,
	}

	res, _ = MessageCreate(messageData)

	assert.Equal(convID, (*res)["conversationId"])
}
