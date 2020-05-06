package database

import (
	"fiberapi/models"
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

var db *gorm.DB
var err error

// Connect to db
func Connect() {
	isLoad := godotenv.Load()
	if isLoad != nil {
		log.Fatal("Error loading .env file")
	}

	dbUser := os.Getenv("db_user")
	dbPass := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")
	db, _ = gorm.Open("mysql", dbUser+":"+dbPass+"@tcp("+dbHost+")/"+dbName+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
	}
	db.AutoMigrate(&models.Role{})
	db.AutoMigrate(&models.User{})
	db.Model(&models.User{}).AddForeignKey("role_id", "roles(id)", "RESTRICT", "CASCADE")

	db.AutoMigrate(&models.Test{})

}

//Instance ...
func Instance() *gorm.DB {
	return db
}

// Close conn
func Close() {
	db.Close()
}
