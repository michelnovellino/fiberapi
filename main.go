package main

import (
	"fiberapi/database"
	"fiberapi/routes"

	"github.com/gofiber/cors"
	"github.com/gofiber/fiber"
	"github.com/gofiber/logger"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	app := fiber.New()

	app.Settings.Prefork = true
	app.Settings.CaseSensitive = true
	app.Settings.StrictRouting = true
	app.Settings.ServerHeader = "Fiberabid"
	database.Connect()
	app.Static("/static", "./public", fiber.Static{
		Compress:  true,
		ByteRange: true,
		Browse:    true,
		Index:     "lol.txt",
	})
	app.Use(cors.New())
	app.Use(logger.New())

	app.Get("/", func(c *fiber.Ctx) {

		c.Send("Hola, Mundo!")
	})

	routes.TestRoutes(app)
	routes.RoleRoutes(app)
	routes.UserRoutes(app)

	app.Use(func(c *fiber.Ctx) {
		c.SendStatus(404) // => 404 "Not Found"
		c.Send("Not Found")
	})
	app.Listen(3000)
}
