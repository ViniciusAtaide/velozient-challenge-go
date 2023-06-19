package main

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	_ "github.com/viniciusataide/velozient-challenge-go/docs"
)

// @title PasswordCard's API
// @version 1.0
// @description This is an api for Passwordcards
// @contact.name Vinicius Ataide
// @contact.email viniciusataid@gmail.com
// @host localhost:3000
// @BasePath /api/v1
func main() {
	err := godotenv.Load()
	app := fiber.New()

	if err != nil {
		panic(err)
	}

	bootstrap(app)

	app.Listen(fmt.Sprintf(":%s", os.Getenv("SERVER_PORT")))
}
