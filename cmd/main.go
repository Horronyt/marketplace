package main

import (
	"github.com/Horronyt/marketplace"
	"github.com/Horronyt/marketplace/pkg/handler"
	"github.com/Horronyt/marketplace/pkg/repository"
	"github.com/Horronyt/marketplace/pkg/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := IninConfig(); err != nil {
		logrus.Fatalf("Error reading config: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("Error loading .env file: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:    viper.GetString("db.host"),
		Port:    viper.GetString("db.port"),
		User:    viper.GetString("db.username"),
		DBName:  viper.GetString("db.dbname"),
		SSLMode: viper.GetString("db.sslmode"),
		Pass:    os.Getenv("DB_PASSWORD"),
	})
	if err != nil {
		logrus.Fatalf("Error connecting to database: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	srv := new(marketplace.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		logrus.Fatalf("Error starting server: %s", err.Error())
	}
}

func IninConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
