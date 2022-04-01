package service

import (
	"github.com/RaphaSalomao/gin-budget-control/database"
	"github.com/RaphaSalomao/gin-budget-control/model"
	"github.com/RaphaSalomao/gin-budget-control/security"
	"github.com/RaphaSalomao/gin-budget-control/utils"
)

type userService struct{}

var UserService = userService{}

func (us *userService) CreateUser(u *model.UserRequest) error {
	hash := utils.HashPassword(u.Password)
	user := model.User{
		Email:    u.Email,
		Password: hash,
	}
	tx := database.DB.Create(&user)
	return tx.Error
}

func (us *userService) Authenticate(u *model.UserRequest) (model.TokenResponse, error) {
	user := model.User{}
	if tx := database.DB.Where("email = ?", u.Email).First(&user); tx.Error != nil {
		return model.TokenResponse{}, model.ErrNotFound
	}
	if utils.IsHashAndPasswordMatched(u.Password, user.Password) {
		token, err := security.GenerateJWT(u.Email)
		return model.TokenResponse{
			Token: token,
		}, err
	}
	return model.TokenResponse{}, model.ErrInvalidUserPassword
}
