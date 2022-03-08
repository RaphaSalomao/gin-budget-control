package entity

import (
	"strings"

	"gorm.io/gorm"
)

type User struct {
	Base
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.Base.BeforeCreate(tx)
	u.Email = strings.ToLower(u.Email)
	return nil
}
