package main

import (
	"context"
	"fmt"
	"github.com/Gagan2004bansal/LetsLearnGo/internal/config"
	"github.com/Gagan2004bansal/LetsLearnGo/internal/handler"
	"github.com/Gagan2004bansal/LetsLearnGo/internal/repository"
	"github.com/Gagan2004bansal/LetsLearnGo/internal/service"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"log/slog"
	"net/http"
)

func main() {

	conn, err := repository.Connect()
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close(context.Background())

	_, err = conn.Exec(
		context.Background(),
		`CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			name TEXT NOT NULL
		)`,
	)
	if err != nil {
		log.Fatal(err)
	}

	_, err = conn.Exec(
		context.Background(),
		"INSERT INTO users (name) VALUES ($1), ($2)",
		"Gagan",
		"Rahul",
	)
	if err != nil {
		log.Fatal(err)
	}

	// Fir retrieve
	repository.QueryData(conn)

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
