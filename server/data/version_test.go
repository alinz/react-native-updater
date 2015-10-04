package data_test

import (
	"encoding/json"
	"testing"

	"github.com/alinz/react-native-updater/server/data"
	"github.com/stretchr/testify/assert"
)

func TestVersion(t *testing.T) {
	type internal struct {
		Version data.Version
	}

	b := internal{}
	receiveMessage := []byte(`{"Version":"1.2.10"}`)

	err := json.Unmarshal(receiveMessage, &b)
	if err != nil {
		assert.Equal(t, nil, err, err)
	}

	assert.Equal(t, uint64(b.Version), uint64(281483566645258), "should be the same")

	parsedMessage, err := json.Marshal(b)
	if err != nil {
		assert.Equal(t, nil, err, err)
	}

	assert.Equal(t, `{"Version":"1.2.10"}`, string(parsedMessage), "should be the same json")
}
