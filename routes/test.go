package routes

import (
	"fiberapi/controllers"

	"github.com/gofiber/fiber"
)

//TestRoutes register
func TestRoutes(app *fiber.App) {
	namespace := app.Group("/users")
	namespace.Get("/get", controllers.GetAll)
	namespace.Post("/", controllers.Post)
	namespace.Get("/admin/*", func(c *fiber.Ctx) {
		c.Send("API path: " + c.Params("*"))
	})
}
