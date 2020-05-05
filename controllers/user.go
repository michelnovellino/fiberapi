package controllers

import (
	"github.com/gofiber/fiber"
)

// GetAllUsers handler
func GetAllUsers(c *fiber.Ctx) {
	c.Send("user route get all")
}
