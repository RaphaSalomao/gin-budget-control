package controller

import (
	"net/http"

	"github.com/RaphaSalomao/gin-budget-control/model"
	"github.com/RaphaSalomao/gin-budget-control/model/entity"
	"github.com/RaphaSalomao/gin-budget-control/service"
	"github.com/gin-gonic/gin"
)

// Month Balance Sumary
// @Description get month balance sumary
// @Tags Sumary
// @Param year path string true "Year"
// @Param month path string true "Month"
// @Success 200 {object} model.BalanceSumaryResponse
// @Router /budget-control/api/v1/summary/{year}/{month} [get]
func MonthBalanceSumary(c *gin.Context) {
	user, _ := c.Get("user")
	var balanceSumary model.BalanceSumaryResponse
	err := service.BalanceSumary(&balanceSumary, c.Param("year"), c.Param("month"), user.(*entity.User).Id)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, model.ErrorResponse{
			Error:   err.Error(),
			Message: "Error getting balance sumary",
		})
		return
	}
	c.JSON(http.StatusOK, balanceSumary)
}
