package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"
)

type Server struct {
	fiber.Router
}

func ProvideServer(app *fiber.App, config *Config) *Server {
	app.Use(cors.New(cors.Config{
		AllowOrigins:     config.BackendCors,
		AllowMethods:     "*",
		AllowHeaders:     "*",
		AllowCredentials: true,
	}))

	app.Get("/swagger/*", swagger.HandlerDefault)
	app.Use(recover.New())
	app.Use(logger.New())

	group := app.Group("/api")
	v1 := group.Group("/v1")

	return &Server{v1}
}
