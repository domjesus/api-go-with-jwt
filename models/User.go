package models

import (
	"errors"
	"net/mail"
	"reflect"

	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name             string `validate:"nonzero" creating: "nonzero" json:"name"`
	Email            string `validate:"nonzero" creating: "nonzero" gorm:"unique" json:"email"`
	Password         string `creating:"nonzero" json:"password"`
	Password_confirm string `gorm:"-" gorm:"-:migration" json:"password_confirm"`
	Role             string `json:"role"`
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

func ValidaPasswords(v interface{}, passwordconfirm string) error {
	st := reflect.ValueOf(v)

	if st.String() != passwordconfirm {
		return errors.New("Passwords not matches!")
	}

	return nil

}
