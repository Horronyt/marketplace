package main

import (
	"github.com/Horronyt/marketplace"
	"github.com/Horronyt/marketplace/pkg/handler"
	"github.com/Horronyt/marketplace/pkg/repository"
	"github.com/Horronyt/marketplace/pkg/service"
	"log"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	srv := new(marketplace.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("Error starting server: %s", err.Error())
	}
}
