package routes

import (
	"github.com/gofiber/fiber"
)

//Register routes
func Register(app *fiber.App) {
	v1 := app.Group("/users")
	v1.Get("/get/*", func(c *fiber.Ctx) {
		c.Send("API path: " + c.Params("*"))
	})
	v1.Get("/admin", func(c *fiber.Ctx) {
		c.Send("API path: " + c.Params("*"))
	})
}
