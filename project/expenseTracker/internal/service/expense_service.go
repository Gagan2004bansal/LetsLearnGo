package service

import (
	"letslearngo/project/expenseTracker/internal/model"
)

type ExpenseService struct {
}

func NewExpenseService() *ExpenseService {
	return &ExpenseService{}
}

func (e *ExpenseService) CreateExpense(request model.CreateExpenseRequest) *model.ExpenseResponse {

	expense := model.NewExpense(request.Title, request.Desc, request.Category, request.TotalAmount)

	return expense.ToResponse()
}
