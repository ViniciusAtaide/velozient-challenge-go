package utils

import (
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/viniciusataide/velozient-challenge-go/server"
)

func TestCrypt(t *testing.T) {
	err := godotenv.Load("../.env")

	if err != nil {
		panic(err)
	}

	config := server.ProvideConfig()
	crypto := ProvideCrypto(config)
	password := "password"

	assert.NotEqualf(t, password, crypto.EncryptAES(password), "should not fail encryption")
}

func TestDecrypt(t *testing.T) {
	err := godotenv.Load("../.env")

	if err != nil {
		panic(err)
	}

	config := server.ProvideConfig()
	crypto := ProvideCrypto(config)
	password := "password"

	encrypted := crypto.EncryptAES(password)

	decripted := crypto.DecryptAES(encrypted)

	assert.Equalf(t, password, decripted, "should not fail encryption")
}
