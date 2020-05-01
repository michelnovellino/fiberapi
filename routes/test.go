package routes

import (
	"fiberapi/controllers"

	"github.com/gofiber/fiber"
)

//TestRoutes register
func TestRoutes(app *fiber.App) {
	namespace := app.Group("/users")
	namespace.Get("/", controllers.GetAll)
	namespace.Get("/:id", controllers.Get)
	namespace.Post("/", controllers.Post)
	namespace.Put("/:id", controllers.Put)
	namespace.Delete("/:id", controllers.Delete)
	namespace.Get("/admin/*", func(c *fiber.Ctx) {
		c.Send("API path: " + c.Params("*"))
	})
}
