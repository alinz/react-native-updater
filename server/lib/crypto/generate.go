package crypto

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/gob"
	"encoding/pem"
	"fmt"
	"os"
)

func Generate(size int, path string) {
	// generate private key
	privatekey, err := rsa.GenerateKey(rand.Reader, size)

	if err != nil {
		fmt.Println(err.Error)
		os.Exit(1)
	}

	// save private key
	privatekeyfile, err := os.Create(path + "/private.key")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	privatekeyencoder := gob.NewEncoder(privatekeyfile)
	privatekeyencoder.Encode(privatekey)
	privatekeyfile.Close()

	// save public key
	var publickey *rsa.PublicKey
	publickey = &privatekey.PublicKey

	publickeyfile, err := os.Create(path + "/public.key")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	publickeyencoder := gob.NewEncoder(publickeyfile)
	publickeyencoder.Encode(publickey)
	publickeyfile.Close()

	// save private key as pem file
	pemfile, err := os.Create(path + "/private.pem")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var pemkey = &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privatekey),
	}

	err = pem.Encode(pemfile, pemkey)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	pemfile.Close()

	// save public key as pem file
	pemfile, err = os.Create(path + "/public.pem")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	bytes, _ := x509.MarshalPKIXPublicKey(publickey)

	pemkey = &pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: bytes,
	}

	err = pem.Encode(pemfile, pemkey)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	pemfile.Close()
}
