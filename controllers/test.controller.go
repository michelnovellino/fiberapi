package controllers

import "github.com/gofiber/fiber"

// GetAll controller
func GetAll(c *fiber.Ctx) {
	c.Send("hello")
}
