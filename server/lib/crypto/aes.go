package crypto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha1"
	"fmt"
	"math/rand"
	"time"
)

var letters = []byte("1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randVI(b []byte) {
	length := len(letters)
	for i := range b {
		b[i] = letters[rand.Intn(length)]
	}
}

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
	randVI(encrypted[:aes.BlockSize])

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

func init() {
	rand.Seed(time.Now().UnixNano())
}
