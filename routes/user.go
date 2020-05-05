package routes

import (
	"fiberapi/controllers"

	"github.com/gofiber/fiber"
)

//UserRoutes ...
func UserRoutes(app *fiber.App) {
	namespace := app.Group("/users")
	namespace.Get("/", controllers.GetAllUsers)
}
