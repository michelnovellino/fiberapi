package main

import (
	"fiberapi/routes"

	"github.com/gofiber/cors"
	"github.com/gofiber/fiber"
	"github.com/gofiber/logger"
)

func main() {
	app := fiber.New()

	app.Settings.Prefork = true
	app.Settings.CaseSensitive = true
	app.Settings.StrictRouting = true
	app.Settings.ServerHeader = "Fiberabid"

	app.Static("/static", "./public", fiber.Static{
		Compress:  true,
		ByteRange: true,
		Browse:    true,
		Index:     "lol.txt",
	})
	app.Use(cors.New())
	app.Use(logger.New())

	app.Get("/greetings", func(c *fiber.Ctx) {
		c.Send("Welcome!")
	})
	v1 := app.Group("/v1")
	routes.Register(app, v1)

	app.Use(func(c *fiber.Ctx) {
		c.SendStatus(404) // => 404 "Not Found"
		c.Send("This is a dummy route")
	})
	app.Listen(3000)
}
