package models

import (
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title       string `validate:"nonzero" json:"title"`
	Author      Author `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	AuthorID    int
	YearPublish string `json:"year_publish"`
	ISBN        string `json:"isbn"`
}

var Books []Book

func ValidaDadosDaRequest(book *Book) error {
	if err := validator.Validate(book); err != nil {
		return err
	}

	return nil
}
