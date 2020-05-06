package models

import (
	"github.com/jinzhu/gorm"
)

// Role Inject fields `ID`, `CreatedAt`, `UpdatedAt`, `DeletedAt` into model`
type Role struct {
	gorm.Model
	Name        string `gorm:"unique;not null" json:"name" xml:"name" form:"name" query:"name"`
	Description string `json:"description" gorm:"type:varchar(100);" `
}
