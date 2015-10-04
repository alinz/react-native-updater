package crypto

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/pem"
	"hash"
	"io/ioutil"
	"log"
)

//RSA implementation
type RSA struct {
	privateKey *rsa.PrivateKey
}

func (r *RSA) Decrypt(data []byte) ([]byte, error) {
	var sha1Hash hash.Hash
	var label []byte

	sha1Hash = sha1.New()

	decrypted, err := rsa.DecryptOAEP(sha1Hash, rand.Reader, r.privateKey, data, label)
	if err != nil {
		return nil, err
	}

	return decrypted, nil
}

func NewRSAWithFile(privateKeyPath string) *RSA {
	var pemData []byte
	var err error

	if pemData, err = ioutil.ReadFile(privateKeyPath); err != nil {
		log.Fatalf("Error reading pem file: %s", err)
	}

	return NewRSAWithData(pemData)
}

func NewRSAWithData(privateKeyContent []byte) *RSA {
	var privateKey *rsa.PrivateKey
	var block *pem.Block
	var err error

	if block, _ = pem.Decode(privateKeyContent); block == nil || block.Type != "RSA PRIVATE KEY" {
		log.Fatal("No valid PEM data found")
	}

	if privateKey, err = x509.ParsePKCS1PrivateKey(block.Bytes); err != nil {
		log.Fatalf("Private key can't be decoded: %s", err)
	}

	return &RSA{
		privateKey: privateKey,
	}
}
