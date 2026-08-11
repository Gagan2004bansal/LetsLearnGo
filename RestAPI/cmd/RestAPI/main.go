package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Gagan2004bansal/LetsLearnGo/internal/config"
	"github.com/Gagan2004bansal/LetsLearnGo/internal/http/handlers/student"
	"github.com/Gagan2004bansal/LetsLearnGo/internal/storage/sqlite"
)

func main() {

	// load config
	cfg := config.MustLoad()

	// database setup
	storage, errr := sqlite.New(cfg)
	if errr != nil {
		log.Fatal(errr)
	}

	slog.Info("storage initialize")

	// setup router
	router := http.NewServeMux()

	router.HandleFunc("POST /api/students", student.New(storage))

	// setup server
	server := http.Server{
		Addr:    cfg.Addr,
		Handler: router,
	}

	slog.Info("server started", slog.String("address", cfg.Addr))
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal("failed to start server")
		}
		fmt.Println("server started")
	}()

	<-done

	slog.Info("shutting down the server")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := server.Shutdown(ctx)
	if err != nil {
		slog.Error("failed to shutdown server", slog.String("error", err.Error()))
	}
	slog.Info("server shutdown successfully")
}
