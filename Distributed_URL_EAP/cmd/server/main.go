package main

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"

	"github.com/Gagan2004bansal/LetsLearnGo/internal/config"
	"github.com/Gagan2004bansal/LetsLearnGo/internal/handler"
	"github.com/Gagan2004bansal/LetsLearnGo/internal/repository"
	"github.com/Gagan2004bansal/LetsLearnGo/internal/service"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		slog.Warn("no .env file found")
	}

	cfg, error := config.Load()
	if error != nil {
		log.Fatal("failed to load configuration")
	}

	// Database connection pool
	db, err := repository.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Initialize Repository
	urlRepository := repository.NewPostgresUrlRepository(db)

	// Initialize Service
	urlService := service.NewUrlService(urlRepository)

	// Initialize Handler
	urlHandler := handler.NewUrlHandler(urlService)

	// Router Setup
	router := mux.NewRouter()
	urlHandler.UrlRoutes(router)

	// HTTP SERVER
	port := fmt.Sprintf(":%d", cfg.PORT)
	slog.Info("server run...")

	server := &http.Server{Addr: port, Handler: router}

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal("server error")
	}
}
