package main

import (
	"fmt"
	"letslearngo/project/expenseTracker/internal/config"
	"letslearngo/project/expenseTracker/internal/handler"
	"letslearngo/project/expenseTracker/internal/service"
	"log"
	"log/slog"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	// Load Env file
	godotenv.Load()

	// Load Configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("failed to load configuration")
	}

	// Initialize Service
	expenseService := service.NewExpenseService()

	// Initialize Handler
	expenseHandler := handler.NewExpenseHandler(expenseService)

	// Router Setup
	router := mux.NewRouter()

	expenseHandler.RegisterRoutes(router)

	// HTTP SERVER
	port := fmt.Sprintf(":%d", cfg.PORT)
	slog.Info("http server started")

	server := &http.Server{
		Addr:    port,
		Handler: router,
	}

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal("server error")
	}
}
