package model

import (
	"strings"

	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type User struct {
	Base
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserRequest struct {
	Email    string `json:"email" validate:"nonzero, regexp=^[aA-zZ\\.\\-\\_\\d]+@[aA-zZ\\d]+\\.[aA-zZ]+$"`
	Password string `json:"password" validate:"nonzero"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.Base.BeforeCreate(tx)
	u.Email = strings.ToLower(u.Email)
	return nil
}

func (u *UserRequest) Validate() error {
	if err := validator.Validate(u); err != nil {
		return err
	}
	return nil
}
