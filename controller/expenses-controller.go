package controller

import (
	"net/http"

	"github.com/RaphaSalomao/gin-budget-control/model"
	"github.com/RaphaSalomao/gin-budget-control/model/entity"
	"github.com/RaphaSalomao/gin-budget-control/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Create Expense
// @Summary Create a new expense
// @Description Create a new expense. Obs.: you cannot create two expenses with the same description in a single month.
// @Tags Expenses
// @Param expense body entity.ExpenseRequest true "Expense"
// @Success 201 {object} uuid.UUID
// @Router /budget-control/api/v1/expense [post]
func CreateExpense(c *gin.Context) {
	userId := uuid.MustParse(c.GetString("userId"))
	var expense entity.ExpenseRequest
	if err := c.ShouldBindJSON(&expense); err != nil {
		c.JSON(http.StatusBadRequest, model.ValidationErrorResponse(err))
		return
	}
	id, err := service.ExpenseService.CreateExpense(&expense, userId)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, model.ErrorResponse{
			Error:   err.Error(),
			Message: "Error creating expense",
			Id:      id,
		})
		return
	}
	c.JSON(http.StatusCreated, struct{ Id uuid.UUID }{id})
}

// Find All Expenses
// @Summary Find all expenses
// @Description Find all expenses
// @Tags Expenses
// @Success 200 {array} entity.ExpenseResponse
// @Router /budget-control/api/v1/expense [get]
func FindAllExpenses(c *gin.Context) {
	userId := uuid.MustParse(c.GetString("userId"))
	expenses := []entity.ExpenseResponse{}
	description := c.Query("description")
	service.ExpenseService.FindAllExpenses(&expenses, description, userId)
	c.JSON(http.StatusOK, expenses)
}

// Find Expense By Id
// @Summary Find expense by id
// @Description Find expense by id
// @Tags Expenses
// @Param id path string true "Expense ID"
// @Success 200 {object} entity.ExpenseResponse
// @Router /budget-control/api/v1/expense/{id} [get]
func FindExpense(c *gin.Context) {
	userId := uuid.MustParse(c.GetString("userId"))
	var expense entity.ExpenseResponse
	id := uuid.MustParse(c.Param("id"))
	err := service.ExpenseService.FindExpense(&expense, id, userId)
	if err != nil {
		c.JSON(http.StatusNotFound, model.ErrorResponse{
			Error:   err.Error(),
			Message: "Expense not found",
		})
		return
	}
	c.JSON(http.StatusOK, expense)
}

// Update Expense
// @Summary Update an expense
// @Description Update an expense
// @Tags Expenses
// @Param id path string true "Expense ID"
// @Param expense body entity.ExpenseRequest true "Expense"
// @Success 204
// @Router /budget-control/api/v1/expense/{id} [put]
func UpdateExpense(c *gin.Context) {
	userId := uuid.MustParse(c.GetString("userId"))
	var expense entity.ExpenseRequest
	id := uuid.MustParse(c.Param("id"))
	if err := c.ShouldBindJSON(&expense); err != nil {
		c.JSON(http.StatusBadRequest, model.ValidationErrorResponse(err))
		return
	}
	id, err := service.ExpenseService.UpdateExpense(&expense, id, userId)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, model.ErrorResponse{
			Error:   err.Error(),
			Message: "Error updating expense",
			Id:      id,
		})
		return
	}
	c.JSON(http.StatusNoContent, struct{ Id uuid.UUID }{id})
}

// Delete Expense
// @Summary Delete an expense
// @Description Delete an expense
// @Tags Expenses
// @Param id path string true "Expense ID"
// @Success 204
// @Router /budget-control/api/v1/expense/{id} [delete]
func DeleteExpense(c *gin.Context) {
	userId := uuid.MustParse(c.GetString("userId"))
	id := uuid.MustParse(c.Param("id"))
	service.ExpenseService.DeleteExpense(id, userId)
	c.JSON(http.StatusNoContent, nil)
}

// Find All Expenses By Period
// @Summary Find all expenses by period
// @Description Find all expenses by period
// @Tags Expenses
// @Param year path int true "Year"
// @Param month path int true "Month"
// @Success 200 {array} entity.ExpenseResponse
// @Router /budget-control/api/v1/expense/period/{year}/{month} [get]
func ExpensesByPeriod(c *gin.Context) {
	userId := uuid.MustParse(c.GetString("userId"))
	var expenses []entity.ExpenseResponse
	err := service.ExpenseService.ExpensesByPeriod(&expenses, c.Param("year"), c.Param("month"), userId)
	if err != nil {
		c.JSON(http.StatusNotFound, model.ErrorResponse{
			Error:   err.Error(),
			Message: "Expenses not found",
		})
		return
	}
	c.JSON(http.StatusOK, expenses)
}
