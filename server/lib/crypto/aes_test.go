package crypto_test

import (
	"testing"

	"github.com/alinz/react-native-updater/server/lib/crypto"
	"github.com/stretchr/testify/assert"
)

func TestWithValidation(t *testing.T) {
	data := []byte("hello")
	key := []byte("123456789012345678901234")

	encrypted, err := crypto.Encrypt(data, key, true)

	if err != nil {
		panic(err)
	}

	decrypted, err := crypto.Decrypt(encrypted, key, true)

	assert.Equal(t, "hello", string(decrypted), "should be the same")
	assert.Equal(nil, err, err)
}

func TestWithoutValidation(t *testing.T) {
	data := []byte("hello")
	key := []byte("123456789012345678901234")

	encrypted, err := crypto.Encrypt(data, key, false)

	if err != nil {
		panic(err)
	}

	decrypted, err := crypto.Decrypt(encrypted, key, false)

	assert.Equal(t, "hello", string(decrypted), "should be the same")
	assert.Equal(nil, err, err)
}
