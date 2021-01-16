package main

import (
	GEO_Support "github.com/Kushkaftar/geo_support"
	"github.com/Kushkaftar/geo_support/pkg/handler"
	"github.com/Kushkaftar/geo_support/pkg/repository"
	"github.com/Kushkaftar/geo_support/pkg/service"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"log"
	"os"
)

func main() {

	// инициируем конигурацию
	if err := initConfig(); err != nil {
		log.Fatalf("error read config: %s", err.Error())
	}

	// иницируем переменные env
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err.Error())
	}

	db, err := repository.NewMysqlDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.name"),
	})
	if err != nil {
		log.Fatalf("error connect to Data Base: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(GEO_Support.Server)
	if err := srv.Run(viper.GetString("bind_addr"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error run server: %s", err.Error())
	}

}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
