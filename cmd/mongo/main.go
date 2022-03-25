package main

import (
	"context"
	"log"
	"net/http"

	"github.com/ribeirohugo/golang_startup/internal/config"
	"github.com/ribeirohugo/golang_startup/internal/controller"
	"github.com/ribeirohugo/golang_startup/internal/database/mongodb"
	"github.com/ribeirohugo/golang_startup/internal/service"
)

func main() {
	cfg := config.Load()

	database, err := mongodb.New(context.Background(), cfg.Database.Address, cfg.Database.Name)
	if err != nil {
		log.Fatalf("failed to initialise the database client: %v", err)
	}

	userService := service.New(database)

	controllers := controller.New(userService)

	hostAddress := cfg.GetServerAddress()

	err = http.ListenAndServe(hostAddress, controllers)
	if err != http.ErrServerClosed {
		log.Printf("http server terminated unexpectedly: %v", err)
	}
}
