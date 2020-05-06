package controllers

import (
	"fiberapi/database"
	"fiberapi/models"
	"log"

	"github.com/gofiber/fiber"
)

// GetAllUsers handler
func GetAllUsers(c *fiber.Ctx) {
	db := database.Instance()
	db.LogMode(true)
	var User []models.User

	if res := db.Find(&User); res.Error != nil {
		log.Fatal("have error ", res.Error)
		c.Send("not records found")
	}
	c.JSON(User)
	db.Close()
}

//GetUser handler
func GetUser(c *fiber.Ctx) {
	db := database.Instance()
	db.LogMode(true)
	User := new(models.User)
	id := c.Params("id")
	if res := db.Find(&User, id); res.Error != nil {
		println("have error ", res.Error)
	}
	c.JSON(User)
}

// AddUser handler
func AddUser(c *fiber.Ctx) {
	db := database.Instance()
	db.LogMode(true)
	User := new(models.User)
	if err := c.BodyParser(User); err != nil {
		log.Fatal(err)
	}

	if res := db.Create(&User); res.Error != nil {
		log.Fatal("have error ", res.Error)
	}

	log.Println(User.FirstName)
	db.Close()
	c.Send(User.FirstName)
}

// EditUser handler
func EditUser(c *fiber.Ctx) {
	id := c.Params("id")
	User := new(models.User)

	if err := c.BodyParser(User); err != nil {
		log.Fatal(err)
	}
	log.Println(User.FirstName)
	newName := User.FirstName
	db := database.Instance()

	db.First(&User, id) // find User with id
	User.FirstName = newName

	db.Save(&User)
	db.Close()
	log.Println(User)
	c.Send("updated User with id: " + id)
}

// RemoveUser handler
func RemoveUser(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.Instance()

	var User models.User
	db.Find(&User)
	if res := db.Find(&User); res.Error != nil {
		log.Fatal("have error ", res.Error)
	}
	db.Delete(&User)

	c.Send("deleted User with id: ", id)
}
