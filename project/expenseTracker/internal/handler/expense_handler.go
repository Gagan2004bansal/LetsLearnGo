package handler

import (
	"encoding/json"
	"letslearngo/project/expenseTracker/internal/model"
	"letslearngo/project/expenseTracker/internal/service"
	"net/http"

	"github.com/gorilla/mux"
)

// Expense Services
type ExpenseHandler struct {
	service *service.ExpenseService
}

func NewExpenseHandler(service *service.ExpenseService) *ExpenseHandler {
	return &ExpenseHandler{service: service}
}

// Expense Routes
func (h *ExpenseHandler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/expense", h.CreateExpense).Methods(http.MethodPost)
}

func (h *ExpenseHandler) CreateExpense(w http.ResponseWriter, r *http.Request) {
	var request model.CreateExpenseRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	if request.TotalAmount < 0 || request.Title == "" || request.Category == "" {
		http.Error(w, "all fields are required and amount must be greater than or equal to 0", http.StatusBadRequest)
		return
	}

	response := h.service.CreateExpense(request)

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(response)
}
