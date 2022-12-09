package models

import (
	"gorm.io/gorm"
)

type Trash struct {
	gorm.Model
	Type     string `validate:"nonzero" json:"type"`
	Name     string `validate:"nonzero" json:"name"`
	ImageUrl string
	Active   bool `gorm:"default:true"`
}

var Trashes []Trash
