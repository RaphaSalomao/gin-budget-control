package service

import (
	"errors"
	"fmt"

	"github.com/RaphaSalomao/gin-budget-control/database"
	"github.com/RaphaSalomao/gin-budget-control/model"
	"github.com/RaphaSalomao/gin-budget-control/model/entity"
	"github.com/RaphaSalomao/gin-budget-control/security"
	"github.com/RaphaSalomao/gin-budget-control/utils"
	"gorm.io/gorm"
)

type userService struct{}

var UserService = userService{}

func (us *userService) CreateUser(u *entity.UserRequest) error {
	hash := utils.HashPassword(u.Password)
	user := entity.User{
		Email:    u.Email,
		Password: hash,
	}
	tx := database.DB.Create(&user)
	return tx.Error
}

func (us *userService) Authenticate(u *entity.UserRequest) (model.TokenResponse, error) {
	user := entity.User{}
	tx := database.DB.Where("email = ?", u.Email).First(&user)
	if tx.Error != nil {
		if tx.Error == gorm.ErrRecordNotFound {
			return model.TokenResponse{}, fmt.Errorf("user %s not found", u.Email)
		} else {
			return model.TokenResponse{}, tx.Error
		}
	}
	if utils.ValidadePasswordHash(u.Password, user.Password) {
		token, err := security.GenerateJWT(u.Email)
		return model.TokenResponse{
			Token: token,
		}, err
	} else {
		return model.TokenResponse{}, errors.New("invalid user/password")
	}
}
