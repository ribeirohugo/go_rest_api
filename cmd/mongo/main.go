package main

import (
	"context"
	"log"

	"github.com/ribeirohugo/golang_startup/internal/common"
	"github.com/ribeirohugo/golang_startup/internal/config"
	"github.com/ribeirohugo/golang_startup/internal/controller"
	"github.com/ribeirohugo/golang_startup/internal/database/mongodb"
	"github.com/ribeirohugo/golang_startup/internal/server"
	"github.com/ribeirohugo/golang_startup/internal/service"
)

func main() {
	cfg := config.Load()

	database, err := mongodb.New(context.Background(), cfg.Database.Address, cfg.Database.Name)
	if err != nil {
		log.Fatalf("failed to initialise the database client: %v", err)
	}

	timer := common.NewTimer()
	services := service.New(database, timer)

	controllers := controller.New(services)

	srv := server.New(controllers, cfg.GetServerAddress())

	srv.ServeHTTP()
}
