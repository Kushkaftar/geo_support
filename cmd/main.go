package main

import (
	GEO_Support "github.com/Kushkaftar/geo_support"
	"github.com/Kushkaftar/geo_support/pkg/handler"
	"github.com/Kushkaftar/geo_support/pkg/repository"
	"github.com/Kushkaftar/geo_support/pkg/service"
	"log"
)

func main()  {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(GEO_Support.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("error run server: %s", err.Error())
	}

}