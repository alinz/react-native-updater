package data_test

import (
	"encoding/json"
	"testing"

	"github.com/alinz/react-native-updater/server/data"
	"github.com/stretchr/testify/assert"
)

func TestID(t *testing.T) {
	data.SetSecureIDKey("123456789012345678901234")

	type internal struct {
		ID data.SecureID
	}

	sendMessage := internal{
		ID: 10,
	}

	b, err := json.Marshal(sendMessage)
	if err != nil {
		assert.Equal(t, nil, err, err)
	}

	receiveMessage := internal{}

	err = json.Unmarshal(b, &receiveMessage)
	if err != nil {
		assert.Equal(t, nil, err, err)
	}

	assert.Equal(t, int64(10), int64(receiveMessage.ID), "should be the same value")
}
