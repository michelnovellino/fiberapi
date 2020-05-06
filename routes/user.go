package routes

import (
	"fiberapi/controllers"

	"github.com/gofiber/fiber"
)

//UserRoutes register
func UserRoutes(app *fiber.App) {
	namespace := app.Group("/users")
	namespace.Get("/", controllers.GetAllUsers)
	namespace.Get("/:id", controllers.GetUser)
	namespace.Post("/", controllers.AddUser)
	namespace.Put("/:id", controllers.EditUser)
	namespace.Delete("/:id", controllers.RemoveUser)
	namespace.Get("/admin/*", func(c *fiber.Ctx) {
		c.Send("API path: " + c.Params("*"))
	})
}
