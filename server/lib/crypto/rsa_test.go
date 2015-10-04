package crypto_test

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testData = []struct {
	privateKey       []byte
	passphrase       string
	encryptedMessage string
	decryptedMessage string
}{
	{
		privateKey: []byte(`-----BEGIN RSA PRIVATE KEY-----
Proc-Type: 4,ENCRYPTED
DEK-Info: DES-EDE3-CBC,41744071146EB9E0

ijDP1crU7Lj51m0UQTcvvqPUEJfVYZus00Oe1cbv5BWFE6u7WzcP+kIaZElSfmOV
03+UXvtKYXoUPA8YhJEOB2hewoVO5i+BIZNGPxnVZVy/KFxzGemHQqUmItZSYtYB
wDoZR4zSirGCR5sB3lNb31EPAAWZNqnfhvVqeXi4qndBmla6iDYtCrqg7GTKe8/w
8c+xDjFfGbUW5uMNtT3zXHGgbMCD+QprXkpLehxu5ORssppHlOM4sYZZkyxZexGZ
RXGnnVVe5vy6xOXT84s91O9K8ixWyxeKC48XVjGvmerXg1H3r1MbueKy1NM8DgLG
N6twpoyOUYytljXPnCC/7ou4unXh0eNig4GYOsaOf1hULCmRniWcZMzn61B9k3ip
W6LKv7xDCenD1XkavcL9adWgMsvkC6o/uGcUcpj6Y6aM34kwVC7uIRUYHJkLBHjX
g17fj7V/6Qzoj3cOxCbmr/8lhJFU6jFfXMZylxdBqsovzciR6QQU01ISmf3NcC/h
g4aBpZ3rpXzUhoELomH4DzJzXKknZUHPrBhO1YSBro7RlmXCSYK8aVHvn1X+lp2y
uRAXki+Vqhh0o8axCQbWIT622Wnv50Gd2hEyrgixuvugCqq4EPvnVe7813j73iUg
PapcybNWxQWSkCQhYdZYzP5jrJhaWPW0jDNrSA1IEih7KOscaBRa2OYqG20cLerC
dSBDt4m4M6UDgOa5tBiYuVADKzBar1jOeDu5oWw+x4qfG7/7vLyZALS8r5tcJn1k
SVAOAE/nQMAkjrm4cdKj4FAOOSrNyZcBBFlLNOD4EmPqT9ex5ZoogTwYiHY3VOy5
IhztCycpoCLOOzEOHj5P9fuMCEVS/4vKQRjpVmCKYCLyeEhLNcK6Q491wKDilJAV
HMxNxBJA08D1dPYeA7wI0v2CK6io3KMPihDTMcxSDT5CNQmJisc7Y0onr6L07MJF
RV0hsIwCaPf3pT2re8psqd/v1NkH0Unr5w4VboBfHzo5821QnBKrlAJtk6W4dJ/5
PxUlZY+qA1y0+zYf2jbRWjU1lpe6ls1vfhf4rml6Ob3tsYI20iFoEyo7eauGcGQ0
CkqGgk0OtGjvFged8z27/UBDX3jW/2YT7jM8KTIlC7RjjqqfUJk1DTY2e/+l4ndi
OcdcVzTrmKqUiT/VrYE4lVEaI1CRsesQN1efeV+G3q57K1N/2TlpUGmnWCODhMkc
vbe+8TP1TJNbbJ5EYxULHIPC31p+rbK9/Z/hQeMQdVDkAUmCiKXLZBYEKeZ8xLEG
UN9AeWacOBeKGX4szqhKrWsCgARtw5WtKH8qosLa+jJlgq3IwH9q2XVL4BsRQ3ds
scSYW3k6T9PUqX8PEEhfP7PTjrvBNc+k3i5+lGJVvCEZXaACO0WCLsixXXlQqy/J
V6I1Pi09HS03p+YlsKZWRx4zAk+YBNhJRl8cc9Z4aj0tJYRI6pcq9vmw4/ZlOSOW
xLMLCmHgGaMoKCFcK2m8owOcR0oq5H98a8+gxxr/YnU3SMYZRsD0hds17G33BYcH
h7pwRWj9VGyrjV9umBlZN8OidE30U15Cz5oDpdO1Vzs55s1qyN65fA==
-----END RSA PRIVATE KEY-----
`),
		passphrase:       "test",
		encryptedMessage: "r9ZWuqNWyjZEKLKInLVzi4KqoSN9HF8W2lMw1Q1brYNZ7kBIvfHT8wKjh37PHP9ac01tFeg0bT6bXK0OwDmAyxVAILp/tMGQwUGRMVKF9w+hFni+kAvYK6Sr2K8aXwWZXL/9k2b0f+KHBMSgGIvNM0L+u6iY/G2RVKvWZzptdLoZ49IHA249G+k65cvS1a2lM11Vg5SslYyfAJvLpi+cWeJmsvAxwwrd1hfT5bSfZjprGaEFNStVT5IZGHQX53oR79zcPqFDEK+gAhsFWmb7ZjsCL/MAoxNq1E2/1fYYTzzB7b2lz3wxL1ZHnO/mWpnz7wBoFOsqFHNp/LMOY0VrvQ==",
		decryptedMessage: "Hello world",
	},
}

func TestDecrypt(t *testing.T) {
	block, rest := pem.Decode(testData[0].privateKey)

	if len(rest) > 0 {
		t.Error("extra data")
	}

	der, err := x509.DecryptPEMBlock(block, []byte(testData[0].passphrase))

	if err != nil {
		t.Error("decrypt failed: ", err)
	}

	var privateKey *rsa.PrivateKey
	if privateKey, err = x509.ParsePKCS1PrivateKey(der); err != nil {
		t.Error("private key failed: ", err)
	}

	b, _ := base64.StdEncoding.DecodeString(testData[0].encryptedMessage)

	decrypted, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, b)

	if err != nil {
		t.Error("failed to decrypt, ", err)
	}

	assert.Equal(t, testData[0].decryptedMessage, string(decrypted), "should be the same")
}
