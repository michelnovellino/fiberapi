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
	println("params", c.Params("*"))
	test := models.Test{Name: "testone"}
	if res := db.Create(&test); res.Error != nil {
		println("have error ", res.Error)
	}
	db.Close()
	c.Send("success!")
}

//Get handler
func Get(c *fiber.Ctx) {

	res := c.Params("id")
	c.Send(res)
}

// Post handler
func Post(c *fiber.Ctx) {
	// db := database.Instance()
	// db.LogMode(true)
	test := new(models.Test)
	if err := c.BodyParser(test); err != nil {
		log.Fatal(err)
	}

	log.Println(test.Name)

	c.Send(test.Name)
}

// Put handler
func Put(c *fiber.Ctx) {
	res := c.Params("id")
	test := new(models.Test)
	if err := c.BodyParser(test); err != nil {
		log.Fatal(err)
	}

	log.Println(test.Name)
	c.Send(res)
}

// Delete handler
func Delete(c *fiber.Ctx) {
	res := c.Params("id")
	c.Send(res)
}
