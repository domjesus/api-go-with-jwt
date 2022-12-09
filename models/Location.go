package models

import (
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type Location struct {
	gorm.Model
	Latitude  string `validate:"nonzero" json:"latitude"`
	Longitude string `validate:"nonzero" json:"longitude"`
	UserId    int
	User      User
	TrashId   int
	Trash     Trash
	Valid     bool
	Active    bool
}

var Locations []Location

func ValidaDadosLocation(location *Location) error {
	if err := validator.Validate(location); err != nil {
		return err
	}

	return nil
}
