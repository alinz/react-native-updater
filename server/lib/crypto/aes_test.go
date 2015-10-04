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

	assert.Equal(t, err, nil, err)

	decrypted, err := crypto.Decrypt(encrypted, key, true)

	assert.Equal(t, "hello", string(decrypted), "should be the same")
	assert.Equal(nil, err, err)
}

func TestWithoutValidation(t *testing.T) {
	data := []byte("hello")
	key := []byte("123456789012345678901234")

	encrypted, err := crypto.Encrypt(data, key, false)

	assert.Equal(t, err, nil, err)

	decrypted, err := crypto.Decrypt(encrypted, key, false)

	assert.Equal(t, "hello", string(decrypted), "should be the same")
	assert.Equal(nil, err, err)
}

func TestDecryptSecureInt64(t *testing.T) {
	value := int64(876545)
	key := []byte("123456789012345678901234")

	encrypted, err := crypto.EncryptSecureInt64(value, key)

	assert.Equal(t, err, nil, err)

	decryptedValue, err := crypto.DecryptSecureInt64(encrypted, key)

	assert.Equal(nil, err, err)
	assert.Equal(t, value, decryptedValue, "should be the same")
}

func TestDecryptSecureInt64Base64(t *testing.T) {
	value := int64(876545)
	key := []byte("123456789012345678901234")

	encryptedBase64, err := crypto.EncryptSecureInt64AsBase64(value, key)

	assert.Equal(t, err, nil, err)

	decryptedValue, err := crypto.DecryptSecureInt64FromBase64(encryptedBase64, key)

	assert.Equal(nil, err, err)
	assert.Equal(t, value, decryptedValue, "should be the same")
}
