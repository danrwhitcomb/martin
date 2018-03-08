package models

import (
	"github.com/jinzhu/gorm"
)

type Schema struct {
	gorm.Model

	Version int `gorm:"Column:version"`
}
