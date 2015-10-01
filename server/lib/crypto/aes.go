package crypto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha1"
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"io"
)

func encryptAESCFB(dst, src, key, iv []byte) error {
	aesBlockEncrypter, err := aes.NewCipher([]byte(key))
	if err != nil {
		return err
	}

	aesEncrypter := cipher.NewCFBEncrypter(aesBlockEncrypter, iv)
	aesEncrypter.XORKeyStream(dst, src)
	return nil
}

func decryptAESCFB(dst, src, key, iv []byte) error {
	aesBlockDecrypter, err := aes.NewCipher([]byte(key))
	if err != nil {
		return err
	}

	aesDecrypter := cipher.NewCFBDecrypter(aesBlockDecrypter, iv)
	aesDecrypter.XORKeyStream(dst, src)
	return nil
}

//Encrypt bytes of data with key
func Encrypt(data, key []byte, withHash bool) ([]byte, error) {
	//key must be 16, 24 or 32 bytes
	keyLen := len(key)
	if keyLen != 16 && keyLen != 24 && keyLen != 32 {
		return nil, fmt.Errorf("key length must be 16, 24 or 32 bytes")
	}

	var hash []byte

	if withHash {
		hasher := sha1.New()
		hasher.Write(data)
		hash = hasher.Sum(nil)
		// fmt.Printf("HASH: %v\n", hash)
	}

	encrypted := make([]byte, aes.BlockSize+len(data)+len(hash))
	//randVI(encrypted[:aes.BlockSize])

	io.ReadFull(rand.Reader, encrypted[:aes.BlockSize])

	if _, err := io.ReadFull(rand.Reader, encrypted[:aes.BlockSize]); err != nil {
		return nil, fmt.Errorf("something went wrong with generating new IV")
	}

	var source []byte

	if len(hash) > 0 {
		copy(encrypted[aes.BlockSize:], data)
		copy(encrypted[aes.BlockSize+len(data):], hash)

		source = encrypted[aes.BlockSize:]
	} else {
		source = data
	}

	err := encryptAESCFB(encrypted[aes.BlockSize:], source, key, encrypted[:aes.BlockSize])

	if err != nil {
		return nil, err
	}

	return encrypted, nil
}

//Decrypt bytes of data with key
func Decrypt(data, key []byte, withHash bool) ([]byte, error) {
	//key must be 16, 24 or 32 bytes
	keyLen := len(key)
	if keyLen != 16 && keyLen != 24 && keyLen != 32 {
		return nil, fmt.Errorf("key length must be 16, 24 or 32 bytes")
	}

	decrypted := make([]byte, len(data)-aes.BlockSize)

	err := decryptAESCFB(decrypted, data[aes.BlockSize:], key, data[:aes.BlockSize])

	var hash []byte

	if withHash {
		hasher := sha1.New()
		hasher.Write(decrypted[:len(decrypted)-20])
		hash = hasher.Sum(nil)

		if bytes.Compare(hash, decrypted[len(decrypted)-20:]) != 0 {
			return nil, fmt.Errorf("data is invalid")
		}

		decrypted = decrypted[:len(decrypted)-20]
	}

	if err != nil {
		return nil, err
	}

	return decrypted, nil
}

func EncryptSecureInt64(value int64, key []byte) ([]byte, error) {
	data := make([]byte, 8)

	binary.PutVarint(data, value)

	encrypted, err := Encrypt(data, key, true)

	if err != nil {
		return nil, err
	}

	return encrypted, nil
}

func DecryptSecureInt64(data, key []byte) (int64, error) {
	decrypted, err := Decrypt(data, key, true)

	if err != nil {
		return 0, err
	}

	value, _ := binary.Varint(decrypted)

	return value, nil
}

func EncryptSecureInt64AsBase64(value int64, key []byte) (string, error) {
	encrypted, err := EncryptSecureInt64(value, key)

	if err != nil {
		return "", err
	}

	return base64.URLEncoding.EncodeToString(encrypted), nil
}

func DecryptSecureInt64FromBase64(value string, key []byte) (int64, error) {
	encrypted, err := base64.URLEncoding.DecodeString(value)

	if err != nil {
		return 0, err
	}

	result, err := DecryptSecureInt64(encrypted, key)
	if err != nil {
		return 0, err
	}

	return result, nil
}
