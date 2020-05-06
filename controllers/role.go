package controllers

import (
	"fiberapi/database"
	"fiberapi/models"
	"log"

	"github.com/gofiber/fiber"
)

// GetAllRoles handler
func GetAllRoles(c *fiber.Ctx) {
	db := database.Instance()
	db.LogMode(true)
	var Role []models.Role

	if res := db.Find(&Role); res.Error != nil {
		log.Fatal("have error ", res.Error)
		c.Send("not records found")
	}
	c.JSON(Role)
	db.Close()
}

//GetRole handler
func GetRole(c *fiber.Ctx) {
	db := database.Instance()
	db.LogMode(true)
	Role := new(models.Role)
	id := c.Params("id")
	if res := db.Find(&Role, id); res.Error != nil {
		println("have error ", res.Error)
	}
	c.JSON(Role)
}

// AddRole handler
func AddRole(c *fiber.Ctx) {
	db := database.Instance()
	db.LogMode(true)
	Role := new(models.Role)
	if err := c.BodyParser(Role); err != nil {
		log.Fatal(err)
	}
	if res := db.Create(&Role); res.Error != nil {
		log.Fatal("have error ", res.Error)
	}

	log.Println(Role.Name)
	db.Close()
	c.Send(Role.Name)
}

// EditRole handler
func EditRole(c *fiber.Ctx) {
	id := c.Params("id")
	Role := new(models.Role)

	if err := c.BodyParser(Role); err != nil {
		log.Fatal(err)
	}
	log.Println(Role.Name)
	newName := Role.Name
	db := database.Instance()

	db.First(&Role, id) // find Role with id
	Role.Name = newName

	db.Save(&Role)
	db.Close()
	log.Println(Role)
	c.Send("updated Role with id: " + id)
}

// RemoveRole handler
func RemoveRole(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.Instance()

	var Role models.Role
	db.Find(&Role)
	if res := db.Find(&Role); res.Error != nil {
		log.Fatal("have error ", res.Error)
	}
	db.Delete(&Role)

	c.Send("deleted Role with id: ", id)
}
