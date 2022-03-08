package service

import (
	"github.com/RaphaSalomao/gin-budget-control/model"
	"github.com/google/uuid"
)

func BalanceSumary(bs *model.BalanceSumaryResponse, year string, month string, userId uuid.UUID) error {
	totalReceipt, err := ReceiptService.TotalReceiptValueByPeriod(year, month, userId)
	if err != nil {
		return err
	}
	totalExpense, categoryBalance, err := ExpenseService.TotalExpenseValueByPeriod(year, month, userId)
	if err != nil {
		return err
	}
	bs.CategoryBalance = categoryBalance
	bs.TotalExpense = totalExpense
	bs.TotalReceipt = totalReceipt
	bs.MonthBalance = totalReceipt - totalExpense
	return nil
}
