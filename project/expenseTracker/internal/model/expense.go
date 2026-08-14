package model

import "time"

type Expense struct {
	ID          int64
	Title       string
	Desc        string
	Category    string
	TotalAmount float64
	CreatedAt   time.Time
}

type CreateExpenseRequest struct {
	Title       string  `json:"title"`
	Desc        string  `json:"desc"`
	Category    string  `json:"category"`
	TotalAmount float64 `json:"total_amount"`
}

type ExpenseResponse struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title"`
	Desc        string    `json:"desc"`
	Category    string    `json:"category"`
	TotalAmount float64   `json:"total_amount"`
	CreatedAt   time.Time `json:"created_at"`
}

func NewExpense(title string, desc string, category string, totalAmount float64) *Expense {
	return &Expense{
		Title:       title,
		Desc:        desc,
		Category:    category,
		TotalAmount: totalAmount,
		CreatedAt:   time.Now(),
	}
}

func (e *Expense) ToResponse() *ExpenseResponse {
	return &ExpenseResponse{
		ID:          e.ID,
		Title:       e.Title,
		Desc:        e.Desc,
		Category:    e.Category,
		TotalAmount: e.TotalAmount,
		CreatedAt:   e.CreatedAt,
	}
}
