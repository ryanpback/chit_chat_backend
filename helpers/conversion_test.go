package helpers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertReceiverIDs(t *testing.T) {
	assert := assert.New(t)
	ids := map[string]interface{}{
		"ids": []int{1, 2, 3, 4, 5},
	}

	newIDs := ConvertReceiverIDs(ids["ids"])

	assert.Equal(int64(1), newIDs[0])
}
