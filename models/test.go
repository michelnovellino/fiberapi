package models

import (
	"github.com/jinzhu/gorm"
)

// Test Inject fields `ID`, `CreatedAt`, `UpdatedAt`, `DeletedAt` into model`
type Test struct {
	gorm.Model
	Name string
}
