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
	userId := uuid.MustParse(c.GetString("userId"))
	var receipt entity.ReceiptRequest
	if err := c.ShouldBindJSON(&receipt); err != nil {
		c.JSON(http.StatusBadRequest, model.ValidationErrorResponse(err))
		return
	}
	id, err := service.ReceiptService.CreateReceipt(&receipt, userId)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, model.ErrorResponse{
			Error:   err.Error(),
			Message: "Error creating receipt",
			Id:      id,
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
	userId := uuid.MustParse(c.GetString("userId"))
	receipts := []entity.ReceiptResponse{}
	description := c.Query("description")
	service.ReceiptService.FindAllReceipts(&receipts, description, userId)
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
	userId := uuid.MustParse(c.GetString("userId"))
	var receipt entity.ReceiptResponse
	id := uuid.MustParse(c.Param("id"))
	err := service.ReceiptService.FindReceipt(&receipt, id, userId)
	if err != nil {
		c.JSON(http.StatusNotFound, model.ErrorResponse{
			Error:   err.Error(),
			Message: "Receipt not found",
			Id:      id,
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
	userId := uuid.MustParse(c.GetString("userId"))
	var receipt entity.ReceiptRequest
	id := uuid.MustParse(c.Param("id"))
	if err := c.ShouldBindJSON(&receipt); err != nil {
		c.JSON(http.StatusBadRequest, model.ValidationErrorResponse(err))
		return
	}
	id, err := service.ReceiptService.UpdateReceipt(&receipt, id, userId)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, model.ErrorResponse{
			Error:   err.Error(),
			Message: "Error updating receipt",
			Id:      id,
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
	userId := uuid.MustParse(c.GetString("userId"))
	id := uuid.MustParse(c.Param("id"))
	service.ReceiptService.DeleteReceipt(id, userId)
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
	userId := uuid.MustParse(c.GetString("userId"))
	var receipts []entity.ReceiptResponse
	err := service.ReceiptService.ReceiptsByPeriod(&receipts, c.Param("year"), c.Param("month"), userId)
	if err != nil {
		c.JSON(http.StatusNotFound, model.ErrorResponse{
			Error:   err.Error(),
			Message: "Receipts not found",
		})
	}
	c.JSON(http.StatusOK, receipts)
}
