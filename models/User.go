package models

import (
	"net/mail"

	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `validate: "nonzero" json:"name"`
	Email    string `validate: "nonzero" gorm:"unique" json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

var Users []User

func ValidaUser(user *User) error {

	if err := validator.Validate(user); err != nil {
		return err
	}

	return nil
}

func ValidaEmail(email string) error {
	if _, err := mail.ParseAddress(email); err != nil {
		return err
	}
	return nil
}
