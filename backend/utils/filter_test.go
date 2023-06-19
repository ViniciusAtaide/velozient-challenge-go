package utils

import (
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/viniciusataide/velozient-challenge-go/server"
)

func TestFilter(t *testing.T) {
	godotenv.Load("../.env")
	cards := ProvideTestCards(ProvideCrypto(server.ProvideConfig()))

	firstCardName := (*cards)[0].Name
	result := Filter(*cards, firstCardName)

	assert.Equalf(t, 1, len(result), "Should only get one result")
	assert.Equalf(t, firstCardName, result[0].Name, "This result should have the same name")
}
