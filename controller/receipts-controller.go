package controller

import (
	"net/http"

	"github.com/RaphaSalomao/gin-budget-control/model"
	"github.com/RaphaSalomao/gin-budget-control/model/entity"
	"github.com/RaphaSalomao/gin-budget-control/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Create Receipt
// @Summary Create a new receipt
// @Description Create a new receipt. Obs.: you cannot create two receipts with the same description in a single month.
// @Tags Receipts
// @Param receipt body entity.ReceiptRequest true "Receipt"
// @Success 201 {object} uuid.UUID
// @Router /budget-control/api/v1/receipt [post]
func CreateReceipt(c *gin.Context) {
	user, _ := c.Get("user")
	receipt, _ := c.Get("body")
	id, err := service.ReceiptService.CreateReceipt(receipt.(*entity.ReceiptRequest), user.(*entity.User).Id)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, model.ErrorResponse{
			Error:   err.Error(),
			Message: "Error creating receipt",
			Id:      id.String(),
		})
		return
	}
	c.JSON(http.StatusCreated, struct{ Id uuid.UUID }{id})
}

// Find All Receipts
// @Summary Find all receipts
// @Description Find all receipts
// @Tags Receipts
// @Success 200 {array} entity.ReceiptResponse
// @Router /budget-control/api/v1/receipt [get]
func FindAllReceipts(c *gin.Context) {
	user, _ := c.Get("user")
	receipts := []entity.ReceiptResponse{}
	description := c.Query("description")
	service.ReceiptService.FindAllReceipts(&receipts, description, user.(*entity.User).Id)
	c.JSON(http.StatusOK, receipts)
}

// Find Receipt By Id
// @Summary Find a receipt by id
// @Description Find a receipt by id
// @Tags Receipts
// @Param id path string true "Receipt id"
// @Success 200 {object} entity.ReceiptResponse
// @Router /budget-control/api/v1/receipt/{id} [get]
func FindReceipt(c *gin.Context) {
	user, _ := c.Get("user")
	var receipt entity.ReceiptResponse
	id := uuid.MustParse(c.Param("id"))
	err := service.ReceiptService.FindReceipt(&receipt, id, user.(*entity.User).Id)
	if err != nil {
		c.JSON(http.StatusNotFound, model.ErrorResponse{
			Error:   err.Error(),
			Message: "Receipt not found",
			Id:      id.String(),
		})
		return
	}
	c.JSON(http.StatusOK, receipt)
}

// Update Receipt
// @Summary Update a receipt
// @Description Update a receipt
// @Tags Receipts
// @Param id path string true "Receipt id"
// @Param receipt body entity.ReceiptRequest true "Receipt"
// @Success 200 {object} uuid.UUID
// @Router /budget-control/api/v1/receipt/{id} [put]
func UpdateReceipt(c *gin.Context) {
	user, _ := c.Get("user")
	receipt, _ := c.Get("body")
	id := uuid.MustParse(c.Param("id"))
	id, err := service.ReceiptService.UpdateReceipt(receipt.(*entity.ReceiptRequest), id, user.(*entity.User).Id)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, model.ErrorResponse{
			Error:   err.Error(),
			Message: "Error updating receipt",
			Id:      id.String(),
		})
		return
	}
	c.AbortWithStatus(http.StatusNoContent)
}

// Delete Receipt
// @Summary Delete a receipt
// @Description Delete a receipt
// @Tags Receipts
// @Param id path string true "Receipt id"
// @Success 204
// @Router /budget-control/api/v1/receipt/{id} [delete]
func DeleteReceipt(c *gin.Context) {
	user, _ := c.Get("user")
	id := uuid.MustParse(c.Param("id"))
	service.ReceiptService.DeleteReceipt(id, user.(*entity.User).Id)
	c.AbortWithStatus(http.StatusNoContent)
}

// Find All Receipts By Period
// @Summary Find all receipts by Period
// @Description Find all receipts by Period
// @Tags Receipts
// @Param year path int true "Year"
// @Param month path int true "Month"
// @Success 200 {array} entity.ReceiptResponse
// @Router /budget-control/api/v1/receipt/period/{year}/{month} [get]
func ReceiptsByPeriod(c *gin.Context) {
	user, _ := c.Get("user")
	var receipts []entity.ReceiptResponse
	err := service.ReceiptService.ReceiptsByPeriod(&receipts, c.Param("year"), c.Param("month"), user.(*entity.User).Id)
	if err != nil {
		c.JSON(http.StatusNotFound, model.ErrorResponse{
			Error:   err.Error(),
			Message: "Receipts not found",
		})
	}
	c.JSON(http.StatusOK, receipts)
}
