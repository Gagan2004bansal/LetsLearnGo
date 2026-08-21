package main

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"

	"github.com/Gagan2004bansal/LetsLearnGo/internal/config"
	"github.com/Gagan2004bansal/LetsLearnGo/internal/handler"
	"github.com/Gagan2004bansal/LetsLearnGo/internal/service"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load()
	cfg, error := config.Load()
	if error != nil {
		log.Fatal("failed to load configuration")
	}

	// Initialize Service
	urlservice := service.NewUrlService()

	// Initialize Handler
	urlhandler := handler.NewUrlHandler(urlservice)

	// Router Setup
	router := mux.NewRouter()

	urlhandler.UrlRoutes(router)

	// HTTP SERVER
	port := fmt.Sprintf(":%d", cfg.PORT)
	slog.Info("server run...")

	server := &http.Server{Addr: port, Handler: router}

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal("server error")
	}
}
