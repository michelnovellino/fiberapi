package routes

import (
	"fiberapi/controllers"

	"github.com/gofiber/fiber"
)

//RoleRoutes register
func RoleRoutes(app *fiber.App) {
	namespace := app.Group("/roles")
	namespace.Get("/", controllers.GetAllRoles)
	namespace.Get("/:id", controllers.GetRole)
	namespace.Post("/", controllers.AddRole)
	namespace.Put("/:id", controllers.EditRole)
	namespace.Delete("/:id", controllers.RemoveRole)
	namespace.Get("/admin/*", func(c *fiber.Ctx) {
		c.Send("API path: " + c.Params("*"))
	})
}
