package models

import (
	th "chitChat/testhelpers"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConversationsFindByID(t *testing.T) {
	assert := assert.New(t)
	DBConn = userTC.DBConn
	defer th.TruncateUsers()
	defer th.TruncateMessages()

	const messageCount = 7

	createUsers()
	users, _ := UsersAll()

	convID := createConversations(int(users[0].ID), int(users[1].ID), messageCount)
	_ = createConversations(int(users[0].ID), int(users[2].ID), 4)

	conv, err := ConversationsFindByID(convID)

	assert.Nil(err)
	assert.Equal(messageCount, len(conv.Messages))
}

func TestConversationsFindByUserID(t *testing.T) {
	assert := assert.New(t)
	DBConn = userTC.DBConn
	defer th.TruncateUsers()
	defer th.TruncateMessages()

	const messageCount = 3
	const messageCount2 = 5

	createUsers()
	users, _ := UsersAll()

	_ = createConversations(int(users[0].ID), int(users[1].ID), messageCount)
	_ = createConversations(int(users[0].ID), int(users[2].ID), messageCount2)
	_ = createConversations(int(users[1].ID), int(users[2].ID), messageCount)

	convs, err := ConversationsFindByUserID(int(users[0].ID))

	assert.Nil(err)
	assert.Equal(2, len(convs))
}
