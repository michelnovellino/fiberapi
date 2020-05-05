package routes

import (
	"fiberapi/controllers"

	"github.com/gofiber/fiber"
)

//TestRoutes register
func TestRoutes(app *fiber.App) {
	namespace := app.Group("/test")
	namespace.Get("/", controllers.GetAllTest)
	namespace.Get("/:id", controllers.GetTest)
	namespace.Post("/", controllers.AddTest)
	namespace.Put("/:id", controllers.EditTest)
	namespace.Delete("/:id", controllers.RemoveTest)
	namespace.Get("/admin/*", func(c *fiber.Ctx) {
		c.Send("API path: " + c.Params("*"))
	})
}
