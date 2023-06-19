//go:build wireinject
// +build wireinject

package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/wire"
	"github.com/viniciusataide/velozient-challenge-go/domain"
	"github.com/viniciusataide/velozient-challenge-go/passwordcards"
	"github.com/viniciusataide/velozient-challenge-go/server"
	"github.com/viniciusataide/velozient-challenge-go/utils"
)

func bootstrap(app *fiber.App) *passwordcards.Controller {
	wire.Build(
		passwordcards.PasswordCardsModule,
		server.ProvideConfig,
		server.ProvideServer,
		utils.ProvideCrypto,
		domain.ProvideCards,
	)

	return &passwordcards.Controller{}
}
func bootstrapTest(app *fiber.App) *passwordcards.Controller {
	wire.Build(
		passwordcards.PasswordCardsModule,
		utils.ProvideTestCards,
		server.ProvideConfig,
		server.ProvideServer,
		utils.ProvideCrypto,
	)

	return &passwordcards.Controller{}
}
