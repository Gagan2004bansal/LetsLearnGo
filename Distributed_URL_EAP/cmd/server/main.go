package main

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"

	"github.com/Gagan2004bansal/LetsLearnGo/internal/config"
	"github.com/Gagan2004bansal/LetsLearnGo/internal/handler"
	"github.com/Gagan2004bansal/LetsLearnGo/internal/redis"
	"github.com/Gagan2004bansal/LetsLearnGo/internal/repository"
	"github.com/Gagan2004bansal/LetsLearnGo/internal/service"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {

	// Load Environment Variables
	if err := godotenv.Load(); err != nil {
		slog.Warn("no .env file found")
	}

	cfg, error := config.Load()
	if error != nil {
		log.Fatal("failed to load configuration")
	}

	// Database Connection Pool
	db, err := repository.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Redis Connection
	rdb, err := redis.ConnectRedis()
	if err != nil {
		log.Fatal("failed to connect to redis", err)
	}
	defer rdb.Close()

	// Initialize Repository
	urlRepository := repository.NewPostgresUrlRepository(db)

	// Initialize Cache Repository
	redisRepository := repository.NewRedisUrlRepository(rdb)

	// Initialize Service
	urlService := service.NewUrlService(urlRepository, redisRepository)

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
