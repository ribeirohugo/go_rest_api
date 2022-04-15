package main

import (
	"github.com/ribeirohugo/golang_startup/internal/common"
	"log"

	"github.com/ribeirohugo/golang_startup/internal/config"
	"github.com/ribeirohugo/golang_startup/internal/controller"
	"github.com/ribeirohugo/golang_startup/internal/database/postgres"
	"github.com/ribeirohugo/golang_startup/internal/server"
	"github.com/ribeirohugo/golang_startup/internal/service"
)

func main() {
	cfg := config.Load()

	database, err := postgres.New(cfg.Database.Address)
	if err != nil {
		log.Fatalf("failed to initialise the database client: %v", err)
	}

	err = database.Migrate(cfg.Database.Name, cfg.Database.MigrationsPath)
	if err != nil {
		log.Fatalf("failed initialise database migrations: %v", err)
	}

	timer := common.NewTimer()
	services := service.New(database, timer)

	controllers := controller.New(services)

	srv := server.New(controllers, cfg.GetServerAddress())

	srv.ServeHTTP()
}
