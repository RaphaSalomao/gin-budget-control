package controller

import (
	"net/http"

	"github.com/RaphaSalomao/gin-budget-control/model"
	"github.com/RaphaSalomao/gin-budget-control/service"
	"github.com/gin-gonic/gin"
)

// Create User
// @Description create a new user
// @Tags User
// @Param user body model.UserRequest true "User"
// @Success 201 {object} model.UserRequest
// @Router /budget-control/api/v1/user [post]
func CreateUser(c *gin.Context) {
	user, _ := c.Get("body")
	if err := service.UserService.CreateUser(user.(*model.UserRequest)); err != nil {
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
// @Param user body model.UserRequest true "User"
// @Success 201 {string} string "token"
// @Router /budget-control/api/v1/user/authenticate [post]
func Authenticate(c *gin.Context) {
	user, _ := c.Get("body")
	token, err := service.UserService.Authenticate(user.(*model.UserRequest))
	if err != nil {
		c.JSON(http.StatusUnauthorized, model.ErrorResponse{
			Error:   err.Error(),
			Message: "Unauthorized",
		})
		return
	}
	c.JSON(http.StatusOK, token)
}
