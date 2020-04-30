package controllers

import (
	"fiberapi/database"

	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
)

// User ...
type User struct {
	gorm.Model
	Name string
}

// GetAll controller
func GetAll(c *fiber.Ctx) {
	db := database.Instance()
	if res := db.Create(&User{Name: "testname"}); res.Error != nil {
		println("have error ", res.Error)
	}
	db.Close()
	c.Send("success!")
}
