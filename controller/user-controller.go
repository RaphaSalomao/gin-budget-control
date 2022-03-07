package controller

import (
	"fmt"
	"net/http"

	"github.com/RaphaSalomao/gin-budget-control/model"
	"github.com/RaphaSalomao/gin-budget-control/model/entity"
	"github.com/RaphaSalomao/gin-budget-control/service"
	"github.com/gin-gonic/gin"
)

// Create User
// @Description create a new user
// @Tags User
// @Param user body entity.UserRequest true "User"
// @Success 201 {object} entity.UserRequest
// @Router /budget-control/api/v1/user [post]
func CreateUser(c *gin.Context) {
	var user entity.UserRequest
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, model.ValidationErrorResponse(err))
		fmt.Println(err)
		return
	}
	if err := service.UserService.CreateUser(&user); err != nil {
		c.JSON(http.StatusUnprocessableEntity, model.ErrorResponse{
			Error:   err.Error(),
			Message: "Unable to create user",
		})
		return
	}
	c.JSON(http.StatusCreated, user)
}

// Authenticate
// @Description authenticate user
// @Tags User
// @Param user body entity.UserRequest true "User"
// @Success 201 {string} string "token"
// @Router /budget-control/api/v1/user/authenticate [post]
func Authenticate(c *gin.Context) {
	var user entity.UserRequest
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, model.ValidationErrorResponse(err))
		return
	}
	token, err := service.UserService.Authenticate(&user)
	if err != nil {
		c.JSON(http.StatusUnauthorized, model.ErrorResponse{
			Error:   err.Error(),
			Message: "Unauthorized",
		})
		return
	}
	c.JSON(http.StatusOK, token)
}
