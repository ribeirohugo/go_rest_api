package main

import (
	"log"
	"net/http"

	"github.com/ribeirohugo/golang_startup/internal/config"
	"github.com/ribeirohugo/golang_startup/internal/database"
	"github.com/ribeirohugo/golang_startup/internal/server"
	"github.com/ribeirohugo/golang_startup/internal/service"
)

func main() {
	cfg := config.Load()

	postgres, err := database.NewPostgres(cfg.Database.Address)
	if err != nil {
		log.Fatalf("failed to initialise the database client: %v", err)
	}

	userRepo := database.NewUserRepo(postgres.DB)

	userService := service.NewService(userRepo)

	httpServer := server.New(userService)

	err = http.ListenAndServe(":8082", httpServer)
	if err != http.ErrServerClosed {
		log.Printf("http server terminated unexpectedly: %v", err)
	}
}
