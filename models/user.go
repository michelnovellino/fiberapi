package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// User model
type User struct {
	gorm.Model
	FirstName string `json:"first_name" xml:"first_name" form:"first_name" query:"first_name"`
	LastName  string
	birthDay  *time.Time
	Email     string
	RoleID    uint `gorm:"column:role_id" json:"role_id"`
	Role      Role
}
