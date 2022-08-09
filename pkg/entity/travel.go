package entity

import (
	"github.com/jinzhu/gorm"
)

// Flag is the unit of flags
type Travel struct {
	gorm.Model

	NameTravel  string `sql:"type:text"`
	Country     string
	Description string `sql:"type:text"`
}
