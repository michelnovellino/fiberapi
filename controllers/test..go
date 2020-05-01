package controllers

import (
	"fiberapi/database"
	"fiberapi/models"
	"log"

	"github.com/gofiber/fiber"
)

// GetAll handler
func GetAll(c *fiber.Ctx) {
	db := database.Instance()
	db.LogMode(true)
	var test []models.Test

	if res := db.Find(&test); res.Error != nil {
		log.Fatal("have error ", res.Error)
		c.Send("not records found")
	}
	c.JSON(test)
	db.Close()
}

//Get handler
func Get(c *fiber.Ctx) {
	db := database.Instance()
	db.LogMode(true)
	test := new(models.Test)
	id := c.Params("id")
	if res := db.Find(&test, id); res.Error != nil {
		println("have error ", res.Error)
	}
	c.JSON(test)
}

// Post handler
func Post(c *fiber.Ctx) {
	db := database.Instance()
	db.LogMode(true)
	test := new(models.Test)
	if err := c.BodyParser(test); err != nil {
		log.Fatal(err)
	}
	if res := db.Create(&test); res.Error != nil {
		log.Fatal("have error ", res.Error)
	}

	log.Println(test.Name)
	db.Close()
	c.Send(test.Name)
}

// Put handler
func Put(c *fiber.Ctx) {
	id := c.Params("id")
	test := new(models.Test)

	if err := c.BodyParser(test); err != nil {
		log.Fatal(err)
	}
	log.Println(test.Name)
	newName := test.Name
	db := database.Instance()

	db.First(&test, id) // find test with id
	test.Name = newName

	db.Save(&test)
	db.Close()
	log.Println(test)
	c.Send("updated test with id: " + id)
}

// Delete handler
func Delete(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.Instance()

	var test models.Test
	db.Find(&test)
	if res := db.Find(&test); res.Error != nil {
		log.Fatal("have error ", res.Error)
	}
	db.Delete(&test)

	c.Send("deleted test with id: ", id)
}
