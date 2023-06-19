package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"

	"github.com/viniciusataide/velozient-challenge-go/server"
)

type Crypto struct {
	cipher.AEAD
}

func ProvideCrypto(config *server.Config) *Crypto {
	c, err := aes.NewCipher([]byte(config.AesKey))

	gcm, err := cipher.NewGCM(c)

	if err != nil {
		panic(err)
	}

	return &Crypto{gcm}
}

func (gcm *Crypto) EncryptAES(plaintext string) string {
	nonce := make([]byte, gcm.NonceSize())

	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err)
	}

	a := gcm.Seal(nonce, nonce, []byte(plaintext), nil)

	return string(a)
}

func (gcm *Crypto) DecryptAES(ct string) string {
	ciphertext := []byte(ct)
	nonceSize := gcm.NonceSize()

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)

	if err != nil {
		panic(err)
	}

	return string(plaintext)
}
