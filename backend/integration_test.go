package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http/httptest"
	"testing"

	"github.com/go-faker/faker/v4"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/viniciusataide/velozient-challenge-go/domain"
	"github.com/viniciusataide/velozient-challenge-go/passwordcards"
	"github.com/viniciusataide/velozient-challenge-go/utils"
)

var app *fiber.App
var ctrl *passwordcards.Controller

func setup(t *testing.T) func() {
	err := godotenv.Load()

	if err != nil {
		panic(err)
	}

	app = fiber.New()
	ctrl = bootstrapTest(app)

	return func() {
	}
}

func TestList(t *testing.T) {
	defer setup(t)()

	// Act
	req := httptest.NewRequest("GET", "/api/v1/password-cards", nil)
	resp, err := app.Test(req, 1)

	body, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()

	var cards []domain.PasswordCard

	err = json.Unmarshal(body, &cards)

	if err != nil {
		panic(err)
	}

	// Assert
	assert.Equalf(t, fiber.StatusOK, resp.StatusCode, "Should respond 200 when listing")
	assert.Equalf(t, len(cards), utils.CARDS_LENGTH, "Should list Password Cards")
}

func TestCreate(t *testing.T) {
	defer setup(t)()

	newCard := domain.PasswordCardCreate{
		URL:      faker.URL(),
		Name:     faker.Name(),
		Username: faker.Username(),
		Password: faker.Password(),
	}

	payload, _ := json.Marshal(newCard)

	req := httptest.NewRequest("POST", "/api/v1/password-cards", bytes.NewReader(payload))

	req.Header.Set("Content-Type", "application/json")

	res, _ := app.Test(req, 1)

	body, _ := io.ReadAll(res.Body)
	defer res.Body.Close()

	removeQuotes := string(body)[1 : len(string(body))-1]
	lastCreatedId := (*ctrl.Service.Repository.PasswordCards)[utils.CARDS_LENGTH].Id.String()

	assert.Equalf(t, fiber.StatusCreated, res.StatusCode, "Should 200 when create a new Password Card")
	assert.Equalf(t, lastCreatedId, removeQuotes, "Id's should be the same")
	assert.Equalf(t, utils.CARDS_LENGTH+1, len(*ctrl.Service.Repository.PasswordCards), "Should create a new Password Card")

}

func TestGet(t *testing.T) {
	defer setup(t)()
	firstId := (*ctrl.Service.Repository.PasswordCards)[0].Id.String()

	// Act
	req := httptest.NewRequest("GET", fmt.Sprintf("/api/v1/password-cards/%s", firstId), nil)
	resp, err := app.Test(req, 1)

	body, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()

	var card domain.PasswordCard

	err = json.Unmarshal(body, &card)

	if err != nil {
		panic(err)
	}

	// Assert
	assert.Equalf(t, card.Id.String(), firstId, "Should get a Password Card")

}

func TestDelete(t *testing.T) {
	defer setup(t)()
	passwordCards := ctrl.Service.Repository.PasswordCards
	firstId := (*passwordCards)[0].Id.String()

	// Act
	req := httptest.NewRequest("DELETE", fmt.Sprintf("/api/v1/password-cards/%s", firstId), nil)
	app.Test(req, 1)

	// Assert
	assert.Equalf(t, 2, len(*passwordCards), "Should get a Password Card")

}

func TestUpdate(t *testing.T) {
	defer setup(t)()
	firstId := (*ctrl.Service.Repository.PasswordCards)[0].Id.String()

	cardToUpdate := domain.PasswordCardCreate{
		URL:      faker.URL(),
		Name:     faker.Name(),
		Username: faker.Username(),
		Password: faker.Password(),
	}

	payload, _ := json.Marshal(cardToUpdate)

	// Act
	req := httptest.NewRequest("PATCH", fmt.Sprintf("/api/v1/password-cards/%s", firstId), bytes.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req, 1)

	defer resp.Body.Close()

	if err != nil {
		panic(err)
	}

	// Assert
	assert.Equalf(t, fiber.StatusOK, resp.StatusCode, "Should send a status OK")
}
