package models

import "github.com/jinzhu/gorm"

// User model
type User struct {
	gorm.Model
	FirstName string
	LastName  string
}
