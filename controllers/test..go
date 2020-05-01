package controllers

import (
	"fiberapi/database"
	"fiberapi/models"

	"github.com/gofiber/fiber"
)

// GetAll controller
func GetAll(c *fiber.Ctx) {
	db := database.Instance()
	test := models.Test{Name: "testone"}
	if res := db.Create(&test); res.Error != nil {
		println("have error ", res.Error)
	}
	db.Close()
	c.Send("success!")
}
